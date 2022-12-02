use std::fmt::Display;

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

    let game = Game;
    game.load(Enemy);
    game.load(Hero);
}

struct Game;

struct Enemy;

struct Hero;

trait Loadable {
    fn init(&self);
}

impl Loadable for Enemy {
    fn init(&self) {
        println!("Load Enemy");
    }
}

impl Loadable for Hero {
    fn init(&self) {
        println!("Load Hero");
    }
}

impl Game {
    // :Loadable 指定特征范围
    fn load<T: Loadable>(&self, entity: T) {
        entity.init();
    }
}

// 限定泛型范围
fn add_things<T: std::ops::Add>(fst: T, snd: T) {
    let _ = fst + snd;
}

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

