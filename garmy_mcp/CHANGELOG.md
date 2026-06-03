# Changelog

## 2.0.0-ha1

- Fixed Docker build failure by adding `git` to the image (required for `pip` install from `git+https` source).

## 2.0.0-ha0

- Initial Home Assistant add-on for garmy, wrapping upstream `garmy 2.0.0` from `zxdavb/garmy`.
- Dockerized on the HA base image with garmy installed into a virtualenv.
- FastMCP network transport support: `streamable-http` (default), `http`, and `sse`.
- Configurable SQL safety settings: row limits, query logging, strict validation.
- Network-only; `stdio` transport is not supported.
