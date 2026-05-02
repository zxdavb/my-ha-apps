# https://developers.home-assistant.io/docs/apps/configuration#app-dockerfile
ARG BUILD_FROM=ghcr.io/home-assistant/base:3.23
FROM ${BUILD_FROM}

COPY requirements.txt /tmp/requirements.txt

RUN apk add --no-cache python3 py3-pip \
    && python3 -m venv /opt/venv \
    && /opt/venv/bin/pip install --no-cache-dir --upgrade pip \
    && /opt/venv/bin/pip install --no-cache-dir -r /tmp/requirements.txt

ENV PATH="/opt/venv/bin:${PATH}"

COPY rootfs /

LABEL \
    org.opencontainers.image.title="Home Assistant App: Google Workspace MCP" \
    org.opencontainers.image.description="Google Workspace MCP server for Home Assistant." \
    org.opencontainers.image.source="https://github.com/taylorwilsdon/google_workspace_mcp" \
    org.opencontainers.image.licenses="MIT"
