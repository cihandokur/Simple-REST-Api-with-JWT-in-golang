FROM golang:1.19.5-alpine3.17 AS builder
WORKDIR /app/
COPY . .
RUN go build -o /app/simple_rest_api ./cmd


FROM alpine:3.17
RUN mkdir -p app
WORKDIR /app/
COPY --from=builder /app/simple_rest_api/ /app/
COPY --from=builder /app/config/config.toml /app/config/config.toml

EXPOSE 9090
