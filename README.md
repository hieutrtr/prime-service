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

## Installation
You need go installed and `GOBIN` in your `PATH`

install `go install github.com/hieutrtr/prime-service`

start service `prime-service start --limit-prime <prime limit number>`

options
```
-limit-prime (default: 1000000) : Input that greater than limit-prime will not be supported. This configuration is also for the service to know how many prime numbers need to be generated in cache storage.
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
