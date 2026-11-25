# Home Assistant App: WhatsApp MCP Extended

WhatsApp MCP server app for Home Assistant: read, search and send WhatsApp messages from an MCP client.

## Upstream

| | |
| --- | --- |
| **Project** | [FelixIsaac/whatsapp-mcp-extended](https://github.com/FelixIsaac/whatsapp-mcp-extended) (extends [lharries/whatsapp-mcp](https://github.com/lharries/whatsapp-mcp)) |
| **Pinned** | A specific upstream commit (`ARG WHATSAPP_MCP_REF` in the [Dockerfile](./Dockerfile)), bumped by the upstream-update workflow |

The image builds the upstream Go bridge from source, installs the upstream Python MCP server into a virtualenv, and runs both under s6. The upstream Next.js web UI is not included. Pairing is done with the QR code in the log, or with a pairing code (see [DOCS.md](./DOCS.md)).

> **Warning:** this is an unofficial WhatsApp client. WhatsApp can restrict or ban accounts that use one.

## Packaging differences from upstream

- **Single container:** the bridge and MCP server run as two s6 services. The bridge API listens only on `127.0.0.1:8080`, using an internal API key that is generated once and stored in `/data/bridge_api_key`.
- **Launcher:** upstream's `main.py` builds its MCP server object with hard-coded defaults, so `FASTMCP_*` environment variables are ignored. The MCP SDK then only accepts `Host: localhost` and answers everything else with `421 Invalid Host header`. `whatsapp-mcp-launch` imports `main.py` unchanged and sets the path and `Host` check from the app options before starting it.
- **Port not published:** port `8081` is not published by default.
- **WhatsApp client version:** `bridge-overlay/zz_ha_waversion.go` is added to the bridge's `main` package at build time. Its `init()` sets the WhatsApp Web version that whatsmeow announces, because the version in upstream's pinned whatsmeow is rejected with `Client outdated (405)`. Newer whatsmeow commits change APIs the bridge uses, so bumping the library is not an option without upstream changes.

## Authentication and external proxy (NPM)

The MCP endpoint has **no authentication**. Anyone who can reach it can read your messages and send messages as you. As with the other apps here, put authentication at the proxy.

1. Choose how Nginx Proxy Manager reaches the app:
   - **NPM running as a Home Assistant app:** use the internal network and leave port 8081 unpublished. Apps can reach each other by hostname ([app communication](https://developers.home-assistant.io/docs/apps/communication)). The forward host is this app's hostname, normally `<repository-hash>-whatsapp-mcp-extended` (unverified for this repository, so confirm it). The port is `8081`.
   - **NPM running elsewhere:** publish `8081/tcp` in the app's Network settings and forward to `<home-assistant-ip>:<published-port>`. Anything else on your LAN can then reach it too.
2. Enable TLS (Let's Encrypt) and force HTTPS.
3. Add authentication at the proxy. The minimum is an NPM Access List; better is OAuth/OIDC forward auth.
4. Set `mcp.allowed_hosts` to your public hostname (for example `mcp.example.com`) to turn DNS-rebinding protection back on.
5. Point MCP clients at `https://mcp.example.com/mcp`, or your `mcp.path`.

For long-lived MCP HTTP sessions, add these to the NPM advanced config:

```nginx
proxy_buffering off;
proxy_request_buffering off;
proxy_read_timeout 3600;
proxy_send_timeout 3600;
```

## Wrapper references

These existing Home Assistant wrappers of the original `lharries/whatsapp-mcp` were reviewed for reference. No code was taken from them:

- [patricklis91/ha-addons](https://github.com/patricklis91/ha-addons) (`whatsapp_mcp`)
- [klangenk/ha-addons](https://github.com/klangenk/ha-addons) (`whatsapp_bridge`)

![Supports amd64 Architecture][amd64-shield]

## License

- This app packaging in this repository is licensed under MIT.
- See [LICENSE](../LICENSE).
- Upstream project licensing remains under upstream terms (MIT).

[amd64-shield]: https://img.shields.io/badge/amd64-yes-green.svg
