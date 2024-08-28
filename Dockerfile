FROM golang:1.22-alpine AS build-env

SHELL [ "/bin/sh", "-c"]

# Create an /app directory within a given image that will hold the application source files
RUN mkdir /app

# Copy everything in the root directory into /app directory
ADD . /app

# Specify that now any further commands will be executed inside /app directory
WORKDIR /app

RUN go mod download

RUN go build -o main ./cmd

# Expose port 7001 to the outside world
EXPOSE 7001

CMD ["./main"]
