fn fn1() {
    let nums: [u8; 3] = [1, 2, 3];
    let fla = [0.1f64, 0.2, 0.3];
    println!("nums {:?} (2)={}", nums, nums[1]);
    println!("fla {:?}", fla);
}

fn fn2() {
    let num_and_str: (u8, &str, f32) = (40, "s", 0.1);
    println!("{:?}", num_and_str);
    let (num, string, _) = num_and_str;
    println!("{} {}", num, string);
}

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

use std::collections::HashMap;

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

fn main() {
    fn1();
    fn2();
    fn3();
    fn4();
    fn5();
}