fn give_me<T>(value: T) {
    let _ = value;
}

fn main() {
    let a = "";
    let b = 1024;
    give_me(a);
    give_me(b);

    // 1、提供类型
    let v1: Vec<u8> = Vec::new();
    // 2、不提供类型，但是后续有某个方法调用
    let mut v2 = Vec::new();
    v2.push(2);
    // 3、使用符号
    let v3 = Vec::<u8>::new();

    // 需要第三种方式的时候
    let n = str::parse::<u8>("34").unwrap();
    println!("{}", n);
}

struct GenericStruct<T>(T);

struct Container<T> {
    item: T,
}

impl<T> Container<T> {
    fn new(item: T) -> Self {
        Container { item }
    }
}

// 只会出现在Container<u32>
impl Container<u32> {
    fn sum(item: u32) -> Self {
        Container { item }
    }
}

enum Transmission<T> {
    Signal(T),
    NoSignal,
}