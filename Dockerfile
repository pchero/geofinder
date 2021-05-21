FROM golang:1.16-alpine

LABEL maintainer="Sungtae Kim <pchero21@gmail.com>"

WORKDIR /app

# COPY go.mod go.sum ./
COPY . .

# swagger
RUN go get -v github.com/go-swagger/go-swagger/cmd/swagger
RUN go get github.com/swaggo/swag/cmd/swag
RUN go get github.com/swaggo/gin-swagger
RUN go get github.com/swaggo/files
RUN swag init -g cmd/geofinder/main.go -o docsapi

RUN go mod vendor
RUN go build ./cmd/...
