# Build
FROM golang:1.19-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main /app/gen/main.go

# Run Development
FROM golang:1.19-alpine AS dev

ENV CGO_ENABLED=0
ENV ROOT=/go/src/app
WORKDIR ${ROOT}
ENV GO111MODULE=on
COPY . .
RUN apk upgrade --update && apk add bash && apk --no-cache add git

RUN go install github.com/swaggo/swag/cmd/swag@v1.8.0
RUN go install github.com/google/wire/cmd/wire@v0.5.0
