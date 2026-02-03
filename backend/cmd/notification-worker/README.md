# Notification Worker

Production-ready background service that sends push notifications to users when their plants need care.

## Quick Start

### Local Development

```bash
cd backend
go run ./cmd/notification-worker
```

### Production Build

```bash
go build -o bin/notification-worker ./cmd/notification-worker
./bin/notification-worker
```

### Docker

```bash
docker build --build-arg COMPONENT=notification-worker -t plants-notification-worker:latest .
docker run plants-notification-worker:latest
```

## Configuration

### Environment Variables

```bash
DATABASE_URL=mongodb://localhost:27017/plants
MONGODB_USERNAME=username
MONGODB_PASSWORD=password
MONGODB_DATABASE=plants
GOOGLE_APPLICATION_CREDENTIALS=./secret/fb.json
```

### Message Templates

The worker loads notification message variants from `messages/` directory:

- `watering.json` - Water reminders (4 single + 4 multiple variants)
- `fertilizing.json` - Fertilizer reminders (4 single + 4 multiple variants)
- `misting.json` - Misting reminders (4 single + 4 multiple variants)
- `repotting.json` - Repotting reminders (4 single + 4 multiple variants)

## How It Works

### Main Loop

1. **Startup**: Loads message templates, connects to database and Firebase
2. **Immediate Check**: Runs notification check on startup
3. **Scheduled Checks**: Runs every 5 minutes thereafter
4. **Timeout**: 10-minute timeout per check for large batches

### Notification Check Process

1. **Query Plants**: Fetch plants needing care (1000 per batch)
   - Watering: `lastWatered + intervalDays <= now`
   - Fertilizing: `lastFertilized + intervalDays <= now`
   - Misting: `lastMisted + intervalDays <= now`
   - Repotting: `lastRepotted + repottingCycle <= now`

2. **Group by User**: Plants grouped by userId for efficient processing

3. **For Each User**:
   - Fetch notification config
   - Check if notifications enabled
   - Check notification type is enabled (watering, fertilizing, etc.)
   - Check 4-hour cooldown period
   - Filter muted plants
   - Get active device tokens
   - Build message from random template variant
   - Send to FCM in batches of 500
   - Handle failed tokens
   - Update last notification timestamp

4. **Log Results**: Duration, counts, success rate

## Scalability

### Performance Characteristics

- **50,000 users with 100 plants each** = 5,000,000 plants
- **10% need water daily** = 500,000 plants
- **Execution time**: ~2 minutes per check
- **Memory usage**: ~200MB per worker
- **Database queries**: ~600 per check

### Batch Sizes

- `plantsBatchSize = 1000` - Plants fetched per query
- `FCMBatchSize = 500` - FCM tokens per multicast (Firebase limit)
- `NotificationCooldown = 4 hours` - Minimum time between notifications per user

### Optimization

- MongoDB aggregation with date calculations
- Indexes on date fields for fast queries
- In-memory grouping by user
- Parallel processing of notification types
- Random template selection for variety

## Failed Token Handling

When FCM tokens fail (expired, unregistered, invalid):

1. **Detection**: `SendMulticastNotification()` returns batch response with failures
2. **Extraction**: `extractFailedTokens()` identifies which tokens failed
3. **Logging**: Failed tokens logged with error details (truncated for security)
4. **Database Update**: `MarkTokensAsInactive()` sets `isActive: false`
5. **Future Skipping**: Inactive tokens skipped in subsequent checks
6. **User Recovery**: User must re-register token via API

### Why This Matters

- **Reduces FCM quota usage**: Don't waste sends on invalid tokens
- **Improves success rate**: Only send to valid devices
- **Faster execution**: Skip known-bad tokens
- **Better metrics**: Accurate success/failure tracking

## Message Variants

### Benefits

- **Reduced monotony**: Users don't see same message repeatedly
- **Better engagement**: Variety keeps notifications interesting
- **A/B testing**: Can track which variants perform better
- **Localization ready**: Easy to add language-specific variants

