# Architecture (Phase 0)

## Overview

JiYuNext is designed as a local-only desktop application:

- Rust (`app/`): host process, desktop UI, lifecycle.
- Go (`core/`): core logic module and action handling.

## Planned ABI Contract

Rust will call Go through a stable C ABI bridge. Planned exported functions:

- `jy_init(config_json)`
- `jy_execute(action_json)`
- `jy_shutdown()`

Command protocol is JSON-based:

- Request: `action`, `payload`, `trace_id`
- Response: `code`, `message`, `result`, `trace_id`

## Current Scaffold

- `app/`: minimal Rust binary to establish build chain.
- `core/`: minimal Go module with request/response primitives and tests.

This scaffold is intentionally small so CI and branch flow can be validated first.
