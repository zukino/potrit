# Performance Tests

This directory contains k6 performance tests for the Social Media REST API.

## Prerequisites

Install k6: https://k6.io/docs/getting-started/installation/

## Available Tests

### Basic Load Test
Tests core functionality under moderate load:
- User registration
- User login
- Posts retrieval
- Post creation
- User search

Run with:
```bash
k6 run tests/performance/scripts/basic_load_test.js
```

### Authentication Stress Test
Tests authentication endpoints under high load:
- Registration endpoint stress testing
- Login endpoint stress testing

Run with:
```bash
k6 run tests/performance/scripts/auth_stress_test.js
```

## Expected Behavior

**Currently**, all tests are designed to **fail** with connection errors since the server implementation doesn't exist yet. This is intentional and follows the TDD methodology.

The tests check for:
- Connection refused errors (status 0)
- Server errors (status 5xx)

**After implementation**, these tests should be updated to expect successful responses (2xx status codes).

## Configuration

All tests target: `http://localhost:9191/api/v1`

## Performance Targets

- 95th percentile response time < 500ms
- Error rate < 10%
- Support up to 100 concurrent users (MVP target)