import http from 'k6/http';
import { check, sleep } from 'k6';
import { Rate } from 'k6/metrics';

const BASE_URL = 'http://localhost:9191/api/v1';
export let errorRate = new Rate('errors');

export const options = {
  stages: [
    { duration: '1m', target: 20 }, // Ramp up to 20 users quickly
    { duration: '3m', target: 20 }, // Stay at 20 users
    { duration: '1m', target: 100 }, // Ramp up to 100 users
    { duration: '3m', target: 100 }, // Stay at 100 users
    { duration: '1m', target: 0 },  // Ramp down
  ],
  thresholds: {
    http_req_duration: ['p(95)<300'], // Auth endpoints should be fast
    http_req_failed: ['rate<0.1'],
    errors: ['rate<0.1'],
  },
};

export default function() {
  let userEmail = `stresstest${__VU}${__ITER}@example.com`;

  // Test registration endpoint load
  let registerPayload = JSON.stringify({
    email: userEmail,
    password: 'testpassword123',
    name: `Stress Test User ${__VU}`,
  });

  let params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  let registerResponse = http.post(`${BASE_URL}/auth/register`, registerPayload, params);
  let registerSuccess = check(registerResponse, {
    'registration handles load': (r) => r.status === 0 || r.status >= 500,
  });

  errorRate.add(!registerSuccess);

  // Test login endpoint load
  let loginPayload = JSON.stringify({
    email: userEmail,
    password: 'testpassword123',
  });

  let loginResponse = http.post(`${BASE_URL}/auth/login`, loginPayload, params);
  let loginSuccess = check(loginResponse, {
    'login handles load': (r) => r.status === 0 || r.status >= 500,
  });

  errorRate.add(!loginSuccess);

  sleep(0.5); // Shorter sleep for stress test
}