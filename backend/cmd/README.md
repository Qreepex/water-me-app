# Water Me Backend - Build & Deployment

The backend is now split into two components:

## Components

### 1. API Server (`cmd/api`)

- REST API for plant management
- Handles authentication, plant CRUD, photo uploads, notifications settings
- Runs on port 8080
- Serves HTTP endpoints

### 2. Notification Worker (`cmd/notification-worker`)

- Background service that sends push notifications
- Checks for plants needing water every 5 minutes
- Sends FCM push notifications to users
- No HTTP endpoints

## Shared Code

Both components share:

- `types/` - Data structures (plants, notifications, uploads, errors)
- `services/` - Database, Firebase, S3, authentication, rate limiting
- `constants/` - Collection names, limits, MIME types
- `validation/` - Input validation logic
- `util/` - Helper functions

## Local Development

### Run API Server

```bash
cd backend
go run ./cmd/api
```

### Run Notification Worker

```bash
cd backend
go run ./cmd/notification-worker
```

### Run Both (separate terminals)

```bash
# Terminal 1 - API
go run ./cmd/api

# Terminal 2 - Notification Worker
go run ./cmd/notification-worker
```

## Docker Build

### Build API Server

```bash
docker build --build-arg COMPONENT=api -t plants-backend-api:latest .
```

### Build Notification Worker

```bash
docker build --build-arg COMPONENT=notification-worker -t plants-notification-worker:latest .
```

### Build Both for Production

```bash
# API
docker build --build-arg COMPONENT=api -t ghcr.io/qreepex/plants-backend-api:0.0.6 .
docker push ghcr.io/qreepex/plants-backend-api:0.0.6

# Notification Worker
docker build --build-arg COMPONENT=notification-worker -t ghcr.io/qreepex/plants-notification-worker:0.0.6 .
docker push ghcr.io/qreepex/plants-notification-worker:0.0.6
```

## Kubernetes Deployment

Deploy both components:

```bash
kubectl apply -f infra/k8s/namespace.yaml
kubectl apply -f infra/k8s/secrets.yaml
kubectl apply -f infra/k8s/api.yaml
```

This creates:

- **backend-api** Deployment (2 replicas) - API server with HTTP service + Ingress
- **notification-worker** Deployment (1 replica) - Background worker

## Environment Variables

Both components require:

```bash
DATABASE_URL=mongodb://localhost:27017
MONGODB_USERNAME=test2
MONGODB_PASSWORD=test
MONGODB_DATABASE=plants

GOOGLE_APPLICATION_CREDENTIALS=./secret/fb.json

AWS_REGION=us-east-1
AWS_S3_BUCKET=your-bucket-name

PORT=8080  # API only
```

## Architecture Benefits

- **Separation of Concerns**: API handles HTTP, worker handles scheduled tasks
- **Independent Scaling**: Scale API and worker separately based on load
- **Shared Code**: Types and services are reused across components
- **Easier Testing**: Test API and notification logic independently
- **Fault Isolation**: Worker crashes don't affect API availability

## Next Steps

The notification worker skeleton is created but not yet implemented. To complete:

1. Implement query for plants needing water in `cmd/notification-worker/main.go`
2. Fetch user notification tokens from MongoDB
3. Send FCM push notifications via Firebase service
4. Update last notification timestamp
5. Add error handling and retry logic
