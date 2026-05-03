# Home Assistant App: Google Workspace MCP

Google Workspace MCP server app for Home Assistant.

## Upstream

| | |
|---|---|
| **Project** | [taylorwilsdon/google_workspace_mcp](https://github.com/taylorwilsdon/google_workspace_mcp) |
| **PyPI package** | [workspace-mcp](https://pypi.org/project/workspace-mcp/) |
| **Releases** | [GitHub Releases](https://github.com/taylorwilsdon/google_workspace_mcp/releases) |

This app installs the `workspace-mcp[disk]` PyPI package into a virtualenv on top of the HA base image. The PyPI package is the canonical distribution artifact — no intermediate Docker image is used.

## Wrapper references

No intermediary wrapper repository is used for this app. Packaging is maintained directly in this repository against upstream project and PyPI releases.

![Supports amd64 Architecture][amd64-shield]

## License

- This app packaging in this repository is licensed under MIT.
- See [LICENSE](../LICENSE).
- Upstream project licensing remains under upstream terms.

[amd64-shield]: https://img.shields.io/badge/amd64-yes-green.svg
