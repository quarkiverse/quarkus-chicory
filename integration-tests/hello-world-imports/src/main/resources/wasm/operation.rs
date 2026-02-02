
#[link(wasm_import_module = "env")]
extern "C" {
    fn host_log(x: i32);
}

#[no_mangle]
pub extern fn operation(a: i32, b: i32) -> i32 {
    unsafe { host_log(a) };
    unsafe { host_log(b) };
    a + b
}
