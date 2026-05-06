## 1.20.3-ha2

- Fixed OAuth 2.1 proxy disk persistence by explicitly setting `WORKSPACE_MCP_OAUTH_PROXY_DISK_DIRECTORY`.
- Added `oauth21.workspace_mcp_oauth_proxy_disk_directory` option to control where FastMCP stores OAuth proxy disk state.
- Defaulted OAuth proxy disk state path to `/data/google_workspace_credentials/oauth-proxy`.

## 1.20.3-ha1

- Reorganized configuration UI so only core credentials and external URL remain top-level.
- Kept advanced options available under grouped sections aligned with upstream environment reference categories.
- Added authentication.single_user while preserving legacy single-user config path compatibility.

## 1.20.3-ha0

- Initial repository baseline for this app.
- Tracks upstream `workspace-mcp` 1.20.3 packaging for Home Assistant.
