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

```
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
```

### match 表达式

    类似与switch表达式

```
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
```

### 循环

- loop
- while
- for

```
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
```

### 自定义数据类型

基础类型的包装器：

#### 结构体

```
// 单元结构体
struct Dummy;

fn fn1() {
    let v = Dummy;
}

// 元组结构体
struct Color(u8, u8, u8);

fn fn2() {
    let white = Color(255, 255, 255);
    let red = white.0;
    let green = white.1;
    let blue = white.2;
    println!("{} {} {}", red, green, blue);

    let orange = Color(255, 165, 0);
    // 解构字段
    let Color(r, g, b) = orange;
    println!("{} {} {}", r, g, b);
    // 忽略
    let Color(r, _, b) = orange;
}

// 类c结构体
struct Player {
    name: String,
    iq: u8,
    friends: u8,
    score: u16,
}

fn bump_player_score(mut player: Player, score: u16) {
    player.score += score;
    println!("{} {} {} {}", player.name, player.iq, player.friends, player.score);
}

fn fn3() {
    let name = "name".to_string();
    let player = Player {
        name,
        iq: 171,
        friends: 134,
        score: 1129,
    };
    bump_player_score(player, 120);
}

fn main() {
    fn1();
    fn2();
    fn3();
}
```

#### 枚举

```
#[derive(Debug)]
enum Direction {
    N,
    E,
    S,
    W,
}

enum PlayerAction {
    Move {
        direction: Direction,
        speed: u8,
    },
    Wait,
    Attack(Direction),
}

fn fn4() {
    let s = PlayerAction::Move {
        direction: Direction::N,
        speed: 2,
    };
    match s {
        PlayerAction::Wait => {
            println!("Wait");
        }
        PlayerAction::Move { direction, speed } => {
            println!("move {:?} with {}", direction, speed);
        }
        PlayerAction::Attack(d) => {
            println!("attack to {:?}",d);
        }
    };
}
```

### 类型上的函数和方法

#### 结构体上的 impl 块

```
// 类c结构体
struct Player {
    name: String,
    iq: u8,
    friends: u8,
}

impl Player {
    // 关联方法
    fn with_name(name: &str) -> Player {
        Player {
            name: name.to_string(),
            iq: 100,
            friends: 100,
        }
    }

    // 实例方法 可读访问
    fn get_friends(&self) -> u8 {
        self.friends
    }
}

impl Player {
    // 实例方法 可写访问
    fn set_friends(&mut self, count: u8) {
        self.friends = count
    }
}

fn fn1() {
    let mut player = Player::with_name("Dave");
    player.set_friends(23);
    println!("{} {}", player.name, player.get_friends());
    //
    let _ = Player::get_friends(&player);
}

fn main() {
    fn1()
}
```

#### impl 块和枚举

```
enum PaymentMode {
    Debit,
    Credit,
    Paypal,
}

fn pay_by_debit(amt: u64) {
    println!("debit {}", amt);
}

fn pay_by_credit(amt: u64) {
    println!("credit {}", amt);
}

fn pay_by_paypal(amt: u64) {
    println!("paypal {}", amt);
}

impl PaymentMode {
    fn pay(&self, amount: u64) {
        match self {
            PaymentMode::Debit => pay_by_debit(amount),
            PaymentMode::Credit => pay_by_credit(amount),
            PaymentMode::Paypal => pay_by_paypal(amount),
        }
    }
}

fn get_payment_mode() -> PaymentMode {
    PaymentMode::Debit
}

fn fn2() {
    let m = get_payment_mode();
    m.pay(215);
}
```

### 集合

#### 数组

```
fn fn1() {
    let nums: [u8; 3] = [1, 2, 3];
    let fla = [0.1f64, 0.2, 0.3];
    println!("nums {:?} (2)={}", nums, nums[1]);
    println!("fla {:?}", fla);
}
```

#### 元组

```
fn fn2() {
    let num_and_str: (u8, &str, f32) = (40, "s", 0.1);
    println!("{:?}", num_and_str);
    let (num, string, _) = num_and_str;
    println!("{} {}", num, string);
}
```

#### 列表

```
fn fn3() {
    let mut v: Vec<u8> = Vec::new();
    v.push(1);
    v.push(2);

    let mut v_macro = vec![1];
    v_macro.push(3);
    v_macro.push(4);

    let _ = v_macro.pop();

    let message = if v == v_macro {
        "equal"
    } else {
        "different"
    };
    // different [1, 2] [1, 3]
    println!("{} {:?} {:?}", message, v, v_macro);
}
```

#### hash

```
fn fn4() {
    let mut fruits = HashMap::new();
    fruits.insert("apple", 3);
    fruits.insert("mango", 6);
    fruits.insert("orange", 2);
    fruits.insert("avocado", 7);
    for (k, v) in &fruits {
        println!("{} {}", k, v);
    }
    fruits.remove("orange");
    let old_avo = fruits["avocado"];
    fruits.insert("avocado", old_avo + 5);
    for (k, v) in &fruits {
        println!("{} {}", k, v);
    }
    // mango 6
    // apple 3
    // orange 2
    // avocado 7
    // mango 6
    // apple 3
    // avocado 12
}
```

#### 切片

```
fn fn5() {
    let mut nums = [1, 2, 3, 4];
    {
        let all: &[u8] = &nums[..];
        println!("{:?}", all) // [1, 2, 3, 4]
    }
    {
        let ft: &mut [u8] = &mut nums[1..3];
        ft[0] = 100; // [100, 3]
        println!("{:?}", ft)
    }
    println!("{:?}", nums) // [1, 100, 3, 4]
}
```

### 迭代器