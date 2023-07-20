# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the Golang code into the container
COPY . .

# Build the Golang binary
RUN go build -o main .

# Set the entry point to the Golang binary

EXPOSE 50051/tcp 
EXPOSE 50051/udp
CMD ["./main"]
