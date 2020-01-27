# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/client

# Build inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
WORKDIR /go/src/client
RUN go get .

# Run the client command by default when the container starts.
ENTRYPOINT /go/bin/client

# Document that the service listens on port 8080.
EXPOSE 8080