# Home Assistant App: Garmy MCP

## What this app does

Runs the upstream `garmy-mcp` server inside Home Assistant as an app.

The server provides read-only MCP access to a local Garmy SQLite database so AI assistants can query and analyze your Garmin data.

## Required setup

1. Set `garmin.email` and `garmin.password` to your Garmin Connect credentials.
2. Set `database.path` (default `/data/health.db` is fine for most users).
3. Start the app — the sync service runs immediately and then on the configured interval.

Alternatively, if you have existing token files from a prior `garmy-sync` run, you can copy them instead of entering credentials (see the README).

If the database does not exist yet, the add-on creates an empty one so the MCP server can start while the first sync runs.

## Important options

- `database.path`: Path to the SQLite database file (default `/data/health.db`).
- `garmin.email`: Garmin Connect account email (leave empty if using token files).
- `garmin.password`: Garmin Connect password (leave empty if using token files).
- `sync.interval_hours`: Hours between sync runs (default `6`, max `168`).
- `sync.history_days`: Days of history to sync on each run (default `7`, max `90`).
- `mcp.host`: MCP server bind address (default `0.0.0.0`).
- `mcp.path`: URL path (default `/mcp`).

## Notes

- `/share` is mapped so you can use a database under shared storage.
- This app always serves MCP over `streamable-http` on port 8000 (see `ports:`); other transports are not available.
- Connect to `http://<home-assistant-host>:8000/mcp` unless you changed host/path.
- Upstream MCP behavior and tool implementations are maintained by the upstream project.
