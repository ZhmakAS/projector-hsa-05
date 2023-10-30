# projector-hsa-05

Endpoints:

- `/users POST` - insert single record with fake data

## How to launch

Execute `docker compose up -d` and that's all, after that service can be accessed on http://localhost:8081

## Run load tests:
```
docker run --rm -t --network host -v  $(pwd)/urls/urls.txt:/etc/urls/urls.txt yokogawa/siege -d2 -t5M -c300 -i -f /etc/urls/urls.txt
```


Urls for test:

```
http://localhost:8081/users POST
http://localhost:8081/users GET

```

## Results

`
-d2 -t5M -c300
`
```
Transactions:                  85943 hits
Availability:                  99.99 %
Elapsed time:                 299.98 secs
Data transferred:               1.07 MB
Response time:                  0.04 secs
Transaction rate:             286.50 trans/sec
Throughput:                     0.00 MB/sec
Concurrency:                   12.47
Successful transactions:       85943
Failed transactions:              10
Longest transaction:            0.25
Shortest transaction:           0.00

```

`
-d1 -t1 -c400
`

```
Transactions:                  34562 hits
Availability:                  96.20 %
Elapsed time:                  49.82 secs
Data transferred:               0.43 MB
Response time:                  0.05 secs
Transaction rate:             693.74 trans/sec
Throughput:                     0.01 MB/sec
Concurrency:                   35.29
Successful transactions:       34562
Failed transactions:            1367
Longest transaction:            0.42
Shortest transaction:           0.00
```

`
-d1 -t1 -c500
`
```
Transactions:                  33417 hits
Availability:                  95.79 %
Elapsed time:                  39.03 secs
Data transferred:               0.41 MB
Response time:                  0.05 secs
Transaction rate:             856.19 trans/sec
Throughput:                     0.01 MB/sec
Concurrency:                   44.82
Successful transactions:       33417
Failed transactions:            1468
Longest transaction:            0.32
Shortest transaction:           0.00
```
