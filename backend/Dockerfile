FROM golang:1.23.1-alpine AS builder

# 作業ディレクトリの設定
WORKDIR /var/www/backend

COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go install github.com/air-verse/air@latest

CMD mkdir -p /var/www/backend/tmp/upload && air -c .air.toml