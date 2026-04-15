# Home Assistant App: Google Workspace MCP

## What this app does

Runs the upstream Google Workspace MCP server inside Home Assistant as an app.

## Required setup

1. Create OAuth credentials in Google Cloud.
2. Put the Client ID in `google_oauth_client_id`.
3. Optionally add `google_oauth_client_secret` for confidential clients.

At minimum you usually need:

- `google_oauth_client_id`
- `transport` set to `streamable-http` (default)

## Important options

- `mcp_enable_oauth21`: Enable OAuth 2.1 mode.
- `workspace_mcp_stateless_mode`: Container-friendly stateless mode.
- `tool_tier`: Choose `core`, `extended`, or `complete`.
- `tools`: Optional comma-separated list like `gmail,drive,calendar`.
- `permissions`: Optional service levels like `gmail:organize drive:readonly`.

## Notes

- Credential files persist in `/data/google_workspace_credentials`.
- By default local file reads are restricted to `/share` via `allowed_file_dirs`.
- Upstream docs and behavior are maintained by the upstream project.
