use std::fmt::Debug;

#[derive(Debug)]
struct Square(f32);

#[derive(Debug)]
struct Rectangle(f32, f32);

trait Area: Debug {
    fn get_area(&self) -> f32;
}

impl Area for Square {
    fn get_area(&self) -> f32 {
        self.0 * self.0
    }
}

impl Area for Rectangle {
    fn get_area(&self) -> f32 {
        self.0 * self.1
    }
}

// 使用 &dyn Trait 作为参数
fn display_area(item: &dyn Area) {
    println!("{:?}  {}", item, item.get_area());
}

fn main() {
    let l: Vec<&dyn Area> = vec![&Square(3f32), &Rectangle(3f32, 2f32)];
    for s in l {
        // 运行时确定要调配用的方法
        println!("{:?}  {}", s, s.get_area());
        display_area(s);
    }
}