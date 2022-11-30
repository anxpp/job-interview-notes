// 单元结构体
struct Dummy;

fn fn1() {
    let v = Dummy;
}

// 元组结构体
struct Color(u8, u8, u8);

fn fn2() {
    let white = Color(255, 255, 255);
    let red = white.0;
    let green = white.1;
    let blue = white.2;
    println!("{} {} {}", red, green, blue);

    let orange = Color(255, 165, 0);
    // 解构字段
    let Color(r, g, b) = orange;
    println!("{} {} {}", r, g, b);
    // 忽略
    let Color(r, _, b) = orange;
}

// 类c结构体
struct Player {
    name: String,
    iq: u8,
    friends: u8,
    score: u16,
}

fn bump_player_score(mut player: Player, score: u16) {
    player.score += score;
    println!("{} {} {} {}", player.name, player.iq, player.friends, player.score);
}

fn fn3() {
    let name = "name".to_string();
    let player = Player {
        name,
        iq: 171,
        friends: 134,
        score: 1129,
    };
    bump_player_score(player, 120);
}

#[derive(Debug)]
enum Direction {
    N,
    E,
    S,
    W,
}

enum PlayerAction {
    Move {
        direction: Direction,
        speed: u8,
    },
    Wait,
    Attack(Direction),
}

fn fn4() {
    let s = PlayerAction::Move {
        direction: Direction::N,
        speed: 2,
    };
    match s {
        PlayerAction::Wait => {
            println!("Wait");
        }
        PlayerAction::Move { direction, speed } => {
            println!("move {:?} with {}", direction, speed);
        }
        PlayerAction::Attack(d) => {
            println!("attack to {:?}",d);
        }
    };
}

fn main() {
    fn1();
    fn2();
    fn3();
    fn4();
}