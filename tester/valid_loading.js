import http from 'k6/http';
import { check } from 'k6';
import { server } from './config.js';

export const options = {
    scenarios: {
        send: {
            executor: 'shared-iterations',
            vus: 1, // number of threads
            iterations: 1,
            maxDuration: '30s',
        },
    },
};

const TotalCount = server.VUs * server.PerVUsIter

function Check_Record(r) {
    if (r.length != TotalCount) return false
    for (let t = 0; t < r.length; ++t) {
        let res = r[t].location == "l1" && r[t].timestamp == "2023-01-01T00:00:00.000+08:00"
        res &= r[t].signature == "MTk=" && r[t].material == 19
        res &= (r[t].a == 1 && r[t].b == 1 && r[t].c == 1 && r[t].d == 1)
        if (!res) return false;
    }
    return true
}


function Check_Report(r) {
    let res = r.count == ans.TotalCount && r.material == TotalCount * 19
    res &= (r.a == TotalCount && r.b == TotalCount && r.c == TotalCount && r.d == TotalCount)
    return res
}

// 記得要修改 r 的東西，因為不知道回傳的東西名稱是什麼 !!!

export default function () {
    // const res1 = http.get(`https://www.google.com.tw?`)
    // const res2 = http.get(`https://www.google.com.tw?`)
    const res1 = http.get(`http://${server.host}:${server.port}/${server.record_endpoint}?location=l1&date=2023-01-01`)
    const res2 = http.get(`http://${server.host}:${server.port}/${server.report_endpoint}?location=l1&date=2023-01-01`)

    check(res1, {
        '> Record Status is 200': (r) => r.status === 200,
        '> Record is Correct': (r) => Check_Record(JSON.parse(r)),
    });
    check(res2, {
        '> Report Status is 200': (r) => r.status === 200,
        '> Report is Correct': (r) => Check_Report(JSON.parse(r)),
    });
}
