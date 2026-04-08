fn main() {
    let ffi_enabled = std::env::var_os("CARGO_FEATURE_FFI_BRIDGE").is_some();
    if !ffi_enabled {
        return;
    }

    if let Some(dir) = std::env::var_os("JY_CORE_LIB_DIR") {
        println!("cargo:rustc-link-search=native={}", dir.to_string_lossy());
    } else {
        println!("cargo:warning=ffi_bridge enabled but JY_CORE_LIB_DIR is not set");
    }
}