### Adding New Variants

1. Edit JSON file in `messages/` directory
2. Add to `single` or `multiple` array
3. Use template variables: `{plantName}`, `{count}`, `{plantNames}`, `{remaining}`
4. Restart worker to reload

Example:

```json
{
  "single": [
    {
      "title": "New creative title! 💧",
      "body": "{plantName} would love some water"
    }
  ]
}
```

## Monitoring

### Expected Logs

```
========================================
Notification Worker Started
========================================
Scalable for 50k+ users with 100+ plants each
Configuration:
  - Plants batch size: 1000
  - FCM batch size: 500
  - Check interval: 5m0s
  - Cooldown period: 4h0m0s
========================================
Loaded 4 single and 4 multiple messages for watering
Loaded 4 single and 4 multiple messages for fertilizing
Loaded 4 single and 4 multiple messages for misting
Loaded 4 single and 4 multiple messages for repotting
========== Starting notification check ==========
Found 234 plants needing watering
Sent watering notifications to user abc123 for 5 plants
Sent watering notifications to user def456 for 3 plants
Token failed: eyJhbGc..., Error: messaging/registration-token-not-registered
Handling 2 failed tokens for user ghi789
========== Notification check complete ==========
Duration: 1m23s
Plants checked: 234
Notifications sent: 89
Notifications failed: 2
Users notified: 45
Success rate: 97.80%
================================================
```

### Health Checks

- Worker should complete checks in < 5 minutes
- Success rate should be > 95%
- Failed tokens should decrease over time (cleanup working)
- Check interval of 5 minutes is consistent

### Troubleshooting

**Worker not sending notifications:**

- Check Firebase credentials are valid
- Verify MongoDB connection
- Ensure message files exist in `messages/` directory
- Check user has notifications enabled
- Verify plants have intervals configured

**High failure rate:**

- Many users haven't used app recently (tokens expired)
- Need token cleanup campaign
- Check FCM quota limits
- Verify Firebase project configuration

**Slow execution:**

- Increase batch sizes
- Add database indexes
- Check MongoDB performance
- Consider horizontal scaling

**Messages not varying:**

- Check message JSON files are loaded
- Verify random selection is working
- Ensure multiple variants exist

## Architecture

### Code Organization

- **`cmd/notification-worker/main.go`**: Minimal entry point (117 lines)
  - Config loading
  - Service initialization
  - Main loop
  - Logging

- **`services/notification_worker.go`**: Core logic (400+ lines)
  - Message template loading
  - Notification processing
  - User filtering
  - Batch sending
  - Token handling
  - Message building

- **`services/database.go`**: Data access
  - `GetPlantsNeedingWatering()`
  - `GetPlantsNeedingFertilizer()`
  - `GetPlantsNeedingMisting()`
  - `GetPlantsNeedingRepotting()`
  - `GetNotificationConfig()`
  - `UpdateNotificationLastSent()`
  - `MarkTokensAsInactive()`

- **`services/firebase.go`**: FCM integration
  - `SendNotification()` - Single device
  - `SendMulticastNotification()` - Up to 500 devices

### Why This Structure?

- **Separation of concerns**: Main is simple, logic is testable
- **Reusability**: Services can be used by other components
- **Maintainability**: Clear boundaries between layers
- **Testability**: Each service can be unit tested
- **Scalability**: Easy to add new notification types

## Future Enhancements

1. **Preferred Time Support**: Send at user's preferred time
2. **Quiet Hours**: Skip notifications during quiet hours
3. **Batching Days**: Group notifications by user preference
4. **Smart Scheduling**: ML-based optimal timing
5. **Priority Queue**: Urgent notifications sent immediately
6. **Localization**: Multi-language message variants
7. **Rich Notifications**: Images, actions, deep links
8. **Analytics**: Track open rates, click-through rates
