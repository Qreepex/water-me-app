package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"plants-backend/types"
	"plants-backend/validation"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Database wraps database access.
type Database struct {
	pool *pgxpool.Pool
}

func NewDatabase(ctx context.Context, connString string) (*Database, error) {
	cfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("parse postgres config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("connect postgres: %w", err)
	}

	s := &Database{pool: pool}
	if err := s.ensureSchema(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	return s, nil
}

func (s *Database) Close() {
	if s.pool != nil {
		s.pool.Close()
	}
}

func (s *Database) ensureSchema(ctx context.Context) error {
	const ddl = `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		username TEXT,
		email TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL,
		language TEXT NOT NULL DEFAULT 'en',
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS plants (
		id TEXT PRIMARY KEY,
		user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		species TEXT NOT NULL,
		name TEXT NOT NULL,
		sun_light TEXT NOT NULL,
		prefered_temperature REAL NOT NULL,
		watering_interval_days INTEGER NOT NULL,
		last_watered TIMESTAMPTZ NOT NULL,
		fertilizing_interval_days INTEGER NOT NULL,
		last_fertilized TIMESTAMPTZ NOT NULL,
		prefered_humidity REAL NOT NULL,
		spray_interval_days INTEGER NULL,
		notes JSONB NOT NULL DEFAULT '[]',
		flags JSONB NOT NULL DEFAULT '[]',
		photo_ids JSONB NOT NULL DEFAULT '[]',
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	);

	CREATE INDEX IF NOT EXISTS idx_plants_user_id ON plants(user_id);
	`
	if _, err := s.pool.Exec(ctx, ddl); err != nil {
		return fmt.Errorf("create schema: %w", err)
	}
	return nil
}

func (s *Database) ListPlants(ctx context.Context, userID string) ([]types.Plant, error) {
	const query = `
	SELECT id, user_id, species, name, sun_light, prefered_temperature,
		watering_interval_days, last_watered, fertilizing_interval_days,
		last_fertilized, prefered_humidity, spray_interval_days,
		notes, flags, photo_ids
	FROM plants
	WHERE user_id = $1
	ORDER BY name ASC`

	rows, err := s.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("query plants: %w", err)
	}
	defer rows.Close()

	plants := make([]types.Plant, 0)
	for rows.Next() {
		plant, err := scanPlant(rows)
		if err != nil {
			return nil, err
		}
		plants = append(plants, plant)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return plants, nil
}

func (s *Database) GetPlant(ctx context.Context, id string, userID string) (types.Plant, bool, error) {
	const query = `
	SELECT id, user_id, species, name, sun_light, prefered_temperature,
		watering_interval_days, last_watered, fertilizing_interval_days,
		last_fertilized, prefered_humidity, spray_interval_days,
		notes, flags, photo_ids
	FROM plants
	WHERE id = $1 AND user_id = $2`

	row := s.pool.QueryRow(ctx, query, id, userID)
	plant, err := scanPlant(row)
	if err != nil {
		if errors.Is(err, errNotFound) {
			return types.Plant{}, false, nil
		}
		return types.Plant{}, false, err
	}
	return plant, true, nil
}

func (s *Database) CreatePlant(ctx context.Context, userID string, input types.PlantInput) (types.Plant, error) {
	now := time.Now().UTC()
	id := generateID(input.ID)

	lastWatered := now
	if input.LastWatered != nil {
		parsed, err := parseTime(*input.LastWatered)
		if err != nil {
			return types.Plant{}, fmt.Errorf("invalid lastWatered: %w", err)
		}
		lastWatered = parsed
	}

	lastFertilized := now
	if input.LastFertilized != nil {
		parsed, err := parseTime(*input.LastFertilized)
		if err != nil {
			return types.Plant{}, fmt.Errorf("invalid lastFertilized: %w", err)
		}
		lastFertilized = parsed
	}

	notesJSON, flagsJSON, photoJSON, err := marshalCollections(input)
	if err != nil {
		return types.Plant{}, err
	}

	spray := interface{}(nil)
	if input.SprayIntervalDays != nil {
		spray = *input.SprayIntervalDays
	}

	const query = `
	INSERT INTO plants (
		id, user_id, species, name, sun_light, prefered_temperature,
		watering_interval_days, last_watered, fertilizing_interval_days,
		last_fertilized, prefered_humidity, spray_interval_days,
		notes, flags, photo_ids, created_at, updated_at
	) VALUES (
		$1, $2, $3, $4, $5,
		$6, $7, $8, $9,
		$10, $11, $12,
		$13, $14, $15, $16, $17
	) RETURNING 
		id, user_id, species, name, sun_light, prefered_temperature,
		watering_interval_days, last_watered, fertilizing_interval_days,
		last_fertilized, prefered_humidity, spray_interval_days,
		notes, flags, photo_ids`

	row := s.pool.QueryRow(
		ctx,
		query,
		id,
		userID,
		valueOrDefault(input.Species),
		valueOrDefault(input.Name),
		valueOrDefault(input.SunLight),
		valueOrDefault(input.PreferedTemperature),
		valueOrDefault(input.WateringIntervalDays),
		lastWatered,
		valueOrDefault(input.FertilizingIntervalDays),
		lastFertilized,
		valueOrDefault(input.PreferedHumidity),
		spray,
		notesJSON,
		flagsJSON,
		photoJSON,
		now,
		now,
	)

	plant, err := scanPlant(row)
	if err != nil {
		return types.Plant{}, err
	}
	return plant, nil
}

func (s *Database) UpdatePlant(ctx context.Context, id string, userID string, updates types.PlantInput) (types.Plant, bool, error) {
	existing, found, err := s.GetPlant(ctx, id, userID)
	if err != nil || !found {
		return types.Plant{}, found, err
	}

	merged, err := mergePlant(existing, updates)
	if err != nil {
		return types.Plant{}, false, err
	}

	notesJSON, flagsJSON, photoJSON, err := marshalCollections(types.PlantInput{
		Notes:    &merged.Notes,
		Flags:    &merged.Flags,
		PhotoIDs: &merged.PhotoIDs,
	})
	if err != nil {
		return types.Plant{}, false, err
	}

	spray := interface{}(nil)
	if merged.SprayIntervalDays != nil {
		spray = *merged.SprayIntervalDays
	}

	lastWateredTime, err := parseTime(merged.LastWatered)
	if err != nil {
		return types.Plant{}, false, err
	}
	lastFertilizedTime, err := parseTime(merged.LastFertilized)
	if err != nil {
		return types.Plant{}, false, err
	}

	now := time.Now().UTC()

	const query = `
	UPDATE plants SET
		species = $1,
		name = $2,
		sun_light = $3,
		prefered_temperature = $4,
		watering_interval_days = $5,
		last_watered = $6,
		fertilizing_interval_days = $7,
		last_fertilized = $8,
		prefered_humidity = $9,
		spray_interval_days = $10,
		notes = $11,
		flags = $12,
		photo_ids = $13,
		updated_at = $14
	WHERE id = $15 AND user_id = $16
	RETURNING id, user_id, species, name, sun_light, prefered_temperature,
		watering_interval_days, last_watered, fertilizing_interval_days,
		last_fertilized, prefered_humidity, spray_interval_days,
		notes, flags, photo_ids`

	row := s.pool.QueryRow(
		ctx,
		query,
		merged.Species,
		merged.Name,
		merged.SunLight,
		merged.PreferedTemperature,
		merged.WateringIntervalDays,
		lastWateredTime,
		merged.FertilizingIntervalDays,
		lastFertilizedTime,
		merged.PreferedHumidity,
		spray,
		notesJSON,
		flagsJSON,
		photoJSON,
		now,
		id,
		userID,
	)

	plant, err := scanPlant(row)
	if err != nil {
		return types.Plant{}, false, err
	}
	return plant, true, nil
}

func (s *Database) DeletePlant(ctx context.Context, id string, userID string) (bool, error) {
	const query = `DELETE FROM plants WHERE id = $1 AND user_id = $2`
	cmdTag, err := s.pool.Exec(ctx, query, id, userID)
	if err != nil {
		return false, fmt.Errorf("delete plant: %w", err)
	}
	return cmdTag.RowsAffected() > 0, nil
}

var errNotFound = errors.New("plant not found")

func scanPlant(row interface{ Scan(dest ...any) error }) (types.Plant, error) {
	var (
		id             string
		userID         string
		species        string
		name           string
		sunLight       string
		prefTemp       float64
		watering       int
		lastWatered    time.Time
		fertInterval   int
		lastFertilized time.Time
		prefHumidity   float64
		spray          sql.NullInt32
		notesBytes     []byte
		flagsBytes     []byte
		photoBytes     []byte
	)

	if err := row.Scan(
		&id,
		&userID,
		&species,
		&name,
		&sunLight,
		&prefTemp,
		&watering,
		&lastWatered,
		&fertInterval,
		&lastFertilized,
		&prefHumidity,
		&spray,
		&notesBytes,
		&flagsBytes,
		&photoBytes,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return types.Plant{}, errNotFound
		}
		return types.Plant{}, err
	}

	notes := make([]string, 0)
	if len(notesBytes) > 0 {
		_ = json.Unmarshal(notesBytes, &notes)
	}

	flagsStr := make([]string, 0)
	if len(flagsBytes) > 0 {
		_ = json.Unmarshal(flagsBytes, &flagsStr)
	}
	flags := make([]types.PlantFlag, 0, len(flagsStr))
	for _, f := range flagsStr {
		flag := types.PlantFlag(f)
		if validation.IsPlantFlag(flag) {
			flags = append(flags, flag)
		}
	}

	photoIDs := make([]string, 0)
	if len(photoBytes) > 0 {
		_ = json.Unmarshal(photoBytes, &photoIDs)
	}

	plant := types.Plant{
		ID:                      id,
		UserID:                  userID,
		Species:                 species,
		Name:                    name,
		SunLight:                types.SunlightRequirement(sunLight),
		PreferedTemperature:     prefTemp,
		WateringIntervalDays:    watering,
		LastWatered:             lastWatered.UTC().Format(time.RFC3339Nano),
		FertilizingIntervalDays: fertInterval,
		LastFertilized:          lastFertilized.UTC().Format(time.RFC3339Nano),
		PreferedHumidity:        prefHumidity,
		Notes:                   notes,
		Flags:                   flags,
		PhotoIDs:                photoIDs,
	}
	if spray.Valid {
		val := int(spray.Int32)
		plant.SprayIntervalDays = &val
	}

	return plant, nil
}

