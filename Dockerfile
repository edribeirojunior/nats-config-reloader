# syntax = docker/dockerfile:1-experimental
FROM golang:1.16.3-alpine AS build
ARG TARGETOS
ARG TARGETARCH
WORKDIR /src/
COPY . /src/
RUN --mount=type=cache,target=/root/.cache/go-build \
GOOS=${TARGETOS} GOARCH=${TARGETARCH} CGO_ENABLED=0 go build -o /bin/nats-config-reloader

FROM alpine
COPY --from=build /bin/nats-config-reloader /bin/nats-config-reloader
ENTRYPOINT ["/bin/nats-config-reloader"]