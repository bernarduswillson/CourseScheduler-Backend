# Use an official Golang image as the base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application inside the container
RUN go build -o singleservice .
# RUN go run migrate/migrate.go

# Expose the port on which your application listens
# EXPOSE 8080

# Define the command to run your application when the container starts
ENTRYPOINT ["/app/singleservice"]
