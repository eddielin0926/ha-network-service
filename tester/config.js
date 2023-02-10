var server = {
    host: "172.20.10.4",
    port: 80,
    send_endpoint: "api/order",
    report_endpoint: "api/report",
    record_endpoint: "api/record",
    param: { headers: { "accept": "application/json", "Content-Type": "application/json" } },
    Case_Num: [0, 3, 200, 20000],
    Valid_Num: [0, 1, 30, 217],
    // Change Test Case Here
    TEST_CASE: 1,

    // Change Load Balance Testing Here
    VUs: 1000,
    PerVUsIter: 10,
    ts: "2023-01-05"
};

export {
    server
};