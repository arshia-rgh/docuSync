FROM golang:1.23.3-alpine

WORKDIR /app

COPY . .

RUN env CGO_ENABLED=0 go build -o app ./cmd/api

CMD ["./app"]
