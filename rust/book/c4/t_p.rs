use std::fmt::Display;

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

    show_me("1111111");

    let adder = lazy_adder(34, 354);
    println!("{:?}", adder());
    println!("{}", surround_with_braces("Hello"));
}

fn show_me(val: impl Display) {
    println!("{}", val);
}

fn lazy_adder(a: u32, b: u32) -> impl Fn() -> u32 {
    move || a + b
}

fn surround_with_braces(val: impl Display) -> impl Display {
    format!("{{{}}}", val)
}