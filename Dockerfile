# Use the official Golang image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container
COPY . .
COPY .env .

# Build the Go application
RUN go build -o tech-assignment-backend-engineer

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./tech-assignment-backend-engineer"]