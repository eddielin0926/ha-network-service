import http from 'k6/http';
import { check } from 'k6';
import { server } from './config.js';

export const options = {
    scenarios: {
        send: {
            executor: 'shared-iterations',
            vus: 1, // number of threads
            iterations: server.Case_Num[server.TEST_CASE],
            maxDuration: '30s',
        },
    },
};


const orders = JSON.parse(open(`./data/order_${server.TEST_CASE}.json`));
let idx = 0;

export default function () {
    const res = http.post(`http://${server.host}:${server.port}/${server.send_endpoint}`, JSON.stringify(orders[idx++]), server.param)
    check(res, {
        'status is 200': (r) => r.status === 200,
    });
}
