# Changelog

## 0.4.0-ha6

- Rename the app's display name from "My WhatsApp MCP Extended" to "My WhatsApp MCP". The slug (`whatsapp_mcp_extended`), hostname, options and paired session are unchanged.

## 0.4.0-ha5

- Privacy: the bridge no longer writes message text, contact names or phone numbers to the app log. Upstream logged every history-synced message in full, plus the contact behind each incoming message and each call, at INFO. These lines are now DEBUG (`patches/0003-bridge-log-privacy.patch`) and only appear with `advanced.debug_logging` on. Lines already written stay in Home Assistant's system log until it rotates.
- Pairing page: read the link status from the bridge's unauthenticated `/api/health` instead of `/api/connection`. The bridge writes an `[AUDIT]` line for every authenticated request, so with the page open the app log got one every 5 seconds, pushing the QR codes out of view.

## 0.4.0-ha4

- QR pairing: when the QR codes run out (about 2.5 minutes after start), the bridge now requests a fresh set straight away. Before, it printed nothing for about 15 minutes and then restarted, so you had to restart the app if you missed the first codes.
- QR pairing errors (for example `err-client-outdated`) are now logged instead of silently ignored.
- New pairing page (**Open Web UI**, through Home Assistant ingress) that shows the current QR code as an image and refreshes itself. The QR code printed in the log often can't be scanned because the log viewer adds space between lines.
- The bridge changes come from the new `patches/0002-bridge-qr-refresh.patch`.

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
