# Home Assistant App: Guacamole Client

Clientless remote desktop gateway (Apache Guacamole) for Home Assistant.

## Upstream

| | |
|---|---|
| **Docker image** | [abesnier/docker-guacamole](https://github.com/abesnier/docker-guacamole) · [Docker Hub](https://hub.docker.com/r/abesnier/guacamole) |
| **Releases** | [GitHub Releases](https://github.com/abesnier/docker-guacamole/releases) |
| **Apache Guacamole** | [guacamole.apache.org](https://guacamole.apache.org) · [GitHub](https://github.com/apache/guacamole-server) |

This app uses the `abesnier/guacamole` Docker image, which bundles Apache Guacamole (web frontend), `guacd` (proxy daemon), and SQLite into a single container — making it suitable as a single HA app. The intermediate image is load-bearing; going direct to Apache Guacamole would require running multiple containers.

## Wrapper references

This repository was used as an implementation reference when creating this Home Assistant wrapper:

- [alexbelgium/hassio-addons](https://github.com/alexbelgium/hassio-addons)

![Supports amd64 Architecture][amd64-shield]

## License

- This app packaging in this repository is licensed under MIT.
- See [LICENSE](../LICENSE).
- Upstream project licensing remains under upstream terms.

[amd64-shield]: https://img.shields.io/badge/amd64-yes-green.svg
