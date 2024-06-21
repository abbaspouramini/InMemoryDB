FROM golang:1.21.1-alpine3.17

WORKDIR /app

COPY go.mod ./

RUN go mod download



RUN go run main.go`
RUN go build -o inmemorydb

CMD ["./inmemorydb"]