var server = {
    host: "127.0.0.1",
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
    VUs: 100,
    PerVUsIter: 100,
    ts: "2023-01-02"
};

export {
    server
};