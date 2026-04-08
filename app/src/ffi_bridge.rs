#[cfg(feature = "ffi_bridge")]
mod enabled {
    use std::ffi::{CStr, CString};
    use std::os::raw::c_char;

    #[link(name = "jiyunext_core")]
    extern "C" {
        fn jy_init(config_json: *const c_char) -> *mut c_char;
        fn jy_execute(action_json: *const c_char) -> *mut c_char;
        fn jy_shutdown() -> *mut c_char;
        fn jy_free(ptr: *mut c_char);
    }

    pub fn call_init(config_json: &str) -> Result<String, String> {
        call_with_arg(jy_init, config_json)
    }

    pub fn call_execute(action_json: &str) -> Result<String, String> {
        call_with_arg(jy_execute, action_json)
    }

    pub fn call_shutdown() -> Result<String, String> {
        let ptr = unsafe { jy_shutdown() };
        take_response(ptr)
    }

    fn call_with_arg(
        func: unsafe extern "C" fn(*const c_char) -> *mut c_char,
        payload: &str,
    ) -> Result<String, String> {
        let c_payload = CString::new(payload).map_err(|_| "payload contains NUL byte".to_string())?;
        let ptr = unsafe { func(c_payload.as_ptr()) };
        take_response(ptr)
    }

    fn take_response(ptr: *mut c_char) -> Result<String, String> {
        if ptr.is_null() {
            return Err("ffi returned null pointer".to_string());
        }
        let text = unsafe { CStr::from_ptr(ptr) }
            .to_str()
            .map_err(|_| "ffi returned non-utf8 response".to_string())?
            .to_owned();
        unsafe { jy_free(ptr) };
        Ok(text)
    }
}

#[cfg(not(feature = "ffi_bridge"))]
mod disabled {
    fn unavailable() -> Result<String, String> {
        Err("ffi_bridge feature is disabled".to_string())
    }

    pub fn call_init(_config_json: &str) -> Result<String, String> {
        unavailable()
    }

    pub fn call_execute(_action_json: &str) -> Result<String, String> {
        unavailable()
    }

    pub fn call_shutdown() -> Result<String, String> {
        unavailable()
    }
}

#[cfg(not(feature = "ffi_bridge"))]
pub use disabled::{call_execute, call_init, call_shutdown};
#[cfg(feature = "ffi_bridge")]
pub use enabled::{call_execute, call_init, call_shutdown};

#[cfg(all(test, not(feature = "ffi_bridge")))]
mod tests {
    use super::{call_execute, call_init, call_shutdown};

    #[test]
    fn disabled_feature_returns_clear_error() {
        let init = call_init("{}").expect_err("ffi should be unavailable without feature");
        let exec = call_execute(r#"{"action":"ping"}"#)
            .expect_err("ffi should be unavailable without feature");
        let shutdown = call_shutdown().expect_err("ffi should be unavailable without feature");

        assert!(init.contains("ffi_bridge"));
        assert!(exec.contains("ffi_bridge"));
        assert!(shutdown.contains("ffi_bridge"));
    }
}
