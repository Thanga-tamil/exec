FROM golang:latest

WORKDIR /app

COPY go.mod go.sum config.json exec .

CMD ["./exec"]
