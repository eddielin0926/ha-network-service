import http from 'k6/http';
import { check } from 'k6';
import { server } from './config.js';

export const options = {
    scenarios: {
        send: {
            executor: 'per-vu-iterations',
            vus: server.VUs, // number of threads
            iterations: server.PerVUsIter,
            maxDuration: '30s',
        },
    },
};


const data = { "location": "l1", "timestamp": "2023-01-01T00:00:00.000+08:00", "data": { "a": 1, "b": 1, "c": 1, "d": 1 } }

export default function () {
    const res = http.post(`http://${server.host}:${server.port}/${server.send_endpoint}`, data, server.param)
    check(res, {
        'status is 200': (r) => r.status === 200,
    });
}
