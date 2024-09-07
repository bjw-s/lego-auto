FROM docker.io/library/golang:1.23-alpine@sha256:ac67716dd016429be8d4c2c53a248d7bcdf06d34127d3dc451bda6aa5a87bc06 as build
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

FROM gcr.io/distroless/static:latest@sha256:95eb83a44a62c1c27e5f0b38d26085c486d71ece83dd64540b7209536bb13f6d
COPY --from=build /build/lego-auto /app/lego-auto
COPY --from=build /sbin/tini-static /sbin/tini
ENTRYPOINT [ "/sbin/tini", "--", "/app/lego-auto" ]
CMD [ "run" ]
VOLUME "/config"
LABEL \
  org.opencontainers.image.base.name="gcr.io/distroless/static:nonroot" \
  org.opencontainers.image.title="lego-auto"
