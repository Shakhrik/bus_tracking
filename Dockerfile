FROM golang:1.16 as builder


#
RUN mkdir -p $GOPATH/src/billz/billz_catalog_service_v2
WORKDIR $GOPATH/src/billz/billz_catalog_service_v2

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    make build && \
    mv ./bin/billz_catalog_service_v2 /

ENTRYPOINT ["/billz_catalog_service_v2"]

