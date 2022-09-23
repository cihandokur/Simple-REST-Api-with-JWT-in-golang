FROM golang:1.19-alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /app/
COPY . /app/
RUN go build -o /app/devlab ./cmd


FROM alpine
RUN mkdir -p app
WORKDIR /app/
COPY --from=builder /app/devlab/ /app/
COPY --from=builder /app/config/config.toml /app/config/config.toml

EXPOSE 9090
