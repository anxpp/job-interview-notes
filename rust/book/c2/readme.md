## 模块

### 嵌套模块

```
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
```

### 将文件用作模块

```
// file_foo.rs
pub struct Bar;

impl Bar{
    pub fn init(){
        println!("Bar init")
    }
}


// file_main.rs
mod file_foo;

// crate 绝对导入：指向当前项目的根目录（root模块，main.rs）
// self  相对导入：指向与当前模块相关的元素
// super 相对导入：可以用于从父模块导入元素
use crate::file_foo::Bar;

fn main() {
    let _ = Bar::init();
}
```

### 将目录用作模块

    详见 rust/book/c2/mod

## Cargo 和程序库

```
cargo new <name>
```