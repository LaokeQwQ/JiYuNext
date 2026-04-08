mod ffi_bridge;

const APP_NAME: &str = "JiYuNext";
const DEFAULT_CONFIG_FILE: &str = "JiYuNext.json";

fn main() {
    println!("{APP_NAME} bootstrap is ready.");
    println!("Default local config: {DEFAULT_CONFIG_FILE}");

    if std::env::args().any(|arg| arg == "--ffi-probe") {
        match ffi_bridge::call_init(r#"{"mode":"probe"}"#) {
            Ok(resp) => println!("ffi init response: {resp}"),
            Err(err) => println!("ffi init unavailable: {err}"),
        }
    }
}
