# Use the official Go image from DockerHub
FROM golang:1.17

# Set the Current Working Directory inside the container
WORKDIR /app

# Disable Go Modules
ENV GO111MODULE=off

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o main .

# Run the binary
CMD ["./main"]
