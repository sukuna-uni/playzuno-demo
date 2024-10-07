# Step 1: Use the official Golang image to build the application
FROM golang:1.22.4 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code to the working directory
COPY . .

# Build the Go app
# RUN go mod init app && go build -o playzuno-demo
RUN GOARCH=amd64 go build -o playzuno-demo

# Step 2: Use a smaller image for the runtime
FROM alpine:latest

# Copy the compiled binary from the builder container
COPY --from=builder /app/playzuno-demo /playzuno-demo

# Expose the port the app will run on
EXPOSE 8080

# Run the app binary
ENTRYPOINT ["/playzuno-demo"]
