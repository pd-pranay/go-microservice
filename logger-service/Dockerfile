# Use an official Go runtime as a parent image
FROM golang:1.19-alpine

# Set the working directory in the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Download and install any required third-party dependencies into the container

# Build the GoFiber application
RUN go build -o main ./cmd/api

CMD ["./main"]
