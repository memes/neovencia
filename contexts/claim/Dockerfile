# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:latest

# Create a working directory
WORKDIR /go/src

# Copy files
COPY claim /go/src

# Check into our working directory, build our main.go
RUN cd /go/src && go build -o main

# Document that the service listens on port 9090.
EXPOSE 9090

# Container entry point
ENTRYPOINT "./main"