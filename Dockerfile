# Base image
FROM golang:1.23-alpine

# Create working directory
WORKDIR /app

# Copy go mod files first (for better caching)
COPY go.mod ./
COPY go.su[m] ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o resume-app .

# Expose port
EXPOSE 8080

# Run the application
CMD ["./resume-app"]
# https://hub.docker.com/_/golang