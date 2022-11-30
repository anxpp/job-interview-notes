fn fn1() {
    // 无限循环
    let mut x = 10;
    loop {
        if x < 0 {
            break;
        }
        println!("{}", x);
        x -= 1
    }
}

fn sub(a: i32, b: i32) -> i32 {
    // 嵌套循环使用标签
    let mut r = 0;
    'increment: loop {
        if r == a {
            let mut dec = b;
            'decrement: loop {
                if dec == 0 {
                    break 'increment;
                } else {
                    r -= 1;
                    dec -= 1;
                }
            }
        } else {
            r += 1;
        }
    }
    r
}

fn fn2() {
    let a = 10;
    let b = 4;
    let r = sub(a, b);
    println!("r={}", r);
}

fn fn3() {
    for i in 0..10 {
        println!("{}", i)
    }
    'b:for i in 0..=10 {
        println!("{}", i)
    }
    let start = 1;
    let end = 10;
    for i in start..end {
        println!("{}", i)
    }
}

fn fn4() {
    let mut x = 10;
    'a: while x >= 0 {
        println!("{}", x);
        x -= 1;
    }
}

fn main() {
    fn1();
    fn2();
    fn3();
    fn4();
}