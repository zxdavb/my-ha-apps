## 0.1.7

- Default to all Workspace services (no explicit tool allowlist).
- Set `workspace_external_url` default to `https://hass.example.com`.
- Restrict OAuth dynamic client redirect URIs to Claude hosted callbacks by default.

## 0.1.6

- Restructure app options/schema to align with upstream environment variable categories.
- Set defaults for reverse-proxy OAuth 2.1 usage with persistent OAuth proxy storage (`disk`).
- Default tool selection to `gmail,calendar,tasks,docs,sheets,slides` with `core` tier.
- Update startup mapping to support new grouped keys while keeping backward compatibility with older keys.

## 0.1.5

- Update upstream `workspace-mcp[disk]` to `1.20.3`.
- Move Python dependency pin to `requirements.txt`.
- Enable Dependabot `pip` updates for this app so dependency bumps are proposed automatically.

## 0.1.4

- Move `workspace_external_url` to a top-level option for easier configuration.

## 0.1.3

- Initial release of Google Workspace MCP app.
