# Build stage
FROM golang:1.20-alpine AS builder
WORKDIR /app

COPY ./ ./

RUN go mod download
RUN go build -o main .

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .

# Expose the port on which the server will run
EXPOSE 6969

# Command to run the Go application
CMD ["/app/main"]
