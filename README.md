# Prime Service
Finding highest prime number lower than the input number.

## Description
Finding the highest prime number lower than input number N is a very simple assignment, but using it as an API service, 
we should care about support multiple requests which leading to have duplicated calculation of prime number.
So I initialize a kind of cache layer to support multiple requests and reduce duplicated calculation cost.

## Supported versions

* Go version 1.15
* Echo framework v4.1.17

## Features overview
* Get the nearest prime less than or equal to number N input

## Installation
You need go installed and `GOBIN` in your `PATH`

install `go install github.com/hieutrtr/prime-service`

start service `prime-service start --limit-prime <prime limit number>`

options
```
-limit-prime (default: 1000000) : Input that greater than limit-prime will not be supported. This configuration is also for the service to know how many prime numbers need to be generated in cache storage.
```

## Deployment
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
