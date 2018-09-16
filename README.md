# [jibjib-api](https://github.com/gojibjib/jibjib-api)

[![Go Report Card](https://goreportcard.com/badge/github.com/gojibjib/gopeana)](https://goreportcard.com/report/github.com/gojibjib/jibjib-api) [![Docker Build Status](https://img.shields.io/docker/build/obitech/jibjib-api.svg)](https://hub.docker.com/r/obitech/jibjib-api/builds/)
[![godoc reference](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/gojibjib/jibjib-api/pkg/api)

Go REST API to receive input from the [JibJib Android App](https://github.com/gojibjib/jibjib), [query](https://github.com/gojibjib/jibjib-query) the [model](https://github.com/gojibjib/jibjib-model) and send those information back to the App.

## Install
### Docker
See [deploy instructions](https://github.com/gojibjib/deploy).

### Compile yourself
If you didn't clone the repo, `go get` the package and the `main.go`:

```
go get github.com/gojibjib/jibjib-api/pkg
cd $GOPATH/src/github.com/gojibjib/jibjib-api
wget https://raw.githubusercontent.com/gojibjib/jibjib-api/master/meta/main.go 
```

Compile it:

```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
```

Setup [jibjib-data](https://github.com/gojibjib/jibjib-data).

Start the API:

```
export JIBJIB_DB_URL=test-r:test-r@localhost/birds
./app
```

## API Documentation
### Endpoints

Endpoint|Method|Comment
---|---|---
`/`|GET|Answers with a basic "Pong" response
`/ping`|GET|Answers with a basic "Pong" response
`/birds/dummy`|GET|Sends a JSON Response with randomized IDs and accuracies for testing
`/birds/all`|GET|Retrieves all bird information, without descriptions
`/birds/{id:[0-9]+}`|GET|Retrieves bird information by ID. Use query string `desc_de=false` and `desc_en=false` to omit description fields.
`/detect/binary`|POST|Send a MP4 file to the API to start querying the ML model for bird voice recognition. File needs to be send in binary format and with the Header `Content-Type: application/octet-stream` set.

### Response format

```
{
    "status": <int>,
    "message": <string>,
    "count": <int>,
    "data": {...} | null
}
```

### Examples

```
curl "http://localhost:8080/ping"
{
    "status":200,
    "message":"Pong",
    "count":0,
    "data":null
}
```

```
curl "htttp://localhost:8080/birds/1"
{
    "status":200,
    "message":"Bird found",
    "count":1,
    "data": {
        "id":1,
        "name":"Cuculus canorus",
        "genus":"Cuculus",
        "species":"canorus",
        "title_de":"Kuckuck",
        "title_en":"Common cuckoo",
        "desc_de":"...omitted...",
        "desc_en":"...omitted..."
    }
}
```

```
curl "htttp://localhost:8080/birds/1?desc_de=false&desc_en=false"
{
    "status":200,
    "message":"Bird found",
    "count":1,
    "data": {
        "id":1,
        "name":"Cuculus canorus",
        "genus":"Cuculus",
        "species":"canorus",
        "title_de":"Kuckuck",
        "title_en":"Common cuckoo",
        "desc_de":"",
        "desc_en":""
    }
}
```

```
# For now, only mp4 files are being accepted
curl -H 'Content-Type: application/octet-stream' -X POST --data-binary @larus_canus_3.mp4 http://localhost:8081/detect/binary
{
  "message": "Detection successful",
  "status": 200,
  "count": 3,
  "data": [
    {
      "accuracy": 0.6470588235294118,
      "id": 110
    },
    {
      "accuracy": 0.17647058823529413,
      "id": 7
    },
    {
      "accuracy": 0.17647058823529413,
      "id": 30
    }
  ]
}
```