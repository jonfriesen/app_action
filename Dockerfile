# Start from the latest golang base image
FROM golang:alpine
# Set the Current Working Directory inside the container
WORKDIR /app
# Copy go mod and sum files
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
# Copy the source from the current directory to the Working Directory inside the container
COPY . .
# Build the Go app
RUN go build -o main main.go
# Command to run the executable
RUN chmod +x main

CMD ["./main"]