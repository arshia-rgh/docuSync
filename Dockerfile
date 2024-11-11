FROM alpine:latest

WORKDIR /app

COPY binaryApp /app

CMD ["/app/binaryApp"]
