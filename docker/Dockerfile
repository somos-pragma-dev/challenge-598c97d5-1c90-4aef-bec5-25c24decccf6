# Dockerfile for the microservice
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum./
RUN go mod download

COPY../
RUN go build -o payment-microservice

FROM alpine:latest

COPY --from=builder /app/payment-microservice /app/

WORKDIR /app

EXPOSE 8080

CMD ["./payment-microservice"]