param(
  [switch]$BuildRust
)

$repoRoot = Split-Path -Parent $PSScriptRoot
$coreDir = Join-Path $repoRoot "core"
$appManifest = Join-Path $repoRoot "app/Cargo.toml"
$bridgeOutputDir = Join-Path $coreDir "build"

Push-Location $coreDir
try {
  New-Item -ItemType Directory -Path $bridgeOutputDir -Force | Out-Null
  go build -buildmode=c-archive -o (Join-Path $bridgeOutputDir "jiyunext_core.a") ./bridge
} finally {
  Pop-Location
}

if ($BuildRust) {
  $env:JY_CORE_LIB_DIR = $bridgeOutputDir
  cargo build --manifest-path $appManifest --release --locked --features ffi_bridge
}
