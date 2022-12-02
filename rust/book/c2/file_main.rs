mod file_foo;

// crate 绝对导入：指向当前项目的根目录（root模块，main.rs）
// self  相对导入：指向与当前模块相关的元素
// super 相对导入：可以用于从父模块导入元素
use crate::file_foo::Bar;

fn main() {
    let _ = Bar::init();
}