# Home Assistant App: Garmy MCP

## What this app does

Runs the upstream `garmy-mcp` server inside Home Assistant as an app.

The server provides read-only MCP access to a local Garmy SQLite database so AI assistants can query and analyze your Garmin data.

## Required setup

1. Ensure you have a Garmy SQLite database file (for example by running `garmy-sync` elsewhere).
2. Set `database.path` to that file path.
3. Start the app.

If the path does not exist, the add-on creates an empty (but valid) SQLite database so the server can start; MCP tool calls will return no data until a real database is placed at that path.

## Important options

- Non-advanced (top level):
  - `database.path`: SQLite database path passed to `--database`.
- Advanced MCP (defaults provided):
  - `advanced_mcp.transport`: MCP transport mode (default `streamable-http`).
  - `advanced_mcp.host`: Listen host (default `0.0.0.0`).
  - `advanced_mcp.port`: Listen port (default `8000`).
  - `advanced_mcp.path`: URL path (default `/mcp`).
- Advanced SQL (defaults provided):
  - `advanced_sql.max_rows`: Soft row limit (default `1000`).
  - `advanced_sql.max_rows_absolute`: Hard row limit (default `5000`).
  - `advanced_sql.enable_query_logging`: Enable SQL query logging.
  - `advanced_sql.disable_strict_validation`: Relax SQL validation (not recommended).
  - `advanced_sql.verbose`: Print expanded startup information.

## Notes

- `database.path` has no default and must be set by the user.
- `/share` is mapped so you can use a database under shared storage.
- This app is network transport only; `stdio` is not supported.
- For URL access, use `advanced_mcp.transport: streamable-http` (recommended), then connect to `http://<home-assistant-host>:8000/mcp` unless you changed host/port/path.
- Upstream MCP behavior and tool implementations are maintained by the upstream project.
