FROM golang:1.23-alpine

WORKDIR /app

COPY go.* .

RUN go mod tidy
RUN go mod verify
RUN go mod download

COPY . .