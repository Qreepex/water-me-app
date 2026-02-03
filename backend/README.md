# Plants Backend - Go Services

A RESTful API and notification worker for managing plants with Firebase authentication and MongoDB storage.

## Components

### 1. **API Server** (`cmd/api`)

REST API for plant management with:

- User authentication (Firebase)
- Plant CRUD operations
- Photo uploads (AWS S3)
- Notification settings
- Orphaned upload cleanup worker

### 2. **Notification Worker** (`cmd/notification-worker`)

Background service that:

- Checks for plants needing water every 5 minutes
- Sends push notifications via Firebase Cloud Messaging
- Operates independently from the API

## Features

- **User Authentication**: Firebase Authentication (Google sign-in) with JWT token validation
- **Plant Management**: Full CRUD operations for plants with rich metadata support
- **Photo Storage**: AWS S3 integration with presigned URLs for direct browser uploads
- **Push Notifications**: Automated watering reminders via Firebase Cloud Messaging
- **User Isolation**: Each user can only access their own plants and data
- **Background Jobs**: Automatic cleanup of orphaned S3 uploads every 30 minutes
- **CORS**: Configured for multiple origins (localhost, app.water-me.app, etc.)

## Tech Stack

- **Language**: Go 1.25
- **Database**: MongoDB with official driver
- **Authentication**: Firebase Admin SDK
- **Storage**: AWS S3
- **HTTP**: Gorilla Mux router + CORS middleware
- **Dependencies**: Managed with Go modules

## Environment Variables

Create a `.env` file in the `backend/` directory:

```bash
DATABASE_URL=mongodb://localhost:27017
MONGODB_USERNAME=test2
MONGODB_PASSWORD=test
MONGODB_DATABASE=plants

# Firebase credentials file path
GOOGLE_APPLICATION_CREDENTIALS=./secret/fb.json

# AWS S3 credentials (uses AWS SDK default chain)
AWS_REGION=us-east-1
AWS_S3_BUCKET=your-bucket-name

PORT=8080
```

Firebase credentials should be placed in `backend/secret/fb.json`.

## Setup

1. **Install dependencies**:

```bash
cd backend
go mod tidy
```

1. **Set up MongoDB**:

```bash
# Start MongoDB locally or use a cloud instance
# The server will automatically create collections on first use
```

1. **Configure Firebase**:
   - Download your Firebase service account key from Firebase Console
   - Place it in `backend/secret/fb.json`

2. **Configure AWS S3**:
   - Create an S3 bucket for plant photos
   - Configure AWS credentials (via environment vars, ~/.aws/credentials, or IAM role)

3. **Run the API server**:

```bash
go run ./cmd/api
```

4. **Run the notification worker** (optional, in separate terminal):

```bash
go run ./cmd/notification-worker
```

The API server starts on port 8080 and connects to MongoDB on startup.
The notification worker runs checks every 5 minutes.

## Project Structure

```
backend/
├── cmd/                  Command-line binaries
│   ├── api/              API server entry point
│   ├── notification-worker/  Notification service
│   └── setup-cors/       S3 CORS setup utility
├── constants/            Application constants (collections, limits)
├── main.go               Legacy entry point (deprecated, use cmd/api)
├── middlewares/          HTTP middlewares (auth, rate limiting)
├── routes/               HTTP endpoint handlers
│   ├── plants.go         Plant CRUD + watering
│   ├── uploads.go        S3 presigned URL generation
│   ├── notifications.go  Push notification config
│   └── stats.go          User statistics
├── services/             Business logic layer (shared)
│   ├── mongo.go          MongoDB connection
│   ├── database.go       Plant/notification/upload queries
│   ├── firebase.go       Firebase auth token verification
│   ├── s3.go             S3 upload/download/presigned URLs
│   ├── uploads.go        Upload service with orphan cleanup
│   └── ratelimit.go      Rate limiting service
├── types/                Data models and enums (shared)
├── util/                 Helper functions (shared)
├── validation/           Input validation (mirrors frontend, shared)
├── infra/                Infrastructure as code
│   └── k8s/              Kubernetes manifests
└── openapi.yaml          API specification
```

See [cmd/README.md](cmd/README.md) for detailed build and deployment instructions.

