# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder

RUN apk update
RUN apk add git
RUN apk add curl
RUN apk add nodejs npm

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download and cache the Go modules
RUN go mod download

# Copy the source code into the container
COPY . .

RUN npm install

# Build the Go binary for Linux
RUN LOG_LEVEL=debug GOPRIVATE=github.com/maddalax GOPROXY=direct go run github.com/maddalax/htmgo/cli/htmgo@latest build


# Stage 2: Create the smallest possible image
FROM gcr.io/distroless/base-debian11

# Set the working directory inside the container
WORKDIR /app

# Copy the Go binary from the    builder stage
COPY --from=builder /app/dist .

# Expose the necessary port (replace with your server port)
EXPOSE 3000


# Command to run the binary
CMD ["./htmgo-site"]