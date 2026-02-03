# Notification System - Production-Ready Architecture

The Water Me notification system is designed to handle 50,000+ users with 100+ plants each (5+ million plants total).

## Architecture Overview

### Components

1. **API Server** (`cmd/api/`) - Manages device token registration
2. **Notification Worker** (`cmd/notification-worker/`) - Background service that sends push notifications
3. **Notification Service** (`services/notification_worker.go`) - Core notification logic
4. **Firebase Cloud Messaging (FCM)** - Delivery infrastructure
5. **MongoDB** - Stores notification configs and device tokens
6. **Message Templates** (`messages/`) - JSON files with notification variants

### File Structure

```
backend/
├── cmd/
│   ├── api/main.go                    # API server entry point
│   └── notification-worker/main.go    # Worker entry point (minimal, 117 lines)
├── services/
│   ├── notification_worker.go         # Core notification logic (400+ lines)
│   ├── firebase.go                    # FCM integration
│   ├── database.go                    # Database queries + MarkTokensAsInactive
│   └── mongo.go                       # MongoDB connection
├── messages/                          # Notification message variants
│   ├── watering.json                  # 4 single + 4 multiple variants
│   ├── fertilizing.json               # 4 single + 4 multiple variants
│   ├── misting.json                   # 4 single + 4 multiple variants
│   └── repotting.json                 # 4 single + 4 multiple variants
├── types/
│   └── notifications.go               # DeviceToken, NotificationConfig types
└── routes/
    └── notifications.go               # API endpoints for token management
```

### Scalability Features

#### Batch Processing

- **Plants Query**: Processes 1,000 plants per batch
- **FCM Multicast**: Sends to 500 devices per batch (FCM limit)
- **Parallel Processing**: Handles all notification types concurrently

#### Efficient Queries

- MongoDB aggregation pipeline with date calculations
- Indexes on `userId`, `watering.lastWatered`, `fertilizing.lastFertilized`
- Limits on result sets to prevent memory overflow

#### Rate Limiting

- 4-hour cooldown between notifications per user
- Prevents notification spam
- Respects quiet hours and preferred times (future enhancement)

#### Failed Token Handling

- **Automatic Detection**: Identifies invalid/expired FCM tokens
- **Database Update**: Marks failed tokens as `isActive: false`
- **Retry Prevention**: Skips inactive tokens in future sends
- **Cleanup**: Failed tokens never attempted again until user re-registers

#### Message Variety

- **4 variants per notification type** (single plant)
- **4 variants per notification type** (multiple plants)
- **Random selection** from variants to avoid repetitive messaging
- **Template variables**: `{plantName}`, `{count}`, `{plantNames}`, `{remaining}`

#### Parallelization

- Processes watering, fertilizing, misting, repotting in parallel
- Groups plants by user to minimize database queries
- Batch sends to multiple devices simultaneously

## Message Templates

Notification messages are stored in JSON files (`messages/*.json`) with multiple variants:

### Example: watering.json

```json
{
  "single": [
    {
      "title": "Time to water your plant! 💧",
      "body": "{plantName} needs water"
    },
    {
      "title": "Your plant is thirsty! 💧",
      "body": "{plantName} is ready for watering"
    }
  ],
  "multiple": [
    {
      "title": "{count} plants need water! 💧",
      "body": "{plantNames}"
    },
    {
      "title": "Watering time! 💧",
      "body": "{count} of your plants are thirsty: {plantNames}"
    }
  ]
}
```

### Template Variables

- `{plantName}` - Single plant name
- `{count}` - Number of plants
- `{plantNames}` - Comma-separated plant names (max 3 shown)
- `{remaining}` - Count of plants not shown in name list

### Adding New Messages

1. Edit the appropriate JSON file in `messages/`
2. Add new variants to `single` or `multiple` arrays
3. Restart notification worker to load new messages
4. Worker randomly selects from available variants

## Data Flow

```
┌─────────────────────────────────────────────────────────────┐
│ 1. Mobile App registers FCM token via API                  │
│    POST /api/notifications/tokens                          │
└──────────────────────┬──────────────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────────────┐
│ 2. Token stored in MongoDB notifications collection        │
│    { userId, deviceTokens: [{ token, deviceId, ... }] }    │
└──────────────────────┬──────────────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────────────┐
│ 3. Notification Worker runs every 5 minutes                │
│    - Queries plants needing care (1000 at a time)          │
│    - Groups plants by user                                  │
│    - Fetches notification configs                          │
└──────────────────────┬──────────────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────────────┐
│ 4. For each user:                                           │
│    - Check if notifications enabled                         │
│    - Filter muted plants                                    │
│    - Check cooldown period                                  │
│    - Build notification message                             │
└──────────────────────┬──────────────────────────────────────┘
                       │
                       ▼
┌─────────────────────────────────────────────────────────────┐
│ 5. Send via FCM multicast (up to 500 tokens per call)     │
│    - Handles failures gracefully                           │
│    - Marks invalid tokens as inactive                       │
│    - Updates lastNotificationSentAt timestamp              │
└─────────────────────────────────────────────────────────────┘
```

## API Endpoints

### Register Device Token

```http
POST /api/notifications/tokens
Authorization: Bearer <firebase-id-token>
Content-Type: application/json

{
  "token": "FCM-TOKEN-STRING",
  "deviceId": "unique-device-id",
  "deviceType": "android" // or "ios", "web"
}
```

**Response**: Updated notification config with device token added

### Remove Device Token

```http
DELETE /api/notifications/tokens/{deviceId}
Authorization: Bearer <firebase-id-token>
```

**Response**: Updated notification config with device token removed

### Get Notification Config

