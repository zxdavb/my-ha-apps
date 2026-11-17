# Home Assistant App: WhatsApp MCP Extended

## What this app does

Runs [FelixIsaac/whatsapp-mcp-extended](https://github.com/FelixIsaac/whatsapp-mcp-extended) inside Home Assistant:

- **WhatsApp bridge** (Go, whatsmeow): connects as a linked device and stores messages in `/data/store`. Its REST API listens on `127.0.0.1:8080` inside the container only.
- **MCP server** (Python): exposes the WhatsApp tools over streamable HTTP on port `8081`, at `mcp.path` (default `/mcp`).

The upstream web UI is not included.

> **Warning:** this is an unofficial WhatsApp client. WhatsApp can restrict or ban accounts that use one.

## Security first

The MCP endpoint has **no authentication**. Anyone who can reach it can read your messages and send messages as you.

- Port `8081` is **not published** by default. Enable it only if you need it, and put authentication in front of it (see the README).
- `/data` holds your WhatsApp session keys and message history, and Home Assistant backups include it. Treat backups accordingly.
- In your MCP client, require confirmation for sending tools instead of "always allow".

## First start: linking WhatsApp

1. Start the app and open the **Log** tab.
2. Link it as a device, using either method:
   - **QR code:** scan the QR code printed in the log with WhatsApp > Settings > Linked devices > Link a device. It refreshes about every 20 seconds.
   - **Pairing code:** if the QR code doesn't render properly in the log, set `pairing.phone_number` to your number (international format, digits only, for example `447700900123`) and restart. An 8-character code appears in the log. On the phone, choose *Link with phone number instead* and enter it. Clear the option once linked.
3. History sync then runs in the background and can take several minutes on a busy account.

If the phone unlinks the device, restart the app to pair again.

## Options

- `pairing.phone_number`: see above. Leave empty to use the QR code only.
- `mcp.path`: endpoint path (default `/mcp`).
- `mcp.toolsets`: comma-separated tool groups, or `all` (default). Groups: `core`, `send`, `media`, `history`, `contacts_write`, `message_admin`, `groups`, `presence`, `account_admin`, `newsletter`. Note that `core` already includes `send_message`.
- `mcp.tools`: individual tool names to enable on top of the toolsets.
- `mcp.allowed_hosts`: `Host` header values to accept (for example `mcp.example.com`, or `homeassistant.local:*` for any port). Localhost is always accepted. When this is empty (default), any `Host` is accepted and DNS-rebinding protection is off. If it's set and requests fail with `421 Invalid Host header`, add the host your client or proxy actually sends.
- `advanced.debug_logging`: verbose logging.

## Notes

- If the bridge stays disconnected from WhatsApp for more than 3 minutes, it exits on purpose (upstream behaviour) and is restarted automatically.
- At startup the bridge prints its internal API key in a `curl` example line. That key only works from inside the container.
- MCP behaviour and tool implementations come from upstream.
