# Step 1: Building the binary
FROM golang:1.10 as builder

WORKDIR /go/src/github.com/gojibjib/jibjib-api/api

RUN go get -d -v github.com/gorilla/mux

COPY api/* ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Step 2: Copy & Run
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /go/src/github.com/gojibjib/jibjib-api/api/app ./

CMD ["./app"]