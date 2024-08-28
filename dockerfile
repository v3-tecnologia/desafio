FROM golang:1.23-alpine as build-desafio

WORKDIR /app

COPY . .

RUN apk add make

RUN go mod download

RUN make run-tests

RUN go build -o main .

CMD ["./main"]