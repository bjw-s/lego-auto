FROM docker.io/library/golang:1.21-alpine@sha256:438bf0b13aa6d4e1de111818e3204c717f9287ede61183e82ed4e50f58434698 as build
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

FROM gcr.io/distroless/static:latest@sha256:9be3fcc6abeaf985b5ecce59451acbcbb15e7be39472320c538d0d55a0834edc
COPY --from=build /build/lego-auto /app/lego-auto
COPY --from=build /sbin/tini-static /sbin/tini
ENTRYPOINT [ "/sbin/tini", "--", "/app/lego-auto" ]
CMD [ "run" ]
VOLUME "/config"
LABEL \
  org.opencontainers.image.base.name="gcr.io/distroless/static:nonroot" \
  org.opencontainers.image.title="lego-auto"
