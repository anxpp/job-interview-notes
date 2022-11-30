fn req_status() -> u32 {
    404
}

fn main() {
    let status = req_status();
    match status {
        200 => println!("Success"),
        404 => println!("Not found"),
        oth => {
            println!("other {}", oth)
        }
    }
    // 同样可以返回值使用
    let code = match status {
        200 => "Success",
        404 => "Not fount",
        _other => "other"
    };
    println!("code = {}", code)
}