func generateID(provided *string) string {
	if provided != nil && strings.TrimSpace(*provided) != "" {
		return strings.TrimSpace(*provided)
	}
	return fmt.Sprintf("plant_%d_%s", time.Now().UnixMilli(), uuid.NewString())
}

func marshalCollections(input types.PlantInput) ([]byte, []byte, []byte, error) {
	notes := []string{}
	if input.Notes != nil {
		notes = *input.Notes
	}
	flags := []types.PlantFlag{}
	if input.Flags != nil {
		flags = *input.Flags
	}
	photos := []string{}
	if input.PhotoIDs != nil {
		photos = *input.PhotoIDs
	}

	notesJSON, err := json.Marshal(notes)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("marshal notes: %w", err)
	}
	flagsJSON, err := json.Marshal(flags)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("marshal flags: %w", err)
	}
	photoJSON, err := json.Marshal(photos)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("marshal photoIds: %w", err)
	}

	return notesJSON, flagsJSON, photoJSON, nil
}

func mergePlant(existing types.Plant, updates types.PlantInput) (types.Plant, error) {
	merged := existing

	if updates.Species != nil {
		merged.Species = *updates.Species
	}
	if updates.Name != nil {
		merged.Name = *updates.Name
	}
	if updates.SunLight != nil {
		merged.SunLight = *updates.SunLight
	}
	if updates.PreferedTemperature != nil {
		merged.PreferedTemperature = *updates.PreferedTemperature
	}
	if updates.WateringIntervalDays != nil {
		merged.WateringIntervalDays = *updates.WateringIntervalDays
	}
	if updates.LastWatered != nil {
		merged.LastWatered = *updates.LastWatered
	}
	if updates.FertilizingIntervalDays != nil {
		merged.FertilizingIntervalDays = *updates.FertilizingIntervalDays
	}
	if updates.LastFertilized != nil {
		merged.LastFertilized = *updates.LastFertilized
	}
	if updates.PreferedHumidity != nil {
		merged.PreferedHumidity = *updates.PreferedHumidity
	}
	if updates.SprayIntervalDays != nil {
		val := *updates.SprayIntervalDays
		merged.SprayIntervalDays = &val
	}
	if updates.Notes != nil {
		merged.Notes = *updates.Notes
	}
	if updates.Flags != nil {
		merged.Flags = *updates.Flags
	}
	if updates.PhotoIDs != nil {
		merged.PhotoIDs = *updates.PhotoIDs
	}

	return merged, nil
}

