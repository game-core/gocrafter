# Build
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main /app/gen/main.go

RUN apk update && apk add bash protobuf-dev

# Dev
FROM golang:1.21-alpine AS dev

RUN export PATH="$PATH:$(go env GOPATH)/bin"
ENV PATH=$PATH:$GOPATH/bin
ENV CGO_ENABLED=0
ENV ROOT=/go/src/app
WORKDIR ${ROOT}
ENV GO111MODULE=on
COPY . .
RUN apk upgrade --update && apk add bash && apk --no-cache add git && apk add bash protobuf-dev

RUN go install github.com/google/wire/cmd/wire@v0.5.0
RUN go install github.com/golang/mock/mockgen@v1.6.0
RUN go install golang.org/x/tools/cmd/goimports@v0.17.0
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
