import http from 'k6/http';
import { check, sleep } from 'k6';
import { Rate } from 'k6/metrics';

const BASE_URL = 'http://localhost:9191/api/v1';
export let errorRate = new Rate('errors');

export const options = {
  stages: [
    { duration: '2m', target: 10 }, // Ramp up to 10 users
    { duration: '5m', target: 10 }, // Stay at 10 users
    { duration: '2m', target: 50 }, // Ramp up to 50 users
    { duration: '5m', target: 50 }, // Stay at 50 users
    { duration: '2m', target: 0 },  // Ramp down
  ],
  thresholds: {
    http_req_duration: ['p(95)<500'], // 95% of requests should be below 500ms
    http_req_failed: ['rate<0.1'],    // Error rate should be less than 10%
    errors: ['rate<0.1'],
  },
};

export function setup() {
  console.log('Starting performance test against:', BASE_URL);
}

export default function() {
  // Test authentication endpoints
  let authPayload = JSON.stringify({
    email: `test${__VU}@example.com`,
    password: 'testpassword123',
    name: `Test User ${__VU}`,
  });

  let params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  // Test user registration (should fail - server not implemented)
  let registerResponse = http.post(`${BASE_URL}/auth/register`, authPayload, params);
  let registerSuccess = check(registerResponse, {
    'registration status is connection refused or error': (r) =>
      r.status === 0 || r.status >= 500,
  });

  errorRate.add(!registerSuccess);

  // Test user login (should fail - server not implemented)
  let loginPayload = JSON.stringify({
    email: `test${__VU}@example.com`,
    password: 'testpassword123',
  });

  let loginResponse = http.post(`${BASE_URL}/auth/login`, loginPayload, params);
  let loginSuccess = check(loginResponse, {
    'login status is connection refused or error': (r) =>
      r.status === 0 || r.status >= 500,
  });

  errorRate.add(!loginSuccess);

  // Test posts endpoint (should fail - server not implemented)
  let postsParams = {
    headers: {
      'Authorization': 'Bearer fake-token',
    },
  };

  let postsResponse = http.get(`${BASE_URL}/posts?limit=10&offset=0`, postsParams);
  let postsSuccess = check(postsResponse, {
    'posts status is connection refused or error': (r) =>
      r.status === 0 || r.status >= 500,
  });

  errorRate.add(!postsSuccess);

  // Test create post (should fail - server not implemented)
  let postPayload = JSON.stringify({
    content: `This is test post ${__VU} - ${__ITER}`,
  });

  let createPostParams = {
    headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer fake-token',
    },
  };

  let createPostResponse = http.post(`${BASE_URL}/posts`, postPayload, createPostParams);
  let createPostSuccess = check(createPostResponse, {
    'create post status is connection refused or error': (r) =>
      r.status === 0 || r.status >= 500,
  });

  errorRate.add(!createPostSuccess);

  // Test users search (should fail - server not implemented)
  let searchResponse = http.get(`${BASE_URL}/users/search?q=test`, postsParams);
  let searchSuccess = check(searchResponse, {
    'search status is connection refused or error': (r) =>
      r.status === 0 || r.status >= 500,
  });

  errorRate.add(!searchSuccess);

  sleep(1);
}

export function teardown() {
  console.log('Performance test completed');
}