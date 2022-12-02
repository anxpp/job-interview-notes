## 泛型

### 创建泛型

#### 泛型函数

```
fn give_me<T>(value:T){
    let _= value;
}

fn main(){
    let a = "";
    let b = 1024;
    give_me(a);
    give_me(b);
}
```

#### 泛型

```
struct GenericStruct<T>(T);

struct Container<T>{
    item :T
}

enum Transmission<T>{
    Signal(T),
    NoSignal
}
```

### 泛型实现

```
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
```

### 泛型应用

```
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
```

## 用特征抽象行为

### 特征

```
struct Audio(String);

struct Video(String);

trait Playable {
    // 实例方法
    fn play(&self);
    // 关联方法
    fn pause() {
        println!("paused")
    }
}


impl Playable for Audio {
    fn play(&self) {
        println!("Audio Play: {}", self.0);
    }
}

impl Playable for Video {
    fn play(&self) {
        println!("Video Play: {}", self.0);
    }
}

trait Vehicle {
    fn price(&self) -> u64;
}

trait Car: Vehicle {
    fn model(&self) -> String;
}

struct TeslaRoadster {
    model: String,
    date: u16,
}

impl TeslaRoadster {
    fn new(model: &str, date: u16) -> Self {
        Self { model: model.to_string(), date }
    }
}

impl Vehicle for TeslaRoadster {
    fn price(&self) -> u64 {
        200_000
    }
}

impl Car for TeslaRoadster {
    fn model(&self) -> String {
        "I".to_string()
    }
}

fn main() {
    println!("super player");
    let a = Audio("mp3".to_string());
    let v = Video("mkv".to_string());
    a.play();
    v.play();

    let r = TeslaRoadster::new("telsa", 2020);
    println!("{} price: {}", r.model, r.price());
}
```

### 特征的多种形式

- 标记特征
- 简单特征
- 泛型特征
- 继承特征

## 使用包含泛型的特征--特征区间

```
struct Game;

struct Enemy;
struct Hero;

trait Loadable{
    fn init(&self);
}

impl Loadable for Enemy{
    fn init(&self) {
        println!("Load Enemy");
    }
}

impl Loadable for Hero{
    fn init(&self) {
        println!("Load Hero");
    }
}

impl Game{
    // :Loadable 指定特征范围
    fn load<T:Loadable> (&self, entity:T){
        entity.init();
    }
}

fn add_things<T: std::ops::Add>(fst: T, snd: T) {
    let _ = fst + snd;
}

fn add_things2<T>(fst: T, snd: T)
    where T: std::ops::Add {
    let _ = fst + snd;
}
```

### 类型上的特征区间

```
struct Foo<T: Display> {
    bar: T,
}

struct Bar<F> where F: Display {
    inner: F,
}
```

### 泛型函数和impl代码块上的特征区间

    用到特征区间最常见的地方

```
use std::fmt::Debug;

trait Eatable {
    fn eat(&self);
}

#[derive(Debug)]
struct Food<T>(T);

#[derive(Debug)]
struct Apple;

impl<T> Eatable for Food<T> where T: Debug {
    fn eat(&self) {
        println!("eating {:?}", self);
    }
}

fn eat<T>(val: T) where T: Eatable {
    val.eat();
}

fn main() {
    let apple = Food(Apple);
    eat(apple);
}
```

### 使用“+”将特征组合为区间

    使用了带默认实现的 trait

```
trait Eat {
    fn eat(&self) {
        println!("eat");
    }
}

trait Code {
    fn code(&self) {
        println!("code");
    }
}

trait Sleep {
    fn sleep(&self) {
        println!("sleep");
    }
}

trait Programmer: Eat + Code + Sleep {
    fn animate(&self) {
        self.eat();
        self.code();
        self.sleep();
        println!("repeat !");
    }
}

struct Bob;
impl Eat for Bob {}
impl Code for Bob {}
impl Sleep for Bob {}
impl Programmer for Bob {}

fn main() {
    Bob.animate();
}
```

### 特征区间与impl特征语法

```
fn show_me<T: Display>(val: T) {
    println!("{}", val);
}

// 使用where限定区间范围
fn add_things2<T>(fst: T, snd: T)
    where T: std::ops::Add {
    let _ = fst + snd;
}

struct Foo<T: Display> {
    bar: T,
}

struct Bar<F> where F: Display {
    inner: F,
}
```

## 标准库特征简介

    #[derive(Default, Debug, PartialEq, Copy, Clone)]

- 内置特征
  - Debug: 该特征有助于在控制台输出类型以便调试
  - PartialEq 和 Eq：允许两个元素相互比较以验证是否相等。
  - Copy 和 Clone：定义了类型的复制方式，Copy依赖Clone
- 其他特征
  - std::ops::Add：允许使用“+”运算符将两个变量相加
  - std::convert 的 Into 和 From：使用户可以根据其他类型创建类型
  - Display：能够输出

```
#[derive(Default, Debug, PartialEq, Copy, Clone)]
struct Complex<T> {
    re: T,
    // 实部
    im: T,  // 虚部
}

// pub trait Add<RHS = Self> {
//     type Output;
//     fn add(self, rhs: RHS) -> Self::Output;
// }

// pub trait From<T> {
//     fn from(self) -> T;
// }

impl<T> Complex<T> {
    fn new(re: T, im: T) -> Self {
        Complex { re, im }
    }
}

impl<T: Add<T, Output=T>> Add for Complex<T> {
    type Output = Complex<T>;
    fn add(self, rhs: Complex<T>) -> Self::Output {
        Complex { re: self.re + rhs.re, im: self.im + rhs.im }
    }
}

use std::convert::From;

impl<T> From<(T, T)> for Complex<T> {
    fn from(value: (T, T)) -> Complex<T> {
        Complex { re: value.0, im: value.1 }
    }
}

impl<T: Display> Display for Complex<T> {
    fn fmt(&self, f: &mut Formatter) -> std::fmt::Result {
        write!(f, "{} + {}i", self.re, self.im)
    }
}

#[cfg(test)]
mod tests {
    use Complex;

    #[test]
    fn complex_basics() {
        let first = Complex::new(3, 5);
        let second: Complex<i32> = Complex::default();
    }

    fn complex_addition() {
        let a = Complex::new(1, -2);
        let b = Complex::default();
        let res = a + b;
        assert_eq!(res, a)
    }
}


fn complex_basics() {
    let first = Complex::new(3, 5);
    let second: Complex<i32> = Complex::default();
}

fn complex_addition() {
    let a = Complex::new(1, -2);
    let b = Complex::default();
    let res = a + b;
    assert_eq!(res, a)
}

fn complex_from() {
    let a = (123, 1234);
    let complex = Complex::from(a);
    println!("{:?}", complex);
}

fn complex_display() {
    let m = Complex::new(234, 2345);
    println!("{}", m);
}
```

## 使用特征对象实现真正的多态性

### 分发

- 静态分发
- 动态分发 

### 特征对象