# Quickstart Guide: Social Media REST API

This guide walks through setting up and testing the social media API from scratch.

## Prerequisites

- Go 1.21+ installed
- PostgreSQL 14+ installed and running
- Git for cloning the repository

## Setup

### 1. Clone and Navigate
```bash
git clone <repository-url>
cd potrit
```

### 2. Database Setup
```bash
# Create database
createdb potrit

# Run migrations (will be created during implementation)
psql potrit -f migrations/001_initial_schema.sql
```

### 3. Configure Application
Copy and edit the configuration:
```bash
cp config.json.example config.json
# Edit config.json with your database credentials
```

### 4. Build and Run
```bash
go mod tidy
go run cmd/server/main.go
```

The API will be available at `http://localhost:9191`

## Quick Test Walkthrough

This walkthrough demonstrates the core user journey through the API.

### 1. User Registration
```bash
curl -X POST http://localhost:9191/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "alice@example.com",
    "password": "password123",
    "name": "Alice Johnson"
  }'
```

Expected response:
```json
{
  "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "email": "alice@example.com",
    "name": "Alice Johnson",
    "created_at": "2025-10-05T10:00:00Z",
    "updated_at": "2025-10-05T10:00:00Z"
  }
}
```

### 2. User Login
```bash
curl -X POST http://localhost:9191/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "alice@example.com",
    "password": "password123"
  }'
```

### 3. Create a Post
```bash
# Use the token from login
TOKEN="eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9..."

curl -X POST http://localhost:9191/api/v1/posts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "content": "Hello world! This is my first post on this social media platform."
  }'
```

Expected response:
```json
{
  "id": "660e8400-e29b-41d4-a716-446655440001",
  "user_id": "550e8400-e29b-41d4-a716-446655440000",
  "content": "Hello world! This is my first post on this social media platform.",
  "created_at": "2025-10-05T10:01:00Z",
  "updated_at": "2025-10-05T10:01:00Z",
  "likes_count": 0,
  "comments_count": 0,
  "user": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "email": "alice@example.com",
    "name": "Alice Johnson"
  }
}
```

### 4. Register Second User
```bash
curl -X POST http://localhost:9191/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "bob@example.com",
    "password": "password456",
    "name": "Bob Smith"
  }'
```

### 5. Create Connection Request
```bash
# Login as Bob and get token
BOB_TOKEN=$(curl -s -X POST http://localhost:9191/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "bob@example.com", "password": "password456"}' | \
  jq -r '.token')

# Get Alice's user ID
ALICE_ID="550e8400-e29b-41d4-a716-446655440000"

# Send connection request
curl -X POST http://localhost:9191/api/v1/connections \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $BOB_TOKEN" \
  -d '{
    "addressee_id": "'$ALICE_ID'"
  }'
```

### 6. Accept Connection Request
```bash
# Login as Alice
ALICE_TOKEN=$(curl -s -X POST http://localhost:9191/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "alice@example.com", "password": "password123"}' | \
  jq -r '.token')

# Get connection ID from Bob's request (would be returned in step 5)
CONNECTION_ID="770e8400-e29b-41d4-a716-446655440002"

# Accept the connection
curl -X PATCH http://localhost:9191/api/v1/connections/$CONNECTION_ID \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ALICE_TOKEN" \
  -d '{
    "status": "accepted"
  }'
```

### 7. View Feed (Now includes Bob's posts for Alice)
```bash
curl -X GET http://localhost:9191/api/v1/posts \
  -H "Authorization: Bearer $ALICE_TOKEN"
```

### 8. Like a Post
```bash
POST_ID="660e8400-e29b-41d4-a716-446655440001"

curl -X POST http://localhost:9191/api/v1/posts/$POST_ID/like \
  -H "Authorization: Bearer $ALICE_TOKEN"
```

### 9. Add Comment
```bash
curl -X POST http://localhost:9191/api/v1/posts/$POST_ID/comments \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ALICE_TOKEN" \
  -d '{
    "content": "Great post! Looking forward to seeing more content."
  }'
```

### 10. Get Post Details (with likes and comments)
```bash
curl -X GET http://localhost:9191/api/v1/posts/$POST_ID \
  -H "Authorization: Bearer $ALICE_TOKEN"
```

## Validation Steps

Use these commands to verify the implementation is working correctly:

### 1. Health Check
```bash
curl http://localhost:9191/health
# Should return: {"status": "ok"}
```

### 2. Database Connection
```bash
curl http://localhost:9191/api/v1/health/db
# Should return: {"status": "connected", "database": "potrit"}
```

### 3. API Documentation
```bash
curl http://localhost:9191/api/v1/docs
# Should return OpenAPI specification
```

### 4. Performance Test (basic)
```bash
# Install k6 if not present
# npm install -g k6

# Run basic performance test
k6 run tests/performance/basic_load_test.js
```

## Common Issues

### Database Connection Failed
- Verify PostgreSQL is running: `pg_ctl status`
- Check database exists: `psql -l`
- Verify config.json credentials

### Authentication Fails
- Check JWT secret is properly configured
- Verify token is not expired
- Ensure Bearer token format is correct

### CORS Errors
- For frontend integration, ensure CORS middleware is configured
- Check Origin header is allowed

### Performance Issues
- Monitor database connection pool usage
- Check for slow queries with `EXPLAIN ANALYZE`
- Verify k6 performance test results meet targets

## Next Steps

1. **Run Contract Tests**: Execute all contract tests to verify API compliance
2. **Performance Testing**: Run full k6 test suite
3. **Integration Testing**: Test with actual frontend application
4. **Monitoring Setup**: Configure logging and metrics collection

## Support

For issues or questions:
- Check the implementation logs for error details
- Verify database schema matches expectations
- Review API contract documentation
- Run individual unit tests for specific functionality