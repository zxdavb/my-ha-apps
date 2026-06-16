# https://developers.home-assistant.io/docs/apps/configuration#app-dockerfile
ARG BUILD_FROM=ghcr.io/home-assistant/base:3.23
FROM ${BUILD_FROM}

ARG GARMY_REF=6560ea8e1cab9cbbca8b7fb8fb04f18c0a5a0dbe

RUN apk add --no-cache python3 py3-pip git \
    && python3 -m venv /opt/venv \
    && /opt/venv/bin/pip install --no-cache-dir --upgrade pip \
    && /opt/venv/bin/pip install --no-cache-dir \
        "garmy[all] @ git+https://github.com/zxdavb/garmy.git@${GARMY_REF}" \
        "pydantic>=2.0.0"

ENV PATH="/opt/venv/bin:${PATH}"

COPY rootfs /

LABEL \
    org.opencontainers.image.title="Home Assistant App: Garmy MCP" \
    org.opencontainers.image.description="Garmin health data MCP server for Home Assistant." \
    org.opencontainers.image.source="https://github.com/zxdavb/garmy" \
    org.opencontainers.image.licenses="MIT"
