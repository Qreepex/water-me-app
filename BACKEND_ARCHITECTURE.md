# Backend Architecture

The Water Me backend has been refactored into two independent components:

## Components

### 1. API Server (`backend/cmd/api/`)
- **Purpose**: REST API for plant management
- **Port**: 8080
- **Endpoints**: Plant CRUD, authentication, uploads, notifications settings
- **Scaling**: Horizontal (multiple replicas)
- **Run**: `go run ./backend/cmd/api`

### 2. Notification Worker (`backend/cmd/notification-worker/`)
- **Purpose**: Send push notifications for watering reminders
- **Schedule**: Checks every 5 minutes
- **Endpoints**: None (background service)
- **Scaling**: Vertical (single instance recommended)
- **Run**: `go run ./backend/cmd/notification-worker`

## Shared Code

Both components share:
- `types/` - Data structures
- `services/` - Database, Firebase, S3, authentication
- `constants/` - Application constants
- `validation/` - Input validation
- `util/` - Helper functions

## Quick Start

### Development
```bash
# Terminal 1 - API Server
cd backend
go run ./cmd/api

# Terminal 2 - Notification Worker (optional)
cd backend
go run ./cmd/notification-worker
```

### Production Builds
```bash
cd backend

# Build both components
go build -o bin/api ./cmd/api
go build -o bin/notification-worker ./cmd/notification-worker

# Or use build scripts
./build.sh 0.0.6    # Linux/Mac
./build.ps1 0.0.6   # Windows
```

### Docker
```bash
# API Server
docker build --build-arg COMPONENT=api -t plants-backend-api .

# Notification Worker
docker build --build-arg COMPONENT=notification-worker -t plants-notification-worker .
```

### Kubernetes
```bash
kubectl apply -f backend/infra/k8s/namespace.yaml
kubectl apply -f backend/infra/k8s/secrets.yaml
kubectl apply -f backend/infra/k8s/api.yaml
```

## Architecture Benefits

✅ **Independent Scaling** - Scale API and worker separately  
✅ **Fault Isolation** - Worker crashes don't affect API  
✅ **Code Reuse** - Shared services and types  
✅ **Clear Separation** - HTTP logic vs background jobs  
✅ **Easier Testing** - Test components independently  

## Implementation Status

- ✅ API Server - Fully implemented
- 🚧 Notification Worker - Skeleton created, logic pending
  - TODO: Query plants needing water
  - TODO: Fetch user notification tokens
  - TODO: Send FCM push notifications
  - TODO: Update notification timestamps

See [backend/cmd/README.md](backend/cmd/README.md) for detailed documentation.
