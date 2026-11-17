# Multiple mailboxes (multi-user mode)

How to use one instance of this app with several Google accounts, including
personal accounts outside your Workspace domain.

## How multi-user mode works

In OAuth 2.1 mode, each MCP connection signs in to Google with its own bearer
token. The server ties every tool call to the account behind that token and
ignores any `user_google_email` the client sends.

**One connector in your MCP client = one Google account.** To reach five
mailboxes, add five connectors, each signed in with a different account.

Claude refuses a second connector with a URL it already has ("A connector
with this URL already exists in your organization"). Give each extra
connector a distinct URL that still reaches the same server, for example
`https://mcp.example.com/mcp?account=personal`. The server ignores the query
string, and the label can be anything. Tested with claude.ai and works. Don't
use a different hostname instead: the
server's OAuth metadata names only `workspace_external_url`, so sign-in would
likely fail.

## 1. App configuration

These are the defaults, so a fresh install needs no changes:

```yaml
workspace_external_url: https://mcp.example.com   # your public URL
authentication:
  user_google_email: ''      # leave blank: the account comes from each sign-in
  single_user: false
oauth21:
  mcp_enable_oauth21: true
  workspace_mcp_oauth_proxy_storage_backend: disk
```

If your saved options still contain an `access:` block with
`single_user`, delete it. It is a legacy option that the app no longer reads.

## 2. Google Cloud OAuth client

In the Google Cloud Console, under **APIs & Services → OAuth consent screen**
(or **Google Auth Platform**):

- **User type: External.** With *Internal*, only accounts in your own
  Workspace domain can sign in. Personal `@gmail.com` accounts and other
  domains are rejected.
- **Publishing status: In production.** While the app is in *Testing*,
  refresh tokens for Gmail and other sensitive scopes expire after 7 days, so
  every account has to sign in again each week. In production, an unverified
  app shows an "unverified app" warning on sign-in that you click past. That is
  fine for personal use, which is capped at 100 users.

Under **Credentials**, on the OAuth client used by this app:

- Add `https://<workspace_external_url>/oauth2callback` to the
  **Authorized redirect URIs**.

## 3. Add each account in your MCP client

For each mailbox, add a custom connector pointing at your server's MCP URL
(with a distinct query string for the second and later connectors, as above)
and complete the Google sign-in with that account.

In Claude's connector settings:

- **Authentication: Sign in now.** You pick the Google account immediately, so
  you can check it's the right one for this connector.
- **OAuth client: Use Claude's published identity** (the default). The server
  supports this (CIMD) as well as *Register automatically* (DCR). Either works
  as long as Claude's callback URLs are in
  `oauth21.workspace_mcp_allowed_client_redirect_uris`, which they are by
  default. *Use your own OAuth client* is not needed.

On the Google sign-in screen, use **Use another account** if Google offers to
continue as an account that's already connected.

## Read-only for some accounts

`tool_selection.workspace_mcp_read_only` and
`tool_selection.workspace_mcp_permissions` apply to the **whole server**, not
to individual accounts. A single instance cannot give one mailbox read/write
access and others read-only access.

To split them, run a second instance with `workspace_mcp_read_only: true`,
its own port, its own `workspace_external_url` and its own credentials
directory. Then connect the read-only accounts to that instance only. In
read-only mode the server requests only read-only scopes at sign-in, so Google
enforces the restriction as well. Add the second instance's
`/oauth2callback` URI to the same OAuth client.

## Removing an account

Delete its connector in your MCP client and revoke the app's access in that
account's Google settings (**Security → Third-party apps & services**).
Otherwise the server keeps a working sign-in for it.
