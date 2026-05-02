# Home Assistant App: Google Workspace MCP

## What this app does

Runs the upstream [Google Workspace MCP server](https://github.com/taylorwilsdon/google_workspace_mcp) inside Home Assistant as an app.

## Required setup

1. Create OAuth credentials in Google Cloud.
2. Put the Client ID in `google_oauth_client_id`.
3. Optionally add `google_oauth_client_secret` for confidential clients.

At minimum you usually need:

- `google_oauth_client_id`
- `transport` set to `streamable-http` (default)

## Important options

- `server.workspace_mcp_transport`: `streamable-http` is recommended.
- `tool_selection.workspace_mcp_tools`: Comma-separated services to expose.
- `tool_selection.workspace_mcp_tool_tier`: `core`, `extended`, or `complete`.
- `oauth21.mcp_enable_oauth21`: Enable OAuth 2.1 mode.
- `oauth21.workspace_mcp_oauth_proxy_storage_backend`: `disk` for persistent single-server state.
- `workspace_external_url`: Public URL for reverse proxy deployments.

## Notes

- Credential files persist in `/data/google_workspace_credentials` by default.
- By default local file reads are restricted to `/share` via `allowed_file_dirs`.
- Upstream docs and behavior are maintained by the upstream project.
