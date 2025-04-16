FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY config/config.yml ./config/config.yml
COPY migrations ./migrations

RUN go build -o app ./cmd

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .
COPY --from=builder /app/config/config.yml ./config/config.yml
COPY --from=builder /app/migrations ./migrations

CMD ["./app"]
