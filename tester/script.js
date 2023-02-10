import http from 'k6/http';
import { check } from 'k6';
import { server } from './config.js';

export const options = {
  scenarios: {
    send: {
      executor: 'shared-iterations',
      vus: 10, // number of threads
      iterations: 200,
      maxDuration: '30s',
    },
  },
};

let params = { headers: { "accept": "application/json", "Content-Type": "application/json" } };
let body = { location: "l1", timestamp: "2023-01-01T20:18:57.424+08:00", data: { a: 1, b: 1, c: 1, d: 1 } };

export default function () {
  // const res = http.get('http://localhost:8100/api/record?location=l1&date=2023-01-01')
  const res = http.post('http://localhost:8100/api/order', JSON.stringify(body), params)
  // const res = http.get(`http://${server.host}:${server.port}/${server.send_endpoint}`);
  // console.log(res)
  check(res, {
    'Status is 200': (r) => r.status === 200,
  });
}

