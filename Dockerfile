# workspace (GOPATH) configured at /go
FROM golang:1.15 as builder


#
RUN mkdir -p $GOPATH/src/github.com/Shakhrik/bus_tracking/api
WORKDIR $GOPATH/src/github.com/Shakhrik/bus_tracking/api

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    go mod vendor && make build && \
    mv ./bin/safia_catalogue_service /

FROM alpine
COPY --from=builder safia_catalogue_service .

ENTRYPOINT ["/safia_catalogue_service"]