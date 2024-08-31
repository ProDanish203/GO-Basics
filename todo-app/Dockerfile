FROM golang:1.23.0-alpine3.20 AS builder

WORKDIR /app

copy go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 8000

CMD ["go", "run", "main.go"]