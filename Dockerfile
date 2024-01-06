FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . /app/
RUN CGO_ENABLED=0 GOOS=linux go build -o random_exporter .


# Second stage: create the final image
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/random_exporter /usr/local/bin/random_exporter

# Expose the port the app runs on
EXPOSE 8080

# Run the binary
CMD ["/usr/local/bin/random_exporter"]