```http
GET /api/notifications
Authorization: Bearer <firebase-id-token>
```

**Response**: User's notification configuration including device tokens

## Database Schema

### Notification Config

```typescript
{
  _id: string,
  userId: string,
  deviceTokens: [
    {
      token: string,          // FCM token
      deviceId: string,       // Unique device identifier
      deviceType: string,     // "android", "ios", "web"
      addedAt: Date,
      lastUsedAt: Date,
      isActive: boolean       // False if token becomes invalid
    }
  ],
  isEnabled: boolean,
  preferredTime: string,      // "08:00"
  quietHours: {
    start: string,            // "22:00"
    end: string               // "07:00"
  },
  mutedPlantIds: [string],
  remindWatering: boolean,
  remindFertilize: boolean,
  remindRepotting: boolean,
  remindMisting: boolean,
  lastNotificationSentAt: Date,
  updatedAt: Date
}
```

## Performance Characteristics

### For 50,000 Users with 100 Plants Each

**Total Plants**: 5,000,000
**Plants Needing Care Daily**: ~500,000 (assuming 10% need water each day)

**Worker Execution**:

- Query plants: ~30 seconds (1000/batch × 500 batches)
- Group by user: ~5 seconds (in-memory operation)
- Fetch notification configs: ~10 seconds (500/batch × 100 batches)
- Send notifications: ~60 seconds (500 FCM tokens/batch × 100 batches)
- **Total**: ~2 minutes per check

**Memory Usage**: ~200MB per worker instance
**Database Load**: 600 queries per check (manageable with indexes)
**FCM Rate Limits**: Well within limits (600,000 messages/minute per project)

## Scaling Options

### Horizontal Scaling

- **Multiple Workers**: Run multiple notification worker instances
- **User Sharding**: Partition users across workers by userId hash
- **Geographic Distribution**: Deploy workers in multiple regions

### Vertical Scaling

- Increase batch sizes (plantsBatchSize, configBatchSize)
- Add more RAM for larger in-memory processing
- Use faster database connections

### Optimization Techniques

- **Database Indexes**:

  ```javascript
  db.plants.createIndex({ userId: 1, "watering.lastWatered": 1 });
  db.plants.createIndex({ "watering.intervalDays": 1 });
  db.notifications.createIndex({ userId: 1 });
  ```

- **Redis Caching**: Cache frequently accessed notification configs
- **Message Queue**: Use RabbitMQ/SQS for async notification delivery
- **FCM Topics**: Group users by notification preferences for batch sends

## Monitoring

### Key Metrics

- Plants checked per run
- Notifications sent per run
- Notification failures per run
- Users notified per run
- Average execution time
- FCM success rate
- Invalid token rate

### Alerts

- Worker execution time > 5 minutes
- Notification failure rate > 5%
- Worker crash/restart
- Database query timeout

## Error Handling

#### Invalid FCM Tokens

- **Detection**: Failed FCM sends are detected in batch response
- **Action**: `MarkTokensAsInactive()` updates database
- **Effect**: Marked as `isActive: false` in `deviceTokens` array
- **Recovery**: User must re-register token via API

#### Database Failures

- Worker retries next interval (5 minutes)
- Logs errors for debugging
- Continues processing other users

#### FCM Rate Limits

- Respects 500 tokens per multicast
- Implements exponential backoff if needed
- Queues notifications for retry

#### Message Loading Failures

- Worker fails to start if messages can't be loaded
- Ensures message files exist and are valid JSON
- Falls back to default messages if cache is empty

## Production Deployment

### Prerequisites

```bash
# Ensure message files exist
ls backend/messages/
# Should show: watering.json, fertilizing.json, misting.json, repotting.json
```

### Running Locally

```bash
cd backend
go run ./cmd/notification-worker
```

### Docker Build

```bash
# Build worker image
docker build --build-arg COMPONENT=notification-worker \
  -t plants-notification-worker:latest .
```

### Kubernetes Deployment

The worker deployment is already configured in `infra/k8s/api.yaml`:

- **Replicas**: 1 (recommended for single worker)
- **Resources**: 64Mi-256Mi RAM, 100m-250m CPU
- **Volume**: Firebase secret mounted at `/etc/firebase`

## Monitoring & Logging

### Key Logs

```
========== Starting notification check ==========
Found X plants needing watering
Sent watering notifications to user USER_ID for X plants
Token failed: FCM-TOKEN..., Error: messaging/registration-token-not-registered
Handling X failed tokens for user USER_ID
========== Notification check complete ==========
Duration: 30s
Plants checked: 1234
Notifications sent: 567
Notifications failed: 12
Users notified: 234
Success rate: 97.93%
================================================
```

### Metrics to Track

- **Plants checked per run** - Should match expected load
- **Notifications sent per run** - Indicates active users
- **Notification failures per run** - Track token expiration rate
- **Users notified per run** - Active user engagement
- **Average execution time** - Performance indicator
- **FCM success rate** - Token health indicator
- **Invalid token rate** - Cleanup effectiveness

### Alerts

- Worker execution time > 5 minutes → Performance degradation
- Notification failure rate > 5% → Token hygiene issues
- Worker crash/restart → Infrastructure problem
- Database query timeout → Database performance issue
- Message loading failure → Deployment issue

## Future Enhancements

1. **Preferred Time Support**: Send notifications at user's preferred time
2. **Quiet Hours**: Skip notifications during quiet hours
3. **Batching Days**: Group notifications by user preference (daily, every 2 days)
4. **Smart Scheduling**: ML-based optimal notification timing
5. **Priority Queue**: Urgent notifications (e.g., critical water needs) sent immediately
6. **A/B Testing**: Optimize notification content for user engagement
