# Home Assistant App: RustDesk Server

## What this app does

Runs a self-hosted RustDesk server using:

- `hbbs` (ID and rendezvous service)
- `hbbr` (relay service)

## Basic setup

1. Install the app and start it.
2. Open logs and confirm both `hbbs` and `hbbr` started.
3. Copy the RustDesk public key from the app logs.
4. Configure your RustDesk clients with:
   - ID server: your Home Assistant host or domain
   - Key: the public key shown in app logs

## Options

- `encrypted_only`: if enabled, starts both services with `-k _` to require key-based encryption.
- `relay`: optional hostname or IP that `hbbs` advertises for relay connections.
- `private_key` and `public_key`: optional keypair. Provide both or leave both empty.

If both key options are empty, RustDesk generates a new keypair on first start
under `/config`.

The app logs the active public key during startup so you can paste it directly
into RustDesk clients.

## Exposed ports

- `21115/tcp`: hbbs NAT test and status
- `21116/tcp`: hbbs TCP hole punching
- `21116/udp`: hbbs heartbeat and registration
- `21117/tcp`: hbbr relay
- `21114/tcp`, `21118/tcp`, `21119/tcp`: optional web and extended endpoints

## Notes

- This app uses `host_network: true` for RustDesk networking behavior.
- Keys and server state are stored in app config storage (`/config`).
