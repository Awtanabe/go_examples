# 開発環境用
FROM golang:1.23-alpine AS dev

WORKDIR /app

COPY go.mod .
COPY .air.toml .
COPY . .

RUN go install github.com/air-verse/air@latest

CMD ["air", "-c", ".air.toml"]

# 本番環境用
FROM golang:1.23-alpine AS prod

WORKDIR /app

COPY go.mod .
COPY . .

RUN go build -o main

CMD ["./main"]