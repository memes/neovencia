# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

WORKDIR /go/src

COPY web /go/src

RUN echo $(ls -1 /go/src/web)

# Check into our working directory, build our main.go
RUN cd /go/src && go build -o main

# Document that the service listens on port 8080.
EXPOSE 8080

ENTRYPOINT "./main"