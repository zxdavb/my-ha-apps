# Packaging patches

Applied in order (`git apply`) to the upstream source at `WHATSAPP_MCP_REF`
during the image build. The build fails if a patch no longer applies.

## 0001-bridge-update-whatsmeow.patch

Updates the bridge's `go.mau.fi/whatsmeow` dependency (upstream pins
2026-07-22, which WhatsApp rejects with `Client outdated (405)`) and ports
the three call sites whose whatsmeow API changed:

- `GetBlocklist` / `UpdateBlocklist` take a `dhash` argument (empty = full
  list / no concurrency check); the blocklist now has `Items` with LID and
  phone-number JIDs instead of `JIDs`. The bridge reports the phone-number
  JID when known, else the LID.
- `SetStatusMessage` takes `types.SetStatusInput` and uses a newer GraphQL
  mutation that some accounts reject (tulir/whatsmeow#1331). `SetAboutText`
  falls back to the legacy status IQ that older whatsmeow used.

Intended to be offered upstream (FelixIsaac/whatsapp-mcp-extended). Drop it
once upstream carries an equivalent change.

## 0002-bridge-qr-refresh.patch

Applies on top of 0001. Upstream's `Connect()` reads one whatsmeow QR channel,
which ends after ~6 codes (about 2.5 minutes). It then waited up to 15 minutes
before marking itself disconnected, and the bridge watchdog restarted it about
3 minutes later. This loops to request a fresh QR channel until paired, unless
phone-code pairing is in progress (that path is unchanged). It also checks
the `GetQRChannel` error, which upstream ignored (a nil channel would block
forever), and logs QR events other than `code`/`success`/`timeout`. When
`WA_QR_FILE` is set, it writes the current QR code string there (removed once
the set ends) for the app's pairing page.

## 0003-bridge-log-privacy.patch

Applies on top of 0002. Upstream hard-codes the bridge log level to INFO and
logs, at INFO, every history-synced message in full ("Stored message: ...",
"Message content: ..."), the contact name/number behind each incoming message
("Getting name for contact", "Using contact name"), group renames and calls.
These go to DEBUG, and `BRIDGE_LOG_LEVEL` (DEBUG/INFO/WARN/ERROR, default INFO,
anything else falls back to INFO) sets the level. The app sets DEBUG only
when `advanced.debug_logging` is on.

### Regenerating after a whatsmeow bump

```bash
git clone https://github.com/FelixIsaac/whatsapp-mcp-extended.git && cd whatsapp-mcp-extended
git checkout <WHATSAPP_MCP_REF>
git apply ../patches/0001-bridge-update-whatsmeow.patch  # then commit, so 0002 can be regenerated separately
cd whatsapp-bridge && go get go.mau.fi/whatsmeow@latest && go mod tidy && go build ./... && go test ./...
cd .. && git diff -- whatsapp-bridge > ../patches/0001-bridge-update-whatsmeow.patch
```
