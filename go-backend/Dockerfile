FROM golang:1.23-alpine

WORKDIR /app

ENV GIN_MODE=release
ENV PORT=8080

COPY go.mod go.sum ./

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN go mod download

COPY *.go ./

RUN swag init

RUN go build -o server

EXPOSE $PORT

CMD ["./server"]
