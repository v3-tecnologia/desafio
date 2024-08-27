FROM golang:1.23-alpine AS build-env

ENV host=go-db

WORKDIR /app

COPY . .

RUN apk add make

RUN go mod download

RUN make tests

RUN go build -o desafio . 

EXPOSE 8080

CMD ["./desafio"]
