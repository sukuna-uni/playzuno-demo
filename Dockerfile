# Step 1: Use the official Golang image to build the application
FROM golang:1.22.4 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code to the working directory
COPY . .

# Build the Go app
# RUN go mod init app && go build -o playzuno-demo
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o playzuno-demo

FROM debian:bullseye-slim

# Copy the compiled binary from the builder container
COPY --from=builder /app/playzuno-demo /playzuno-demo

# Verify that the binary exists and is executable
RUN ls -l /playzuno-demo
RUN file /playzuno-demo

# Expose the port the app will run on
EXPOSE 8080

# Run the app binary
ENTRYPOINT ["/playzuno-demo"]
