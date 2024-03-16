FROM docker.io/library/golang:1.22-alpine@sha256:5b45bf7d1c6f77e5d1dbb740bd1d454bf83073a319faea0e9276275ace925b60 as build
ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT=""
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=${TARGETOS} \
    GOARCH=${TARGETARCH} \
    GOARM=${TARGETVARIANT}
# hadolint ignore=DL3018
RUN apk add --no-cache ca-certificates tini-static \
    && update-ca-certificates
WORKDIR /build
COPY . .
RUN go build -ldflags="-s -w" -o lego-auto /build

FROM gcr.io/distroless/static:latest@sha256:7e5c6a2a4ae854242874d36171b31d26e0539c98fc6080f942f16b03e82851ab
COPY --from=build /build/lego-auto /app/lego-auto
COPY --from=build /sbin/tini-static /sbin/tini
ENTRYPOINT [ "/sbin/tini", "--", "/app/lego-auto" ]
CMD [ "run" ]
VOLUME "/config"
LABEL \
  org.opencontainers.image.base.name="gcr.io/distroless/static:nonroot" \
  org.opencontainers.image.title="lego-auto"
