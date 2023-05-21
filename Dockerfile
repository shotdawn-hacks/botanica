FROM golang:1.20.3-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN apk update && \
    apk add build-base

RUN go mod download

COPY . .

RUN go build -o botanica main.go

FROM alpine

WORKDIR /app

COPY --from=builder /app/botanica /app/botanica

CMD ["./botanica", "plants"]

EXPOSE 9000


