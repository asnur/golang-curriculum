FROM golang:1.20.1-alpine3.17

# Install Go Present
RUN go install golang.org/x/tools/cmd/present@latest

# Workdir
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Expose port 3999 to the outside world
EXPOSE 3999

# Command to run the executable
CMD ["present", "-http", ":3999", "-use_playground"]