# Plants API - Go Backend

A RESTful API for managing plants with user authentication and PostgreSQL storage.

## Features

- **User Authentication**: JWT-based authentication with signup and login
- **Plant Management**: Full CRUD operations for plants
- **User Isolation**: Each user can only see and manage their own plants
- **Validation**: Input validation matching the TypeScript validation rules
- **CORS**: Configured for wildcard origin (`*`)

## Environment Variables

```bash
DATABASE_URL=postgres://postgres:password@localhost:5432/plants?sslmode=disable
JWT_SECRET=your-secret-key-change-this-in-production
PORT=8080
```

## Setup

1. Install dependencies:

```bash
cd backend
go mod tidy
```

1. Set up PostgreSQL database:

```bash
createdb plants
```

1. Run the server:

```bash
go run .
```

The server will automatically create the necessary tables on startup.

## API Endpoints

### Authentication

#### POST /api/signup

Create a new user account.

**Request:**

```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response (201):**

```json
{
  "token": "eyJhbGc...",
  "user": {
    "id": "user_1234...",
    "email": "user@example.com",
    "createdAt": "2026-01-17T..."
  }
}
```

#### POST /api/login

Authenticate an existing user.

**Request:**

```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response (200):**

```json
{
  "token": "eyJhbGc...",
  "user": {
    "id": "user_1234...",
    "email": "user@example.com",
    "createdAt": "2026-01-17T..."
  }
}
```

### Plants (Protected - Requires Authorization Header)

All plant endpoints require the `Authorization: Bearer <token>` header.

#### GET /api/plants

List all plants for the authenticated user.

**Response (200):**

```json
[
  {
    "id": "plant_1234...",
    "userId": "user_1234...",
    "species": "Monstera Deliciosa",
    "name": "My Plant",
    ...
  }
]
```

#### POST /api/plants

Create a new plant.

**Request:**

```json
{
  "species": "Monstera Deliciosa",
  "name": "My Plant",
  "sunLight": "Indirect Sun",
  "preferedTemperature": 22,
  "wateringIntervalDays": 7,
  "fertilizingIntervalDays": 30,
  "preferedHumidity": 60,
  "notes": [],
  "flags": [],
  "photoIds": []
}
```

**Response (201):** Returns the created plant object.

#### GET /api/plants/{plantId}

Get a specific plant by ID.

**Response (200):** Returns the plant object or 404 if not found or not owned by user.

#### PUT /api/plants/{plantId}

Replace a plant (requires all fields).

**Request:** Same as POST with all required fields.

**Response (200):** Returns the updated plant object.

#### PATCH /api/plants/{plantId}

Partially update a plant (only provided fields).

**Request:** Any subset of plant fields.

**Response (200):** Returns the updated plant object.

#### DELETE /api/plants/{plantId}

Delete a plant.

**Response (200):**

```json
{
  "success": true
}
```

## Security

- Passwords are hashed using bcrypt
- JWT tokens expire after 7 days
- Each plant is associated with a user via `user_id` foreign key
- All plant operations are filtered by the authenticated user's ID
- Unauthorized requests return 401

## Database Schema

### users

- `id` TEXT PRIMARY KEY
- `email` TEXT UNIQUE NOT NULL
- `password_hash` TEXT NOT NULL
- `created_at` TIMESTAMPTZ NOT NULL

### plants

- `id` TEXT PRIMARY KEY
- `user_id` TEXT NOT NULL (foreign key to users)
- `species` TEXT NOT NULL
- `name` TEXT NOT NULL
- `sun_light` TEXT NOT NULL
- `prefered_temperature` REAL NOT NULL
- `watering_interval_days` INTEGER NOT NULL
- `last_watered` TIMESTAMPTZ NOT NULL
- `fertilizing_interval_days` INTEGER NOT NULL
- `last_fertilized` TIMESTAMPTZ NOT NULL
- `prefered_humidity` REAL NOT NULL
- `spray_interval_days` INTEGER NULL
- `notes` JSONB NOT NULL DEFAULT '[]'
- `flags` JSONB NOT NULL DEFAULT '[]'
- `photo_ids` JSONB NOT NULL DEFAULT '[]'
- `created_at` TIMESTAMPTZ NOT NULL
- `updated_at` TIMESTAMPTZ NOT NULL

## Testing

Example using curl:

```bash
# Signup
curl -X POST http://localhost:8080/api/signup \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'

# Login
TOKEN=$(curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}' \
  | jq -r '.token')

# List plants
curl http://localhost:8080/api/plants \
  -H "Authorization: Bearer $TOKEN"

# Create plant
curl -X POST http://localhost:8080/api/plants \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "species": "Monstera Deliciosa",
    "name": "My Plant",
    "sunLight": "Indirect Sun",
    "preferedTemperature": 22,
    "wateringIntervalDays": 7,
    "fertilizingIntervalDays": 30,
    "preferedHumidity": 60
  }'
```
