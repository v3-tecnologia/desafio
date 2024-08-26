FROM golang:1.23-alpine AS build-env

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o desafio . 

EXPOSE 8080

CMD ["./desafio"]
