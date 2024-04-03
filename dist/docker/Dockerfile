FROM docker.io/library/golang:1.22-alpine@sha256:cdc86d9f363e8786845bea2040312b4efa321b828acdeb26f393faa864d887b0 as build
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

FROM gcr.io/distroless/static:latest@sha256:6d31326376a7834b106f281b04f67b5d015c31732f594930f2ea81365f99d60c
COPY --from=build /build/lego-auto /app/lego-auto
COPY --from=build /sbin/tini-static /sbin/tini
ENTRYPOINT [ "/sbin/tini", "--", "/app/lego-auto" ]
CMD [ "run" ]
VOLUME "/config"
LABEL \
  org.opencontainers.image.base.name="gcr.io/distroless/static:nonroot" \
  org.opencontainers.image.title="lego-auto"
