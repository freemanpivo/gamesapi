# STAGE: BUILD
FROM golang:1.21-alpine AS builder

RUN apk add --no-cache git build-base
WORKDIR /app
COPY go.mod go.sum data/games_seed.json ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN go build -o gamesapi ./cmd

# STAGE: RUNTIME
FROM scratch

WORKDIR /app
COPY --from=builder /app/gamesapi .
COPY --chown=0:0 ./data /app/data
EXPOSE 3000

ENTRYPOINT ["./gamesapi"]