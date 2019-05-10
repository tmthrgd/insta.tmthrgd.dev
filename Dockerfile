# Use the offical Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.12 as builder

# Set the needed GO* environment variables.
ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org
ENV CGO_ENABLED=0

# Set the working directory to a $GOPATH subdirectory.
WORKDIR /go/src/tmthrgd.dev/go/insta.tmthrgd.dev

# Copy only the files needed to resolve go mod dependencies.
COPY go.mod go.mod
COPY go.sum go.sum

# Prefetch all the needed dependencies before copying code in. This way
# code changes that don't affect go.mod or go.sum can use cached modules.
RUN go mod download

# Copy local code to the container image.
COPY . .

# Build the command inside the container.
RUN go build -v -mod=readonly -o /server

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine

# Install needed root TLS certificate authorities.
RUN apk add ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /server /

# Run the web service on container startup.
CMD ["/server"]