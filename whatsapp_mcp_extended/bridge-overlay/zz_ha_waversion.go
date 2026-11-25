// Added by the Home Assistant app packaging; not part of upstream.
//
// whatsmeow hard-codes the WhatsApp Web client version it announces. When
// WhatsApp retires that version, every connection fails with
// "Client outdated (405)" and no QR/pairing code is ever offered. The patch in
// patches/ updates whatsmeow, but WhatsApp retires versions every few weeks,
// so between whatsmeow updates/rebuilds this keeps the announced version
// current.
//
// This init() runs before main() and sets the announced version, in order:
//  1. WA_CLIENT_VERSION env var ("2.3000.N" or just "N"), if set;
//  2. the current version scraped from web.whatsapp.com (whatsmeow helper);
//  3. haFallbackRevision below, if newer than the library's built-in one.

package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store"
)

// Bumped by the app packaging; matches the whatsmeow commit in
// patches/0001-bridge-update-whatsmeow.patch (2026-11-24).
const haFallbackRevision = 1054751086

func init() {
	builtin := store.GetWAVersion()
	chosen, source := builtin, "whatsmeow built-in"

	if env := strings.TrimSpace(os.Getenv("WA_CLIENT_VERSION")); env != "" {
		parts := strings.Split(env, ".")
		if n, err := strconv.ParseUint(parts[len(parts)-1], 10, 32); err == nil && n > 0 {
			chosen, source = store.WAVersionContainer{2, 3000, uint32(n)}, "WA_CLIENT_VERSION"
		} else {
			fmt.Fprintf(os.Stderr, "[ha] ignoring invalid WA_CLIENT_VERSION %q\n", env)
		}
	}

	if source == "whatsmeow built-in" {
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		latest, err := whatsmeow.GetLatestVersion(ctx, nil)
		cancel()
		if err == nil && latest != nil {
			chosen, source = *latest, "web.whatsapp.com"
		} else {
			fmt.Fprintf(os.Stderr, "[ha] could not fetch current WhatsApp Web version: %v\n", err)
			if builtin[2] < haFallbackRevision {
				chosen, source = store.WAVersionContainer{2, 3000, haFallbackRevision}, "app fallback"
			}
		}
	}

	if chosen != builtin {
		store.SetWAVersion(chosen)
	}
	fmt.Printf("[ha] WhatsApp client version %s (%s; library default %s)\n", chosen, source, builtin)
}
