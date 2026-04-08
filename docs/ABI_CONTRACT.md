# ABI Contract (Draft)

This document defines the current draft contract between Rust host (`app/`) and Go core bridge (`core/bridge/`).

## Exported C ABI

```c
char* jy_init(const char* config_json);
char* jy_execute(const char* action_json);
char* jy_shutdown(void);
void  jy_free(char* ptr);
```

## Memory Rule

- Every non-null string returned by `jy_init/jy_execute/jy_shutdown` is heap-allocated on the Go side.
- Rust caller must release it with `jy_free`.

## Request JSON

`jy_execute` expects:

```json
{
  "action": "string-required",
  "payload": {},
  "trace_id": "string-optional"
}
```

## Response JSON

All API responses follow:

```json
{
  "code": 0,
  "message": "ok",
  "result": {},
  "trace_id": "string-optional"
}
```

## Status Codes (current)

- `0`: success
- `1001`: invalid input
- `2000`: internal error / panic fallback

## Notes

- This contract is versioned with source changes.
- Any breaking field/schema change should update this document and add compatibility notes.
