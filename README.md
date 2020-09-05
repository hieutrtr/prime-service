# Prime Service
Finding highest prime number lower than the input number.

## Description
Finding the highest prime number lower than input number N is a very simple assignment, but using it as an API service, 
we should care about support multiple requests which leading to have duplicated calculation of prime number.
So I initialize a kind of cache layer to support multiple requests and reduce duplicated calculation cost.

## Thought you wanna try first !

* Get secret
```
visit: https://onetimesecret.com/secret/n50wn0kid8izr6yekm12ogedahzzhgs
type passphase: WHTMJmhKXQ
```
Please keep the secret on your laptop to use for next time.

* Generate JWT
```
- Visit: http://jwt.io

- Use payload:
    {
      "iss": "stably",
      "exp": 1599185100 
    }
// please use exp as current timestamp + 1 min

- Place secret in "VERIFY SIGNATURE"
- Copy the JWT from the left box
```

* Try API via swagger
```
- Visit: http://34.126.71.65/swagger/index.html
- Click on the "POST /prime" endpoint.
- Click "Try it out"
- Try with input :
{
  "number": 500000
}
- Authorization : Bearer <place generated JWT>
- Click "Execute"
```

* Or Try via curl with default JWT
```
curl --location --request POST '34.126.71.65/api/v1/prime' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdGFibHkifQ.EbW4_ASIThuZkJyRUrrO6eFjymeisdcO3L2r0CS2LLM' \
--header 'Content-Type: application/json' \
--data-raw '{
    "number": 100000
}'
```

## Attack me
```
- Install vegeta: go get -u github.com/tsenart/vegeta
- Modify: attack/prime_endpoint to put JWT token.
- Run: cat attack/prime_endpoint | vegeta attack -duration=10s -rate=4000 -body attack/body.json | tee results.bin | vegeta report
```

## Supported versions

* Go version 1.15
* Echo framework v4.1.17

## Features overview
* Get the nearest prime less than or equal to number N input

## Quick Start
Use docker to start service on you local:

`docker run -d --name prime-service -p 8080:8080 hieutrtr/prime-service:latest`

Call API via curl with default JWT:
```
curl --location --request POST 'localhost:8080/api/v1/prime' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdGFibHkifQ.EbW4_ASIThuZkJyRUrrO6eFjymeisdcO3L2r0CS2LLM' \
--header 'Content-Type: application/json' \
--data-raw '{
    "number": 100000
}'
```

## Installation
You need go installed and `GOBIN` in your `PATH`

install `go install github.com/hieutrtr/prime-service`

start service `prime-service start --limit-prime <prime limit number>`

options
```
-limit-prime (default: 1000000) : Input that greater than limit-prime will not be supported. This configuration is also for the service to know how many prime numbers need to be generated in cache storage.
-jwt-secret (default: SECRET) : secret key to generate JWT using HMAC algorithm.
```

## Build and Deployment
Clone project:
`git clone https://github.com/hieutrtr/prime-service.git`

Run test

`go test ./...`

Build docker image:

`docker build -t <repo>/<image>:<version> .`

Run docker image:

`docker run -d -p <map port>:8080 --name <container name> <repo>/<image>:<version>`

Eg. `docker run -d -p 8080:8080 --name prime_api hieutrtr/prime-service:latest`

## Performance
#### Algorithm performance

A simple solution for finding highest prime lower than number N is to iterate from N-1 to 2, 
and for every number check if it's a prime. 
This solution look fine if there is only one query.
But not efficient if there're multiple queries for different values of N.

An efficient solution for this problem is to generate all primes using Sieve of Sundaram and store then in an array in increasing order. 
Then apply binary-search to search nearest prime less than N. Time complexity of this solution is O(n log n + log n) = O(n log n).

In thi service I will improve the above solution by keeping the marked array of [Sieve of Sundaram](https://www.geeksforgeeks.org/sieve-sundaram-print-primes-smaller-n/).
Tracing back the nearest prime by using Sieve marked array only cost O(1) because we only need to access array.

#### Service performance
By applying Sieve of Sundaram solution, we have a primes cache that can serve multiple requests with response time around 5ms for each.

Benchmarked by using [Vegeta](https://github.com/tsenart/vegeta)
 on my laptop Mac Pro, 2.3 GHz Intel Core i5 ( 4 cores ), 8GB of RAM:
  
* Rate 4000 in 10 seconds:
```
Requests      [total, rate, throughput]         40000, 4000.12, 4000.05
Duration      [total, attack, wait]             10s, 10s, 175.91µs
Latencies     [min, mean, 50, 90, 95, 99, max]  107.408µs, 1.806ms, 143.421µs, 3.747ms, 11.831ms, 30.774ms, 98.022ms
Bytes In      [total, mean]                     840000, 21.00
Bytes Out     [total, mean]                     840000, 21.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:40000
```

* Rate 4000 in 30 seconds: start having some fail requests.
```
Requests      [total, rate, throughput]         120000, 4000.11, 3914.98
Duration      [total, attack, wait]             29.999s, 29.999s, 237.194µs
Latencies     [min, mean, 50, 90, 95, 99, max]  21.577µs, 18.669ms, 1.949ms, 61.938ms, 70.96ms, 113.393ms, 288.78ms
Bytes In      [total, mean]                     2466387, 20.55
Bytes Out     [total, mean]                     2466387, 20.55
Success       [ratio]                           97.87%
Status Codes  [code:count]                      0:2553  200:117447
```

* Rate 5000 in 10 seconds: At 5000 requests per sec, 
it reach the CPU overhead leading to failures of constructing requests.
You can see the status codes has 6994 of zero status.
```
Requests      [total, rate, throughput]         50000, 5000.12, 4276.84
Duration      [total, attack, wait]             10.056s, 10s, 55.795ms
Latencies     [min, mean, 50, 90, 95, 99, max]  24.934µs, 48.614ms, 52.225ms, 74.711ms, 100.483ms, 136.052ms, 246.28ms
Bytes In      [total, mean]                     903126, 18.06
Bytes Out     [total, mean]                     903126, 18.06
Success       [ratio]                           86.01%
Status Codes  [code:count]                      0:6994  200:43006                    200:40000
```

## Security
This service authenticate and authorize requests use JWT with HS256 (HMAC) algorithm.
