# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Go app
RUN go build -o main .

# Expose port 3030 for the Go app to listen on
EXPOSE 3030

# Set environment variables
ENV MYSQL_HOST=localhost MYSQL_USER=username MYSQL_PASSWORD=secretpw MYSQL_DBNAME=dbname

# Run the Go app
CMD ["./main"]
