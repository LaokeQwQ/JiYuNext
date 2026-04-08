# Build Notes

## Prerequisites

- Rust toolchain (stable)
- Go 1.23+

## Validate core module

```powershell
cd core
go test ./...
```

## Build Go C ABI bridge

```powershell
cd core
New-Item -ItemType Directory -Path build -Force | Out-Null
go build -buildmode=c-archive -o build/jiyunext_core.a ./bridge
```

## Build Rust app (scaffold)

```powershell
cargo build --manifest-path app/Cargo.toml --release --locked
```

## Build Rust app with FFI bridge (experimental)

```powershell
$env:RUSTFLAGS="-L native=core/build"
cargo build --manifest-path app/Cargo.toml --release --locked --features ffi_bridge
```
