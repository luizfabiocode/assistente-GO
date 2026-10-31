# Estágio 1: Builder
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY cmd/ ./cmd/
COPY internal/ ./internal/

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o botgo ./cmd/main.go

# Estágio 2: Final
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

RUN mkdir -p data

COPY --from=builder /app/botgo .

ENTRYPOINT ["./botgo"]

