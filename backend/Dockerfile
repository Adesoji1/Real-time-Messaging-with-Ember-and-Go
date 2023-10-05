# # Use the official Go image as a base image
# FROM golang:1.17

# # Set the working directory inside the container
# WORKDIR /app

# # Copy the go.mod and go.sum files to the container
# COPY go.mod go.sum ./

# # Download all dependencies
# RUN go mod download

# # Copy the entire project to the container
# COPY . .

# # Build the Go app
# RUN go build -o main .

# # Expose port 8080 to the outside world
# EXPOSE 8080

# # Command to run the executable
# CMD ["./main"]


# Use the official Go image as a base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy go.mod and go.sum to pre-download dependencies
COPY go.mod go.sum ./

# Download and verify all dependencies. 
# This step will only re-run if go.mod or go.sum changes, leveraging Docker cache.
RUN go mod download && go mod verify

# Copy the entire backend project to the container
COPY . .

# Build the Go app
RUN go build -v -o /usr/local/bin/app .


# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the app
CMD ["app"]
