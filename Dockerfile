# Builder container
FROM golang:1.22.2-alpine3.19 AS builder

# Labels for the image
LABEL maintainer="Cloudenberg"
LABEL authors="Erlend, Arthur, Martin, Oskar"
LABEL version="1.0"
LABEL stage="builder"

# Set the working directory
WORKDIR /db_project

# Copy files
COPY . .

# Compile binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o server .



# Target container
FROM scratch

# Root as working directory to copy the binary
WORKDIR /

# Copy the binary from the builder container
COPY --from=builder /db_project/server .

# Expose internal port
EXPOSE 8080

# Run the binary
CMD ["./server"]
