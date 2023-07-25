# Use the official Golang base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the application code into the container
COPY . .

# Build the Go application
RUN go build -o app

# Set environment variables if needed
ENV PORT=8080

# Expose the port that the API will listen on (optional)
EXPOSE 8080

# Specify the command to run the Go binary
CMD ["./app"]
