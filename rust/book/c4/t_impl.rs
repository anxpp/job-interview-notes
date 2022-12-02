use std::fmt::{Debug, Display, Formatter};
use std::ops::Add;

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

struct Solution;

impl Solution {
    pub fn min_operations(boxes: String) -> Vec<i32> {
        let mut ans = vec![0; boxes.len()];
        for i in 0..boxes.len() {
            let c = boxes.chars().nth(i).unwrap();
            if c == '1' {
                for j in 0..i {
                    ans[j] += (i - j) as i32;
                }
                for j in i + 1..boxes.len() {
                    ans[j] += (j - i) as i32;
                }
            }
        }
        ans
    }
}

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

fn main() {
    let apple = Food(Apple);
    eat(apple);
    Solution::min_operations("0100110".to_string());

    println!("test Complex");
    complex_basics();
    complex_addition();
    complex_from();
    complex_display();
}