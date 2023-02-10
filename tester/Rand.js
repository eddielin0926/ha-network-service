const fs = require('fs')

function randInt(min, max) {
    let interval = max - min + 1
    return Math.floor(Math.random() * interval) + min
}

function ToStr(int) {
    return (int < 10 ? '0' + int : int.toString())
}

function randOrder(maxLoc = 3, maxDate = 31) {
    let location = 'l' + randInt(1, maxLoc);
    let randDay = ToStr(randInt(1, maxDate));
    let randHr = ToStr(randInt(0, 23));
    let randMn = ToStr(randInt(0, 59));
    let randSe = ToStr(randInt(0, 59));
    let randMi = ToStr(randInt(0, 999));
    let timestamp = '2023-01-' + randDay + 'T' + randHr + ':' + randMn + ':' + randSe + '.' + randMi + '+08:00';
    let data = { 'a': randInt(0, 250), 'b': randInt(0, 250), 'c': randInt(0, 250), 'd': randInt(0, 250) }

    return { 'location': location, 'timestamp': timestamp, 'data': data }
}

function Add(d1, d2) {
    return { 'a': d1.a + d2.a, 'b': d1.b + d2.b, 'c': d1.c + d2.c, 'd': d1.d + d2.d }
}

function Get_record(order) {
    let record = { 'location': "", 'timestamp': "", 'signature': "", 'material': 0, 'a': 0, 'b': 0, 'c': 0, 'd': 0 }
    record.location = order.location
    record.timestamp = order.timestamp
    record.a = order.data.a
    record.b = order.data.b
    record.c = order.data.c
    record.d = order.data.d
    record.material = record.a * 3 + record.b * 2 + record.c * 4 + record.d * 10
    let total = record.a + record.b + record.c + record.d
    // base64.b64encode(str(total).encode('UTF-8')).decode('UTF-8')
    record.signature = btoa(String(total))
    return record
}

function Total_for_each(Orders, loc, date) {

    let Result = []
    for (let i = 0; i < Orders.length; ++i) {
        let loc = Orders[i].location
        let ts = Orders[i].timestamp.substring(0, 10)
        let data = Orders[i].data
        let find = false
        let record = Get_record(Orders[i])
        for (let t = 0; t < Result.length; ++t) {
            if (Result[t].location == loc && Result[t].timestamp == ts) {
                find = true
                Result[t].count++;
                Result[t].order.push(record)
                Result[t].total += record.material
                Result[t].a += data.a, Result[t].b += data.b, Result[t].c += data.c, Result[t].d += data.d
                break;
            }
        }
        if (!find) {
            Result.push({ 'location': loc, 'timestamp': ts, 'count': 1, 'total': record.material, 'a': data.a, 'b': data.b, 'c': data.c, 'd': data.d, 'order': [record] })
        }
    }
    return Result
}

function GenOrderFile(orderNum, loc, date) {
    let Orders = []
    while (orderNum--)
        Orders.push(randOrder(loc, date)) // change location & data count here
    return Orders
}

let TotalOrder = 3, MaxLoc = 1, MaxDate = 1
let Orders = GenOrderFile(TotalOrder, MaxLoc, MaxDate)
let Golden = Total_for_each(Orders, MaxLoc, MaxDate)
let Case = 1

fs.writeFile('./Data/order_' + String(Case) + '.json', JSON.stringify(Orders), 'utf8', (err) => { if (err) console.log(err) })
fs.writeFile('./Data/golden_' + String(Case) + '.json', JSON.stringify(Golden), 'utf8', (err) => { if (err) console.log(err) })
