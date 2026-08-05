import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  vus: 1,
  duration: '1s',
};

export default function () {
  const payload = JSON.stringify({});
  const params = { headers: { 'Content-Type': 'application/json' } };
  
  // Ayağa kalkan sahte 3000 portumuza istek atıyoruz
  const res = http.post('http://localhost:3000/api/v1/auth/register', payload, params);
  
  check(res, {
    'register is successful (201)': (r) => r.status === 201,
  });
  
  sleep(1);
}
