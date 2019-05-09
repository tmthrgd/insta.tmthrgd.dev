# Use the offical Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.12 as builder

# Copy local code to the container image.
WORKDIR /go/src/tmthrgd.dev/go/insta.tmthrgd.dev
COPY . .

# Build the command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN CGO_ENABLED=0 GOOS=linux GO111MODULE=on go build -v -o server

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine

RUN apk add ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /go/src/tmthrgd.dev/go/insta.tmthrgd.dev/server /server

# Run the web service on container startup.
CMD ["/server"]