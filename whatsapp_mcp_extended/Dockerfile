# https://developers.home-assistant.io/docs/apps/configuration#app-dockerfile
ARG BUILD_FROM=ghcr.io/home-assistant/base:3.23

# Pinned upstream commit; bumped by .github/workflows/upstream-update-whatsapp-mcp-extended.yaml
ARG WHATSAPP_MCP_REF=be73e8c0f0875a3a470457c19e0fbd3449329ab5

# ---------- stage 1: fetch upstream source at the pinned commit ----------
FROM alpine:3.23 AS source
ARG WHATSAPP_MCP_REF
RUN apk add --no-cache git
WORKDIR /src
RUN git init -q . \
    && git remote add origin https://github.com/FelixIsaac/whatsapp-mcp-extended.git \
    && git fetch -q --depth 1 origin "${WHATSAPP_MCP_REF}" \
    && git checkout -q FETCH_HEAD

# ---------- stage 2: build the Go bridge (go-sqlite3 needs CGO) ----------
FROM golang:1.26-alpine AS bridge
RUN apk add --no-cache build-base
COPY --from=source /src/whatsapp-bridge /src/whatsapp-bridge
# Packaging-only Go file (init() that sets the WhatsApp Web client version);
# no upstream files are modified. See bridge-overlay/zz_ha_waversion.go.
COPY bridge-overlay/ /src/whatsapp-bridge/
WORKDIR /src/whatsapp-bridge
RUN go mod download \
    && CGO_ENABLED=1 go build -trimpath -ldflags "-s -w" -o /out/whatsapp-bridge .

# ---------- stage 3: runtime ----------
FROM ${BUILD_FROM}

RUN apk add --no-cache python3 py3-pip ffmpeg sqlite-libs ca-certificates tzdata curl jq

COPY --from=source /src/whatsapp-mcp-server /opt/whatsapp-mcp-server

# gradio/gradio_client are only used by the upstream local dev UI (gradio-main.py),
# not by main.py; upstream's own Dockerfile.mcp drops them the same way.
RUN grep -vE "^gradio" /opt/whatsapp-mcp-server/requirements.txt > /tmp/requirements.txt \
    && python3 -m venv /opt/venv \
    && /opt/venv/bin/pip install --no-cache-dir --upgrade pip \
    && /opt/venv/bin/pip install --no-cache-dir -r /tmp/requirements.txt \
    && rm -rf /opt/whatsapp-mcp-server/tests /tmp/requirements.txt

ENV PATH="/opt/venv/bin:${PATH}"

COPY --from=bridge /out/whatsapp-bridge /usr/local/bin/whatsapp-bridge

COPY rootfs /

LABEL \
    org.opencontainers.image.title="Home Assistant App: WhatsApp MCP Extended" \
    org.opencontainers.image.description="WhatsApp MCP server (bridge + MCP) for Home Assistant." \
    org.opencontainers.image.source="https://github.com/FelixIsaac/whatsapp-mcp-extended" \
    org.opencontainers.image.licenses="MIT"
