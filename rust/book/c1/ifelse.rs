fn fn1() {
    let v = true;
    if v {
        println!("True")
    } else {
        println!("False")
    }
}

fn fn2() {
    // 通过 if else 得到值
    let result = if 1 == 2 {
        "1!=2"
    } else {
        "1==2"
    };
    println!("{}", result)
}

fn fn3() {
    let result = if 1 == 2 {
        "1!=2"; // 增加分号后编译器认为用户抛弃该值
    } else {
        "1==2";
    };
    println!("{:?}", result)
}

fn main() {
    fn1();
    fn2();
    fn3();
}