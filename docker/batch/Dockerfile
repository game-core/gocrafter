# Build
FROM golang:1.19-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main /app/batch/main.go

# Run Production
FROM alpine AS prod

WORKDIR /app
COPY --from=builder /app/main .
COPY .env .
RUN mkdir /app/log
COPY /log/batch_debug.log /app/log
COPY wait-for.sh .
RUN chmod +x wait-for.sh

# Run Development
FROM golang:1.19-alpine AS dev

ENV CGO_ENABLED=0
ENV ROOT=/go/src/app
WORKDIR ${ROOT}
ENV GO111MODULE=on
COPY . .
RUN apk upgrade --update && apk add bash && apk --no-cache add git