func parseTime(val string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339Nano, val)
	if err != nil {
		return time.Time{}, fmt.Errorf("parse time: %w", err)
	}
	return t.UTC(), nil
}

func valueOrDefault[T any](ptr *T) T {
	var zero T
	if ptr == nil {
		return zero
	}
	return *ptr
}

// User operations

func (s *Database) CreateUser(ctx context.Context, email, passwordHash string) (types.User, error) {
	id := fmt.Sprintf("user_%d_%s", time.Now().UnixMilli(), uuid.NewString())
	now := time.Now().UTC()

	const query = `
	INSERT INTO users (id, username, email, password_hash, language, created_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, username, email, language, created_at`

	var user types.User
	var createdAt time.Time

	err := s.pool.QueryRow(ctx, query, id, "", email, passwordHash, "en", now).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Language,
		&createdAt,
	)
	if err != nil {
		return types.User{}, fmt.Errorf("create user: %w", err)
	}

	user.CreatedAt = createdAt.UTC().Format(time.RFC3339Nano)
	return user, nil
}

func (s *Database) GetUserByEmail(ctx context.Context, email string) (types.User, bool, error) {
	const query = `
	SELECT id, username, email, password_hash, language, created_at
	FROM users
	WHERE email = $1`

	var user types.User
	var createdAt time.Time

	err := s.pool.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.Language,
		&createdAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return types.User{}, false, nil
		}
		return types.User{}, false, fmt.Errorf("get user by email: %w", err)
	}

	user.CreatedAt = createdAt.UTC().Format(time.RFC3339Nano)
	return user, true, nil
}

