# Home Assistant App: Garmy MCP

Garmin health data MCP server app for Home Assistant.

## Upstream

| | |
| --- | --- |
| **Project** | [zxdavb/garmy](https://github.com/zxdavb/garmy) (fork of [bes-dev/garmy](https://github.com/bes-dev/garmy)) |
| **PyPI package** | [garmy](https://pypi.org/project/garmy/) |
| **Releases** | [bes-dev/garmy releases](https://github.com/bes-dev/garmy/releases) |

This app installs `garmy[all]` from the fork repository into a virtualenv on top of the HA base image, then starts the `garmy-mcp` server.

## Garmin Connect authentication

The add-on runs `garmy-sync` on a configurable schedule to pull data from Garmin Connect into the local SQLite database. This add-on does **not** accept a Garmin email/password, since it starts headlessly and has no way to complete Garmin's interactive MFA prompt for 2FA-protected accounts. Authentication is token-only — choose one of the two options below.

### Option A — paste tokens via the config UI (recommended)

Authenticate interactively once on any machine where you can enter an MFA code (a laptop, a dev container, etc.), then hand the resulting tokens to the add-on through its configuration — no SSH/Samba access to the Home Assistant host required.

1. On that machine, install `garmy` and run its sync/auth flow (or any `garmy` login helper) so it prompts for your MFA code and completes login. This produces two files, of which you only need the first:

   ```text
   ~/.garmy/oauth1_token.json
   ~/.garmy/oauth2_token.json   (not needed — see below)
   ```

2. Open `oauth1_token.json` and paste its contents as-is into `garmin.oauth1_token_json` in the add-on configuration. Start (or restart) the add-on.

You don't need `oauth2_token.json` at all: OAuth1 is the durable "login" credential, and `garmy-sync` mints its own short-lived OAuth2 access token from it on demand (`AuthClient.refresh_tokens()`), refreshing it again whenever it expires. Only the OAuth1 token ever needs to come from you.

`garmin.oauth1_token_json` also doubles as a change signal, so the add-on doesn't fight with that refresh cycle:

- **First start, or `garmin.oauth1_token_json` has changed** (you pasted a new login): `oauth1_token.json` is (re)written from config, and any existing `oauth2_token.json` is removed so a fresh one gets minted to match.
- **`garmin.oauth1_token_json` is unchanged** since the last start: `/data/.garmy/` isn't touched, so the OAuth2 token `garmy-sync` has refreshed in the meantime survives the restart.
- **`garmin.oauth1_token_json` is blank**: `/data/.garmy/` is left entirely alone (see Option B).

### Option B — manual token copy

If you have already run `garmy-sync` on another machine, copy the token file(s) directly instead of pasting into the config UI. Make sure `garmin.oauth1_token_json` is blank, since a non-blank/changed value overwrites `oauth1_token.json`:

1. On the machine where you ran `garmy-sync`, locate the token directory (default `~/.garmy/`):

   ```text
   ~/.garmy/oauth1_token.json
   ~/.garmy/oauth2_token.json
   ```

2. Copy both files into the add-on's persistent data volume on your Home Assistant host:

   The easiest path via SSH or Samba is:

   ```text
   /mnt/data/supervisor/addons/data/local_garmy_mcp/.garmy/oauth1_token.json
   /mnt/data/supervisor/addons/data/local_garmy_mcp/.garmy/oauth2_token.json
   ```

   (The exact path varies by HA installation type; adjust `local_garmy_mcp` to match your add-on slug.)

3. Restart the add-on.

The OAuth1 login typically stays valid for months. When it eventually expires, `garmy-sync`'s refresh will start failing, the add-on will log a warning, and you'll need to re-authenticate (interactively, elsewhere) and repeat Option A or B.

### Troubleshooting

Set `advanced.debug_logging: true` in the add-on configuration and restart to get token-file diagnostics (existence, size, whether each file parses as JSON) in the add-on log. Leave it off otherwise — its output includes token file paths/sizes.

## Authentication and External Proxy (NPM)

Recommended model: keep auth outside this app and front it with Nginx Proxy Manager (NPM).

### Why this model

- Keeps MCP runtime and auth stack separate.
- Makes OAuth policy and provider changes easier without rebuilding this app.
- Reduces security and maintenance burden inside the add-on container.

### Setup steps

1. Configure this app for URL transport:
   - `database.path`: your SQLite DB file path.
   - `mcp.host`: `0.0.0.0`.
   - `mcp.path`: `/mcp`.

2. In NPM, create a proxy host:
   - Domain: your public MCP hostname, for example `mcp.example.com`.
   - Forward scheme: `http`.
   - Forward host/IP: Home Assistant/add-on reachable address.
   - Forward port: app port (`8000` by default).

3. Enable TLS in NPM (Let's Encrypt) and force HTTPS.

4. Add authentication at the proxy layer:
   - Minimum: NPM Access List (basic auth).
   - Preferred: OAuth/OIDC-capable forward auth in front of NPM (or via custom Nginx auth_request integration).

5. Point MCP clients to your public MCP URL, for example `https://mcp.example.com/mcp`.

### NPM tuning notes

For long-lived MCP HTTP sessions, consider adding advanced Nginx directives in NPM:

```nginx
proxy_buffering off;
proxy_request_buffering off;
proxy_read_timeout 3600;
proxy_send_timeout 3600;
```

These help reduce issues with streaming responses and long-running requests.

## Wrapper references

No intermediary wrapper repository is used for this app. Packaging is maintained directly in this repository against your fork.

![Supports amd64 Architecture][amd64-shield]

## License

- This app packaging in this repository is licensed under MIT.
- See [LICENSE](../LICENSE).
- Upstream project licensing remains under upstream terms.

[amd64-shield]: https://img.shields.io/badge/amd64-yes-green.svg
