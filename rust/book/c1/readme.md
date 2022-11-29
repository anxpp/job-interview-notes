## RUST简介

### 基元类型

- bool: true false
- char: 
- 整形:

| 有符号  | 无符号  |
|------|------|
| i8   | u8   |
| i16  | u16  |
| i32  | u32  |
| i64  | u64  |
| i128 | u128 |

- isize: 尺寸可变的有符号整形
- usize：尺寸可变的无符号整形
- f32:
- f64:
- \[T; N]: 固定大小的数组，T为元素类型，N为数目
- \[T]：动态大小的连续序列的试图，T表示类型
- str：字符串切片，主要用作引用：&str
- \(T,U,..)：有限序列，T、U可以是不同的类型
- fn\(i32)->i32：一个接收i32类型参数并返回i32类型参数的函数·

### 变量声明和不可变性

- 不可变的绑定变量：`let v = "";`
- 可变的绑定变量：`let mut v = "";`

### 函数

```
fn add(a: u64, b: u64) -> u64 {
    a + b
}

fn main() {
    let a: u64 = 17;
    let b = 3;
    let r = add(a, b);
    println!("Result {}", r);
}
```

```
fn increase_by(mut val: u32, how_much: u32) {
    val += how_much;
    println!("You made {} points", val)
}

fn main() {
    let score = 2048;
    increase_by(score, 30)
}
```

### 闭包

```
fn main() {
    // 闭包1
    let doubler = |x| x * 2;
    let value = 5;
    let twice = doubler(value);
    println!("{} doubled is {}", value, twice);

    // 闭包2
    let big_closure = |b, c| {
        let z = b + c;
        z * twice
    };

    let some_number = big_closure(1, 2);
    println!("result from closure: {}", some_number);
}
```

### 字符串

```
fn main() {
    let question = "How are you?";  // &str类型
    let person: String = "Bob".to_string(); //  String类型
    let namaste = String::from("123123"); // String类型
    println!("{}! {} {}", namaste, question, person);
}
```

### 条件和判断