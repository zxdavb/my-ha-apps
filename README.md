# Home Assistant App Repository

This repository contains local Home Assistant apps that wrap upstream server projects.

 - each app is maintained independently and tracks its upstream source.
 - apps documentation: <https://developers.home-assistant.io/docs/apps>

[![Open your Home Assistant instance and show the app store with a specific repository URL pre-filled.](https://my.home-assistant.io/badges/supervisor_store.svg)](https://my.home-assistant.io/redirect/supervisor_store/?repository_url=https%3A%2F%2Fgithub.com%2Fzxdavb%2Fmy-ha-apps)

Workflow source note: when syncing CI workflows, use `zxdavb/my-ha-apps` as the source of truth (not `home-assistant/apps-example`) to preserve local builder normalization behavior.

## Apps

This repository contains the following apps

### [Google Workspace MCP](./google_workspace_mcp)

![Supports amd64 Architecture][amd64-shield]

_Google Workspace MCP server app based on the upstream project by taylorwilsdon._

Upstream: [taylorwilsdon/google_workspace_mcp](https://github.com/taylorwilsdon/google_workspace_mcp)

### [Guacamole Client](./guacamole)

![Supports amd64 Architecture][amd64-shield]

_Clientless remote desktop gateway (Apache Guacamole) for Home Assistant._

Upstream: [Apache Guacamole](https://guacamole.apache.org) via [abesnier/docker-guacamole](https://github.com/abesnier/docker-guacamole)

### [RustDesk Server](./rustdesk-server)

![Supports amd64 Architecture][amd64-shield]

_Self-hosted RustDesk server app (hbbs + hbbr) for Home Assistant._

Upstream: [rustdesk/rustdesk-server](https://github.com/rustdesk/rustdesk-server)

## License

This repository is licensed under the MIT License. See [LICENSE](./LICENSE).

Upstream projects used by individual apps keep their own licenses.

[amd64-shield]: https://img.shields.io/badge/amd64-yes-green.svg
