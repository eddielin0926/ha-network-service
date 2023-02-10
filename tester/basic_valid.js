import http from 'k6/http';
import { check } from 'k6';
import { server } from './config.js';

export const options = {
    scenarios: {
        send: {
            executor: 'shared-iterations',
            vus: 1, // number of threads
            iterations: server.Valid_Num[server.TEST_CASE],
            maxDuration: '30s',
        },
    },
};


const Golden = JSON.parse(open(`./data/golden_${server.TEST_CASE}.json`));
let idx = 0;

function Check_Record(r, ans) {
    if (r.length != ans.length) return false
    for (let t = 0; t < r.length; ++t) {
        let find = false;
        for (let i = 0; i < ans.length; ++i) {
            let res = r[t].location == ans[i].location && r[t].timestamp == ans[i].timestamp
            res &= r[t].signature == ans[i].signature && r[t].material == ans[i].material
            res &= (r[t].a == ans[i].a && r[t].b == ans[i].b && r[t].c == ans[i].c && r[t].d == ans[i].d)
            if (res) { find = true; break; }
        };
        if (!find) return false;
    }
    return true
}


function Check_Report(r, ans) {
    let res = r.count == ans.count && r.material == ans.total
    res &= (r.a == ans.a && r.b == ans.b && r.c == ans.c && r.d == ans.d)
    return res
}

// 記得要修改 r 的東西，因為不知道回傳的東西名稱是什麼 !!!

export default function () {
    let ans = Golden[idx++]
    // const res1 = http.get(`https://www.google.com.tw?location=${ans.location}&date=${ans.timestamp}`)
    // const res2 = http.get(`https://www.google.com.tw?location=${ans.location}&date=${ans.timestamp}`)
    const res1 = http.get(`http://${server.host}:${server.port}/${server.record_endpoint}?location=${ans.location}&date=${ans.timestamp}`)
    const res2 = http.get(`http://${server.host}:${server.port}/${server.report_endpoint}?location=${ans.location}&date=${ans.timestamp}`)

    check(res1, {
        '> Record Status is 200': (r) => r.status === 200,
        '> Record is Correct': (r) => Check_Record(r, ans.order),
    });
    check(res2, {
        '> Report Status is 200': (r) => r.status === 200,
        '> Report is Correct': (r) => Check_Report(r, ans),
    });
}


