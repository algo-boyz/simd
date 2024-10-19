# syntax = docker/dockerfile:1

FROM --platform=${BUILDPLATFORM} golang:1.23-alpine AS base
WORKDIR /src
ENV CGO_ENABLED=0
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

FROM base AS build
ARG TARGETOS
ARG TARGETARCH
RUN --mount=target=. \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /build/simd .

FROM scratch AS bin-unix
COPY --from=build /build/simd /

FROM bin-unix AS bin-linux
FROM bin-unix AS bin-darwin

FROM scratch AS bin-windows
COPY --from=build /build/simd /simd.exe

FROM bin-${TARGETOS} AS bin