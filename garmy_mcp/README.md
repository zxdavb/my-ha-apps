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

The add-on runs `garmy-sync` on a configurable schedule to pull data from Garmin Connect into the local SQLite database.

### Option A — credentials in add-on config (recommended)

Set `garmin.email` and `garmin.password` in the add-on configuration. On first start the add-on authenticates with Garmin Connect, saves OAuth tokens to `/data/.garmy/`, and refreshes them automatically. Credentials are only used when the tokens are absent or expired.

### Option B — manual token copy

If you have already run `garmy-sync` on another machine (or prefer not to store credentials in the add-on), copy the two token files instead:

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

3. Leave `garmin.email` and `garmin.password` empty and restart the add-on.

The OAuth2 refresh token typically lasts several months. When it eventually expires the add-on will log a warning and you will need to re-authenticate via Option A or repeat the token copy.

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
