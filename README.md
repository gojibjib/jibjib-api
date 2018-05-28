# [jibjib-api](https://github.com/gojibjib/jibjib-api)

[![Go Report Card](https://goreportcard.com/badge/github.com/gojibjib/gopeana)](https://goreportcard.com/report/github.com/gojibjib/jibjib-api) [![Docker Build Status](https://img.shields.io/docker/build/obitech/jibjib-api.svg)](https://hub.docker.com/r/obitech/jibjib-api/builds/)

Go REST API to receive input from the [JibJib Android App](https://github.com/gojibjib/jibjib), query the [Model](https://github.com/gojibjib/jibjib-model) and send those information back to the App.

## Install
### Remotely
[See deploy instructions](https://github.com/gojibjib/jibjib-api/tree/master/deploy)

### Locally
#### Clone the repo
```
git clone https://github.com/gojibjib/jibjib-api
```

Or if you have Go installed, you can clone it into your Gopath:

```
git clone https://github.com/gojibjib/jibjib-api $GOPATH/src/github.com/gojibjib/jibjib-api
```

#### MongoDB
Create directories, if necessary:

```
mkdir -p db/import db/data
```

Grab the JSON data from Github:

```
wget https://raw.githubusercontent.com/gojibjib/voice-grabber/master/meta/birds.json -o /db/data/birds.json
```

We need to launch the container first to initialize the database:

```
docker run --rm -d --name mongo -d \
	-v $(pwd)/data:/data/db \
	-v $(pwd)/conf:/etc/mongo \
	-v $(pwd)/import:/import \
	-v $(pwd)/initdb:/initdb \
	mongo:3.6.5 --config=/etc/mongo/mongod.conf
```

And setup users:

```
docker exec mongo bash /initdb/setup.sh
```

Next stop the container:

```
docker container rm -f mongo
```

#### Compilation
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

Start the database:

```
docker-compose up mongo
```

Start the API:

```
export JIBJIB_DB_URL=test-r:test-r@localhost/birds
./app
```

### Via Docker

```
docker-compose up -d
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