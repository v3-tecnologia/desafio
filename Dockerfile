FROM golang:1.22 AS BUILDER
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o desafio ./cmd/api

FROM golang:1.22-alpine AS RUNNER
COPY --from=BUILDER /app/desafio ./desafio

EXPOSE 5000
CMD ["./desafio"]


