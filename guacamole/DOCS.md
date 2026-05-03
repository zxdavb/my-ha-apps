# Home Assistant App: Guacamole Client

## What this app does

Runs Apache Guacamole as a web-based remote desktop gateway for RDP, VNC, and SSH.

## Access

- Ingress is enabled and recommended.
- The web interface is also exposed on port `8080` (host port `4822` by default).

Default credentials are usually:

- Username: `guacadmin`
- Password: `guacadmin`

Change the default password immediately after first login.

## Options

- `EXTENSIONS`: comma-separated extensions, for example `auth-totp`.
- `recording_search_path`: directory for Guacamole recording search.
- `env_vars`: optional extra environment variables to pass through.
- `TZ`: timezone string such as `Europe/London`.

## Storage

- Guacamole data is stored in `/config` (addon config mapping).
- PostgreSQL data is persisted in `/config/postgres`.

## Notes

This app tracks the direct upstream image published by `abesnier/docker-guacamole`.
