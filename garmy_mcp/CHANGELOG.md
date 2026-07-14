# Changelog

## 2.0.0-ha14

- Dropped the blanket `set -x` shell tracing that `advanced.debug_logging` enabled. It was tracing bashio's own internal implementation (every `bashio::log.trace`, `bashio::cache.*`, `bashio::fs.*` call), producing megabytes of log for a few seconds of runtime — not the intent. `advanced.debug_logging` now only enables our own targeted diagnostics (token file existence/size/JSON validity, bashio debug-level messages), which is what was actually useful.

## 2.0.0-ha13

- `garmy-sync-ha` now attempts `AuthClient.refresh_tokens()` before giving up when `is_authenticated` is false. `is_authenticated` only checks whether the (short-lived, ~1h) OAuth2 access token has expired and never refreshes it itself — so a token pasted in shortly after being generated interactively was being treated as unusable even though the OAuth1 token could mint a fresh OAuth2 token without any interactive step. Root-caused via `advanced.debug_logging` output showing well-formed, correctly-seeded token files with `auth.is_authenticated=False`.

## 2.0.0-ha12

- `garmin.oauth1_token_json`/`garmin.oauth2_token_json` are now written to `/data/.garmy/` unconditionally on every start (per field, whenever it's non-blank), instead of only when the matching file didn't already exist. Removes the ambiguity of stale/invalid pre-existing files silently blocking newly-pasted config values; leave a field blank to preserve whatever's already on disk.

## 2.0.0-ha11

- Added `advanced.debug_logging` config option: enables bash `set -x` tracing and Garmin token-file diagnostics (existence/size/JSON validity) across the sync service, MCP service, and token-seeding init step, to help diagnose auth/sync issues.

## 2.0.0-ha10

- Guarded `sync.interval_hours` in the sync loop against a non-numeric value, which was collapsing the sleep to 0 and retrying every couple of seconds instead of waiting between syncs.

## 2.0.0-ha9

- Removed `garmin.email` and `garmin.password`; this add-on runs headlessly and cannot complete Garmin's interactive MFA prompt, so credential-based login never reliably worked for 2FA accounts.
- Added `garmin.oauth1_token_json` and `garmin.oauth2_token_json` config fields instead: paste the raw contents of `oauth1_token.json`/`oauth2_token.json` obtained from an interactive login elsewhere, and a new `cont-init.d` step seeds `/data/.garmy/` with them on first start (skipped once real token files exist).

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
