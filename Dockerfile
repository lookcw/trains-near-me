# Use an official Golang image as the base image
FROM golang:1.22

# Install CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon@latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Command to run CompileDaemon to watch for changes and rebuild/restart the app
CMD ["CompileDaemon", "--build=go build -buildvcs=false -o main .", "--command=./main"]
