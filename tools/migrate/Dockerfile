FROM golang:1.22-alpine

WORKDIR /app

RUN apk add --no-cache git \
    && go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY migrations migrations
