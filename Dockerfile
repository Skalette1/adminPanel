FROM golang:1.21-alpine AS builder

RUN apk add --no-cache git make bash

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN swag init

RUN go build -o adminPanel main.go

FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/adminPanel .
COPY --from=builder /app/docs ./docs

EXPOSE 8080

ENTRYPOINT ["./adminPanel"]