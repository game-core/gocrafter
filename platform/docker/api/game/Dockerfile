# Build
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY .. .
RUN go build -o main /app/cmd/api/game/main.go

# Prod
FROM alpine AS prod

WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 50051

CMD [ "/app/main" ]

# Local
FROM golang:1.21-alpine AS local

ENV CGO_ENABLED=0
ENV ROOT=/go/src/app
WORKDIR ${ROOT}
ENV GO111MODULE=on
COPY .. .
RUN apk upgrade --update && apk add bash && apk --no-cache add git

EXPOSE 50051

CMD [ "go", "run", "cmd/api/game/main.go" ]
