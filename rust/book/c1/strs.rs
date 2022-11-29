fn main() {
    let question = "How are you?";  // &str类型
    let person: String = "Bob".to_string(); //  String类型
    let namaste = String::from("123123"); // String类型
    println!("{}! {} {}", namaste, question, person);
}