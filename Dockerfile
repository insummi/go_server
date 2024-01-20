FROM golang:1.21-alpine

WORKDIR /opt/code/
ADD ./ /opt/code/

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN go mod download

RUN go build -o bin/goserver cmd/main.go
ENTRYPOINT [ "bin/goserver" ]