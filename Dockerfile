# Build stage
FROM golang:1.22.2 AS BUILDER

# Maintainers and authors
LABEL maintainer="Cloudenberg"
LABEL authors="Erlend, Arthur, Oskar, Martin"

# Workdir name for image
WORKDIR /db_project

# Copy the entire project directory
COPY . .

# Compile binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o executable .

# Expose internal port
EXPOSE 9000

# Run executable binary
ENTRYPOINT ["./executable"]