## API Endpoints

All endpoints are documented in `openapi.yaml`. Key endpoints:

### Plants (Protected)

- `GET /api/plants` - List all user's plants
- `POST /api/plants` - Create plant (auto-generates slug)
- `GET /api/plants/slug/{slug}` - Get plant by slug
- `GET /api/plants/{id}` - Get plant by MongoDB ObjectID
- `PATCH /api/plants/{id}` - Partially update plant
- `DELETE /api/plants/{id}` - Delete plant
- `POST /api/plants/water` - Mark plants as watered (bulk operation)

### Uploads (Protected)

- `GET /api/upload/presigned-url` - Get presigned S3 URL for photo upload
- `GET /api/photo/{key}` - Get presigned download URL for photo

### Notifications (Protected)

- `GET /api/notifications/config` - Get user's notification config
- `POST /api/notifications/config` - Create notification config
- `PATCH /api/notifications/config` - Update notification config
- `DELETE /api/notifications/config` - Delete notification config

All protected endpoints require `Authorization: Bearer <firebase-id-token>` header.

## Authentication Flow

1. Frontend uses Firebase Authentication to sign in user (Google OAuth)
2. Frontend obtains Firebase ID token
3. Frontend sends ID token in `Authorization: Bearer <token>` header
4. Backend middleware validates token with Firebase Admin SDK
5. User ID is extracted and stored in request context
6. All database operations are scoped to the authenticated user

## MongoDB Collections

### plants

- **Fields**: name, species, slug, sunlight, location, watering config, fertilizing config, humidity, soil, seasonality, pest history, flags, notes, photo IDs, growth history
- **Indexes**: userId (for user isolation), slug (for friendly URLs)
- **ID Format**: MongoDB ObjectID

### notifications

- **Fields**: userId, fcmToken, enabled, preferences
- **Purpose**: Store push notification configuration per user

### uploads

- **Fields**: userId, key (S3 object key), sizeBytes, createdAt
- **Purpose**: Track S3 uploads for quota enforcement and orphan cleanup

## Key Design Patterns

### User Scoping

All MongoDB queries include `userId` filter to ensure data isolation. The auth middleware extracts the user ID from the Firebase token and stores it in the request context.

### Slug Generation

Plants have both MongoDB ObjectIDs (for internal operations) and human-readable slugs (for URLs). Slugs are auto-generated from species name on creation and ensured to be unique per user.

### Presigned URLs

Photo uploads use S3 presigned URLs for direct browser-to-S3 uploads, avoiding backend proxy overhead. The backend validates file types and sizes before generating URLs.

### Orphan Cleanup

A background worker runs every 30 minutes to delete S3 objects that exist in the `uploads` collection but aren't referenced by any plant's `photoIds` field.

## Validation

Validation logic exists in two places:

- `backend/validation.ts` - TypeScript validation (used by frontend during development)
- `backend/validation/` - Go validation (enforced by backend)

**When updating validation rules, update both to maintain consistency.**

## Development

```bash
# Run server with auto-reload (use external tool like air)
go run .

# Run tests (if any exist)
go test ./...

# Format code
go fmt ./...
```

## Testing

Example using curl:

```bash
# Get Firebase token (use Firebase Auth in your app)
TOKEN="your-firebase-id-token"

# List plants
curl http://localhost:8080/api/plants \
  -H "Authorization: Bearer $TOKEN"

# Create plant
curl -X POST http://localhost:8080/api/plants \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "My Monstera",
    "species": "Monstera Deliciosa",
    "isToxic": true,
    "sunlight": "Indirect Sun",
    "location": {"room": "Living Room", "position": "Window", "isOutdoors": false},
    "watering": {"intervalDays": 7, "method": "Top", "waterType": "Filtered"},
    "fertilizing": {"intervalDays": 30, "type": "Liquid"},
    "humidity": {"min": 40, "max": 60}
  }'

# Get presigned upload URL
curl -X GET "http://localhost:8080/api/upload/presigned-url?contentType=image/jpeg&size=1048576" \
  -H "Authorization: Bearer $TOKEN"

# Water plants (bulk)
curl -X POST http://localhost:8080/api/plants/water \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"plantIds": ["507f1f77bcf86cd799439011"]}'
```
