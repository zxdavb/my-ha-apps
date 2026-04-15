# https://developers.home-assistant.io/docs/apps/configuration#app-dockerfile
ARG BUILD_FROM=ghcr.io/home-assistant/base:3.23
FROM ${BUILD_FROM}

ARG WORKSPACE_MCP_VERSION=1.19.0

RUN apk add --no-cache python3 py3-pip \
    && pip3 install --no-cache-dir "workspace-mcp[disk]==${WORKSPACE_MCP_VERSION}"

COPY rootfs /

LABEL \
    org.opencontainers.image.title="Home Assistant App: Google Workspace MCP" \
    org.opencontainers.image.description="Google Workspace MCP server for Home Assistant." \
    org.opencontainers.image.source="https://github.com/taylorwilsdon/google_workspace_mcp" \
    org.opencontainers.image.licenses="MIT"
