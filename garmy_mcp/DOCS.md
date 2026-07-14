# Home Assistant App: Garmy MCP

## What this app does

Runs the upstream `garmy-mcp` server inside Home Assistant as an app.

The server provides read-only MCP access to a local Garmy SQLite database so AI assistants can query and analyze your Garmin data.

## Required setup

This add-on runs headlessly and cannot complete Garmin's interactive MFA prompt, so it does not accept a Garmin email/password. Authenticate once elsewhere (interactively, where you can enter an MFA code), then get the resulting tokens into the add-on:

1. Authenticate with `garmy` on any machine you control, producing `~/.garmy/oauth1_token.json`.
2. Paste its contents into `garmin.oauth1_token_json` (see the README for the exact steps), or copy the token file(s) directly into `/data/.garmy/`.
3. Set `database.path` (default `/data/health.db` is fine for most users).
4. Start the app — the sync service runs immediately and then on the configured interval.

If the database does not exist yet, the add-on creates an empty one so the MCP server can start while the first sync runs.

## Important options

- `database.path`: Path to the SQLite database file (default `/data/health.db`).
- `garmin.oauth1_token_json`: Paste the contents of `oauth1_token.json` obtained via an interactive login elsewhere. `garmy-sync` mints its own OAuth2 access token from it, so no second field is needed. Written only when this value changes (i.e. a fresh login), so a token `garmy-sync` refreshes on disk isn't clobbered by config on every restart. See the README.
- `sync.interval_hours`: Hours between sync runs (default `6`, max `168`).
- `sync.history_days`: Days of history to sync on each run (default `7`, max `999`).
- `mcp.host`: MCP server bind address (default `0.0.0.0`).
- `mcp.path`: URL path (default `/mcp`).
- `advanced.debug_logging`: Adds Garmin token-file diagnostics (existence, size, JSON validity) to the log, for troubleshooting sync/auth issues. Leave off for normal use.

## Notes

- `/share` is mapped so you can use a database under shared storage.
- This app always serves MCP over `streamable-http` on port 8000 (see `ports:`); other transports are not available.
- Connect to `http://<home-assistant-host>:8000/mcp` unless you changed host/path.
- Upstream MCP behavior and tool implementations are maintained by the upstream project.
