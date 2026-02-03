# Plants Backend

Go REST API + notification worker for plant management. Uses Firebase auth, MongoDB, AWS S3.

## Components

| Component           | Purpose                          | Port | Scaling         |
| ------------------- | -------------------------------- | ---- | --------------- |
| API Server          | REST endpoints + upload URLs     | 8080 | Horizontal      |
| Notification Worker | Push notifications (every 5 min) | —    | Single instance |

Shared packages: types/, services/, constants/, validation/, util/

## Quick Start

```bash
cd backend
go run ./cmd/api                 # API on :8080
go run ./cmd/notification-worker # Worker (optional)
```

## Key Endpoints

All require: Authorization: Bearer <firebase-id-token>

| Method         | Path                                 | Purpose               |
| -------------- | ------------------------------------ | --------------------- |
| GET            | /api/plants                          | List plants           |
| POST           | /api/plants                          | Create plant          |
| PATCH          | /api/plants/{id}                     | Update plant          |
| DELETE         | /api/plants/{id}                     | Delete plant          |
| POST           | /api/plants/water                    | Mark watered          |
| GET            | /api/upload/presigned-url            | S3 upload URL         |
| GET/PUT/DELETE | /api/notifications                   | Notification config   |
| POST           | /api/notifications/tokens            | Register device token |
| DELETE         | /api/notifications/tokens/{deviceId} | Remove device token   |

## Notification System

- Runs every 5 minutes
- Batch sizes: 1000 plants, 500 FCM tokens
- Cooldown: 4 hours per user
- Failed tokens marked inactive (isActive: false)
- 8 message variants per type (4 single + 4 multiple)

Message templates live in messages/:

```json
{
  "single": [{ "title": "...", "body": "{plantName} needs water" }],
  "multiple": [{ "title": "...", "body": "{count} plants: {plantNames}" }]
}
```

Template variables: {plantName}, {count}, {plantNames}, {remaining}

## Database Schema (Notifications)

```typescript
{
	_id: string,
	userId: string,
	deviceTokens: [
		{ token: string, deviceId: string, deviceType: string, addedAt: Date, isActive: boolean }
	],
	isEnabled: boolean,
	preferredTime: string,
	quietHours: { start: string, end: string },
	mutedPlantIds: [string],
	remindWatering: boolean,
	remindFertilize: boolean,
	remindRepotting: boolean,
	remindMisting: boolean,
	lastNotificationSentAt: Date,
	updatedAt: Date
}
```

## Features

- Firebase Authentication (Google sign-in)
- Plant CRUD with MongoDB
- AWS S3 photo uploads (presigned URLs)
- Push notifications via FCM
- User data isolation
- Orphaned S3 cleanup (30 min interval)
- CORS configured

## Build & Deploy

```bash
./build.sh
docker build --build-arg COMPONENT=api .
docker build --build-arg COMPONENT=notification-worker .
kubectl apply -f infra/k8s/api.yaml
```

## Tests

```bash
go test ./routes -v
```
