# Changelog

## 0.4.0-ha3

- Fix "Invalid configuration" on start and on save: `advanced.wa_client_version` (added in ha2) rejected its own empty default.
- Update the bridge's WhatsApp library (whatsmeow) from 2026-07-22 to 2026-11-24 via `patches/0001-bridge-update-whatsmeow.patch`, and port the bridge to its changed APIs (blocklist, About text). This brings four months of whatsmeow fixes and WhatsApp version updates. Upstream still pins the old whatsmeow.
  - Setting the About text falls back to the legacy request when WhatsApp rejects the newer one (tulir/whatsmeow#1331).
- The startup version lookup from ha2 stays, so the bridge keeps working when WhatsApp retires versions between whatsmeow updates.

## 0.4.0-ha2

- Fix `Client outdated (405)`: WhatsApp no longer accepts the client version in upstream's pinned whatsmeow, so the bridge never offered a QR or pairing code. A packaging-only Go file now sets the current WhatsApp Web version at startup (fetched from `web.whatsapp.com`, with a baked-in fallback). No upstream files are modified.
- New option `advanced.wa_client_version` to set the version manually if the automatic lookup fails.

## 0.4.0-ha1

- Initial release, packaging upstream `whatsapp-mcp-extended` 0.4.0 at commit `be73e8c0f087`.
- Runs the Go bridge and the MCP server (streamable HTTP, port 8081) in one container under s6. The upstream web UI is not included.
- Pairing by QR code in the log, or by an 8-character pairing code via `pairing.phone_number`.
- Adds a small launcher so the MCP path and the `Host`-header check are configurable (`mcp.path`, `mcp.allowed_hosts`). Upstream's defaults only accept `Host: localhost`, which fails behind port mapping or a reverse proxy.
- Port 8081 is not published by default because the endpoint has no authentication.
