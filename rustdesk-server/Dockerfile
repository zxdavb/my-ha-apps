# https://developers.home-assistant.io/docs/apps/configuration#app-dockerfile
ARG BUILD_FROM=ghcr.io/home-assistant/base:3.23
FROM ${BUILD_FROM}

ARG TARGETARCH
ARG RUSTDESK_SERVER_VERSION=1.1.15

RUN apk add --no-cache --virtual .build-deps curl unzip \
    && mkdir -p /opt \
    && case "${TARGETARCH}" in \
        amd64) rustdesk_arch="amd64" ;; \
        arm64) rustdesk_arch="arm64v8" ;; \
        *) echo "Unsupported TARGETARCH: ${TARGETARCH}" && exit 1 ;; \
    esac \
    && curl -fSL -o /tmp/rustdesk-server.zip \
        "https://github.com/rustdesk/rustdesk-server/releases/download/${RUSTDESK_SERVER_VERSION}/rustdesk-server-linux-${rustdesk_arch}.zip" \
    && unzip /tmp/rustdesk-server.zip -d /tmp/rustdesk-server \
    && mv "/tmp/rustdesk-server/${rustdesk_arch}" /opt/rustdesk-server \
    && chmod +x /opt/rustdesk-server/hbbs /opt/rustdesk-server/hbbr /opt/rustdesk-server/rustdesk-utils \
    && rm -rf /tmp/rustdesk-server.zip /tmp/rustdesk-server \
    && apk del --no-cache --purge .build-deps

COPY rootfs /

LABEL \
    org.opencontainers.image.title="Home Assistant App: RustDesk Server" \
    org.opencontainers.image.description="Self-hosted RustDesk server app for Home Assistant." \
    org.opencontainers.image.source="https://github.com/rustdesk/rustdesk-server" \
    org.opencontainers.image.licenses="MIT"
