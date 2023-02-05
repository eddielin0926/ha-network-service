import http from 'k6/http';
import { sleep } from 'k6';

const baseUrl = " http://localhost:8080/api";

export const options = {
    vus: 10,
    duration: '30s',
  };

export default function () {
  http.get(baseUrl);
  sleep(1);
}
