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

### Regenerating after a whatsmeow bump

```bash
git clone https://github.com/FelixIsaac/whatsapp-mcp-extended.git && cd whatsapp-mcp-extended
git checkout <WHATSAPP_MCP_REF>
git apply ../patches/0001-bridge-update-whatsmeow.patch
cd whatsapp-bridge && go get go.mau.fi/whatsmeow@latest && go mod tidy && go build ./... && go test ./...
cd .. && git diff -- whatsapp-bridge > ../patches/0001-bridge-update-whatsmeow.patch
```
