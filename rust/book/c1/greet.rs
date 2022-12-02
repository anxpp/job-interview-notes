use std::env;

fn main() {
    let name: Option<String> = env::args().skip(1).next();
    match name {
        Some(n) => println!("Hi there ! {}", n),
        None => panic!("Didn't receive any name ?")
    }
}