# Create build image and compile the service
FROM golang:1.11.1 as builder
RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/neudesic/neovencia/context/claim

COPY . .
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

# Create the runtime image 
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/neudesic/neovencia/context/claim/claim .

ENTRYPOINT ["./claim"]
# arbitrary change to test build




