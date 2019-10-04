# URL Checker

## Introduction

We have an HTTP proxy that is scanning traffic looking for malware URLs. Before allowing HTTP connections to be made, this proxy asks a service that maintains several databases of malware URLs if the resource being requested is known to contain malware.

This small web service responds to GET requests where the caller passes in a URL and the service responds with some information about that URL.

## REST API Service Via HTTP GET Request

URL checker is a REST API service via HTTP GET request that gives status of given URL.

URL has 3 types of status:
* `safe` : URL is safe to access on basis of the database. This URL is known as good and safe website, and would be considered as non-malicious URL
* `unsafe` : URL is not safe to access on basis of the database. This URL is known as bad and unsafe website, and would be considered as malicious URL
* `Unknown` : Given URL is not in the database, and would be considered as unknown URL

## REST API Service Architecture

![REST API Architecture](https://github.com/doyunbk/URL-Checker/blob/master/REST_API_architecture.png)

## Installation 

#### Using sources

Install Golang, clone the repository and build:

```sh
$ git clone https://github.com/doyunbk/URL-Checker.git
$ cd url-checker
$ go build
```

#### Dependency 
Install Redis on local machine
```sh
On Mac OS
$ brew install redis
```
```sh
On Linux
$ sudo apt-get install redis-server
```
```sh
On Windows
Download from "https://github.com/dmajkic/redis/downloads"
```
Install Golang-Redis package on local machine
```sh
$ go get github.com/gomodule/goredis
```

#### Run Application

Run this application on command line

```sh
$ go run main.go
```

## Docker

Run this application using Docker
```sh
$ docker-compose build
$ docker-compose up
```
## Makefile

Run this application using Makefile
```sh
$ make build
$ make run
```
Test 5 unit test cases using Makefile, make sure redis is intalled
```sh
$ cd test
$ make redis
$ make test
```



## Database

#### Redis Database Setup & Database Seeding
Setup Redis database and seed database with test data
```sh
$ redis-server
$ redis-cli
> HMSET www.example.com url "www.example.com" status "Unsafe"
> HMSET www.example1.com url "www.example1.com" status "Safe"
```

#### Key-Value Database

Redis stores data in a simple key-value pair.
```sh
+==================+==============================+
|        Key       |            Value             |
+==================+==================+===========+
|         -        |        URL       |   Status  |
+------------------+------------------+-----------+
|  www.example.com |  www.example.com |  Unsafe   |
| www.example1.com | www.example1.com |   Safe    |
+------------------+------------------+-----------+
```

## Usage

#### Check URL Status

HTTP GET request checks the given url status from `REDIS` database
`GET /urlinfo/1/{hostname_and_port}/{original_path_and_query_string}`

#### Example

##### Check URL is validated

If URL is invalidated, HTTP response is made with the following log

```sh
$ curl -X GET localhost:8000/;zc3b:-$`www.validateurl.com/zv/?bceq**&dvcse/
$ Cannot validate url
```

##### Once URL is validated

Check URL is a safe website or not, and HTTP response is made with JSON object of a given url
```sh
$ curl -X GET localhost:8000/www.example.com/&qed?cxvvczd#&/z&32d
```
```sh
{
    "URL":"www.example.com",
    "Status":"Unsafe"
}
```

```sh
$ curl -X GET localhost:8000/www.example1.com/?cvde#?bcx34g1dwe/zcv~@#asz/
```
```sh
{
    "URL":"www.example1.com",
    "Status":"Safe"
}
```

Check URL is not in the database, and HTTP response is made with the following log
```sh
$ curl -X GET localhost:8000/www.example2.com/&fvcx$233/vcds!?54dza
$ Unknown url: cannot found in DB
```
Check URL is not given, and HTTP response is made with the following log
```sh
$ curl -X GET localhost:8000/
$ No url is given, please provide URL
```


## Test

There are 5 test cases for this application.
##### 1. Test for invalidated url
Assign URL to test whether it is an invalidated url
##### 2. Test for unsafe URL from DB 
Assign URL to test whether it is an unsafe website on basis of the database
##### 3. Test for safe URL from DB
Assign URL to test whether it is a safe website on basis of the database
##### 4. Test for URL not in DB
Assign URL to test whether it is not in the database, considered to be unknown url
##### 5. Test for given URL empty
Do not assign any URL to test whether this app gives an error message to provide URL

#### Testing all test cases

```sh
$ go test -v

=== RUN   TestInvalidateUrl
--- PASS: TestInvalidateUrl (0.00s)
=== RUN   TestUrlUnsafeFromDb
--- PASS: TestUrlUnsafeFromDb (0.00s)
=== RUN   TestUrlSafeFromDb
--- PASS: TestUrlSafeFromDb (0.00s)
=== RUN   TestUrlNotInDb
--- PASS: TestUrlNotInDb (0.00s)
=== RUN   TestGivenUrlEmpty
--- PASS: TestGivenUrlEmpty (0.00s)
PASS
ok      github.com/url-checker/test     0.020s
```

## Thought Exercise

Please review [Thought_Excercise.pdf](https://github.com/doyunbk/URL-Checker/blob/master/Thought_Exercise.pdf)