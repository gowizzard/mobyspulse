# Dockerfile for the Moby's Pulse application.
# This file details the process for constructing a lightweight and efficient Docker image using a multi-stage build process.
# The chosen base is Alpine Linux for its minimalistic size, while still providing necessary functionalities.

########################################################################################################################

# This phase uses the Alpine-based Go image to compile the source code of the application.
# By parameterizing the Go version, it becomes straightforward to maintain and modify in the future.
ARG GO_VERSION=1.22
FROM golang:${GO_VERSION}-alpine AS build
RUN apk add --no-cache git make
WORKDIR /tmp/src
COPY . .
RUN make build

########################################################################################################################

# The final preparation phase for the production-ready image. The compiled binary is copied from the previous stage,
# and the essential system packages are added and the correct timezone is set.
FROM alpine:latest AS production
RUN apk add --no-cache curl
ENV TZ=Europe/Berlin
WORKDIR /app
COPY --from=build /tmp/src/mobyspulse .
EXPOSE 3000
CMD ["/app/mobyspulse"]
HEALTHCHECK --interval=30s --timeout=10s --retries=3 \
    CMD curl --silent --unix-socket /var/run/docker.sock http://localhost/_ping || exit 1