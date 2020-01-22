# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/JackCloudman/MicroserviciosTest
ENV GOBIN /go/bin
RUN go get github.com/gorilla/mux

RUN go install github.com/JackCloudman/MicroserviciosTest

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/MicroserviciosTest
# Document that the service listens on port 8080.
EXPOSE 8090
