## 2.2.1-ha1

- Updated upstream `workspace-mcp` from 1.21.2 to 2.2.1 (FastMCP 4, MCP Python SDK v2).
- Added `server.workspace_mcp_session_idle_timeout` (default 3660 seconds) so abandoned sessions are closed.
- Upstream breaking change: the Slides `get_page_thumbnail` tool was replaced by `render_page_to_image`.

## 1.21.2-ha3

- Removed the legacy `access` option group; it only duplicated `authentication.single_user`.
- Changed the default `tool_selection.workspace_mcp_tool_tier` from `core` to `extended`.
- Removed legacy config-key fallbacks from the start script; only the current option names are read.

## 1.21.2-ha2

- Version bump only; no functional changes.

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
