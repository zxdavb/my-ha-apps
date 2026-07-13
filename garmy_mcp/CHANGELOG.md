# Changelog

## 2.0.0-ha8

- Renamed `sync.days` to `sync.history_days` for clarity.

## 2.0.0-ha7

- Removed the `advanced_mcp` group; transport is now fixed to `streamable-http` (no longer user-configurable) and `path` moved to `mcp.path`.

## 2.0.0-ha6

- Removed `http` and `sse` from the `advanced_mcp.transport` options; only `streamable-http` is supported now.

## 2.0.0-ha5

- Removed the `mcp.port` option; the MCP server now always listens on 8000 inside the container, matching the fixed port declared in `ports:`. Changing `mcp.port` previously broke connectivity because Supervisor's port publishing is static and doesn't follow runtime option changes.

## 2.0.0-ha4

- Version bump only; no functional changes.

## 2.0.0-ha3

- Added periodic `garmy-sync` service; syncs Garmin data on startup and on a configurable interval.
- Added `garmin.email` and `garmin.password` config fields; credentials used only to obtain initial OAuth tokens.
- OAuth tokens stored in `/data/.garmy/` and reused across restarts; manual token-file copy also supported.
- Added `sync.interval_hours` and `sync.days` config fields.

## 2.0.0-ha2

- Exposed MCP bind address and port as top-level settings (`mcp.host`, `mcp.port`).
- Kept MCP URL suffix/path in advanced settings (`advanced_mcp.path`).
- Set default SQL database path to `/data/health.db` and kept it user-visible.
- Removed advanced SQL tuning options from the add-on configuration.

## 2.0.0-ha1

- Fixed Docker build failure by adding `git` to the image (required for `pip` install from `git+https` source).

## 2.0.0-ha0

- Initial Home Assistant add-on for garmy, wrapping upstream `garmy 2.0.0` from `zxdavb/garmy`.
- Dockerized on the HA base image with garmy installed into a virtualenv.
- FastMCP network transport support: `streamable-http` (default), `http`, and `sse`.
- Configurable SQL safety settings: row limits, query logging, strict validation.
- Network-only; `stdio` transport is not supported.
