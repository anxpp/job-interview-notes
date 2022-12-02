mod food {
    // 使用pub公开
    pub struct Cake;
    struct Smoothie;
    struct Pizza;
}

// 使用use引入
use food::Cake;

fn main() {
    let _ = Cake;
}