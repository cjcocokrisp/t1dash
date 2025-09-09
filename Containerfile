# Build Step
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apk update && \
    apk add \
    go-task && \
    rm -rf /var/lib/apt/lists/*

RUN go-task build

# Runtime Step
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/dist/t1dash .

CMD ["./t1dash", "server"]
