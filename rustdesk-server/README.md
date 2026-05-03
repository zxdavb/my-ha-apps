# Home Assistant App: RustDesk Server

Self-hosted RustDesk server app (hbbs + hbbr) for Home Assistant.

## Upstream

| | |
|---|---|
| **Project** | [rustdesk/rustdesk-server](https://github.com/rustdesk/rustdesk-server) |
| **Releases** | [GitHub Releases](https://github.com/rustdesk/rustdesk-server/releases) |
| **Changelog** | [CHANGELOG.md](https://github.com/rustdesk/rustdesk-server/blob/master/CHANGELOG.md) |

This app downloads the `hbbs` (ID/rendezvous server) and `hbbr` (relay server) binaries directly from upstream GitHub releases and runs both processes under s6-overlay. No intermediate Docker image is used.

## Wrapper references

These repositories were used as implementation references when creating this Home Assistant wrapper:

- [SankeerthBoddu/ha-rustdesk-server](https://github.com/SankeerthBoddu/ha-rustdesk-server)
- [casse-boubou/hassio-addons](https://github.com/casse-boubou/hassio-addons)

![Supports amd64 Architecture][amd64-shield]

## License

- This app packaging in this repository is licensed under MIT.
- See [LICENSE](../LICENSE).
- Upstream project licensing remains under upstream terms.

[amd64-shield]: https://img.shields.io/badge/amd64-yes-green.svg