func (s *Database) GetUserByID(ctx context.Context, userID string) (types.User, bool, error) {
	const query = `
	SELECT id, username, email, language, created_at
	FROM users
	WHERE id = $1`

	var user types.User
	var createdAt time.Time

	err := s.pool.QueryRow(ctx, query, userID).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Language,
		&createdAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return types.User{}, false, nil
		}
		return types.User{}, false, fmt.Errorf("get user by id: %w", err)
	}

	user.CreatedAt = createdAt.UTC().Format(time.RFC3339Nano)
	return user, true, nil
}

func (s *Database) UpdateUser(ctx context.Context, userID string, updates types.UpdateUserRequest) (types.User, error) {
	var setClauses []string
	var args []interface{}
	argIndex := 1

	if updates.Username != nil {
		setClauses = append(setClauses, fmt.Sprintf("username = $%d", argIndex))
		args = append(args, *updates.Username)
		argIndex++
	}

	if updates.Email != nil {
		setClauses = append(setClauses, fmt.Sprintf("email = $%d", argIndex))
		args = append(args, *updates.Email)
		argIndex++
	}

	if updates.Language != nil {
		setClauses = append(setClauses, fmt.Sprintf("language = $%d", argIndex))
		args = append(args, *updates.Language)
		argIndex++
	}

	if len(setClauses) == 0 {
		// No updates provided, just return current user
		user, _, err := s.GetUserByID(ctx, userID)
		return user, err
	}

	args = append(args, userID)
	query := fmt.Sprintf(`
	UPDATE users
	SET %s
	WHERE id = $%d
	RETURNING id, username, email, language, created_at`,
		strings.Join(setClauses, ", "),
		argIndex,
	)

	var user types.User
	var createdAt time.Time

	err := s.pool.QueryRow(ctx, query, args...).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Language,
		&createdAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return types.User{}, fmt.Errorf("user not found")
		}
		return types.User{}, fmt.Errorf("update user: %w", err)
	}

	user.CreatedAt = createdAt.UTC().Format(time.RFC3339Nano)
	return user, nil
}

func (s *Database) UpdateUserPassword(ctx context.Context, userID string, newPasswordHash string) error {
	const query = `
	UPDATE users
	SET password_hash = $1
	WHERE id = $2`

	result, err := s.pool.Exec(ctx, query, newPasswordHash, userID)
	if err != nil {
		return fmt.Errorf("update password: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}
