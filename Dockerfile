# Base image
FROM golang:1.21-alpine

# Create working path
WORKDIR /app

# Copy the necessary files
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Build 
RUN go build -o resume-app .

# Port
EXPOSE 8080

#Command
CMD [ "./resume-app" ]
# https://hub.docker.com/_/golang