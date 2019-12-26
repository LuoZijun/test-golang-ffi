

#[no_mangle]
pub unsafe extern "C" fn rust_main(ptr: *const u8, len: usize) -> i32 {
    let bytes = std::slice::from_raw_parts(ptr, len);
    println!("[RUST]: {:?}", bytes);
    println!("[RUST]: LEN={:?}", bytes.len());
    
    return 0;
}

