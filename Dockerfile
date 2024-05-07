FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/https://github.com/zepif/Test-service
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/Test-service /go/src/https://github.com/zepif/Test-service


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/Test-service /usr/local/bin/Test-service
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["Test-service"]
