# STAGE: BUILD
FROM golang:1.21-alpine AS builder

RUN apk add --no-cache git build-base
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN go build -o /out/games-api ./main.go

# STAGE: RUNTIME
FROM alpine:3.18

RUN apk add --no-cache ca-certificates curl
RUN adduser -D -g '' appuser
USER appuser
WORKDIR /app
COPY --from=builder /out/games-api /app/games-api
RUN mkdir -p /app/data
EXPOSE 3000

ENTRYPOINT ["/app/games-api"]