FROM golang:latest

WORKDIR /app

COPY go.mod go.sum config.json run .

CMD ["./run"]
