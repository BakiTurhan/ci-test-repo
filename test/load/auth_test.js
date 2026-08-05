import http from 'k6/http';
import { check, sleep } from 'k6';

// Asıl projede olduğu gibi dışarıdan ortam değişkeni okuyoruz
const BASE_URL = __ENV.API_URL || 'http://localhost:3000/api/v1';

export const options = {
  vus: 1,
  duration: '1s',
};

export default function () {
  const payload = JSON.stringify({});
  const params = { headers: { 'Content-Type': 'application/json' } };
  
  // Dışarıdan gelen IP adresiyle istek atıyoruz
  const res = http.post(`${BASE_URL}/auth/register`, payload, params);
  
  check(res, {
    'register is successful (201)': (r) => r.status === 201,
  });
  
  sleep(1);
}
