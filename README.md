# [jibjib-api](https://github.com/gojibjib/jibjib-api)

[![Go Report Card](https://goreportcard.com/badge/github.com/gojibjib/gopeana)](https://goreportcard.com/report/github.com/gojibjib/jibjib-api) [![Docker Build Status](https://img.shields.io/docker/build/obitech/jibjib-api.svg)](https://hub.docker.com/r/obitech/jibjib-api/builds/)

Go REST API to receive input from the [JibJib Android App](https://github.com/gojibjib/jibjib), query the [Model](https://github.com/gojibjib/jibjib-model) and send those information back to the App.

## Endpoints

Endpoint|Method|Handler|Comment
---|---|---|---
/|GET|Ping|Answers with a basic "Pong" response
/ping|GET|Ping|Answers with a basic "Pong" response
/dummy|GET|Dummy|Sends a JSON Response with randomized IDs and accuracies for testing

## Install
### Locally
#### From Source
Get the repo:

```
$ go get github.com/gojibjib/jibjib-api/api
```

Compile it:

```
$ go build -o app github.com/gojibjib/jibjib-api
```

Run it:

```
$ ./app
```

#### Via Docker
As a stand-alone container:

```
$ docker run -d -p 8080:8080 obitech/jibjib-api
```

With `docker-compose`:

```
$ wget https://raw.githubusercontent.com/gojibjib/jibjib-api/master/docker-compose.yml -o docker-compose.yml
$ docker-compose up -d
```

#### Test your installation

```
$ curl http://localhost:8080/ping
{"status":200,"message":"Pong","data":null}
$ curl http://localhost:8080/dummy
{"status":200,"message":"Recognition successful","data":[{"id":891,"accuracy":73.57164829056475},{"id":99,"accuracy":41.64308413543738},{"id":209,"accuracy":28.93046969671797},{"id":1010,"accuracy":15.624674535701452}]}
```

### Remotely
[See deploy instructions](https://github.com/gojibjib/jibjib-api/tree/master/deploy)