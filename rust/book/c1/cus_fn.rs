// 类c结构体
struct Player {
    name: String,
    iq: u8,
    friends: u8,
}

impl Player {
    // 关联方法
    fn with_name(name: &str) -> Player {
        Player {
            name: name.to_string(),
            iq: 100,
            friends: 100,
        }
    }

    // 实例方法 可读访问
    fn get_friends(&self) -> u8 {
        self.friends
    }
}

impl Player {
    // 实例方法 可写访问
    fn set_friends(&mut self, count: u8) {
        self.friends = count
    }
}

fn fn1() {
    let mut player = Player::with_name("Dave");
    player.set_friends(23);
    println!("{} {}", player.name, player.get_friends());
    //
    let _ = Player::get_friends(&player);
}

fn main() {
    fn1();
    fn2();
}

enum PaymentMode {
    Debit,
    Credit,
    Paypal,
}

fn pay_by_debit(amt: u64) {
    println!("debit {}", amt);
}

fn pay_by_credit(amt: u64) {
    println!("credit {}", amt);
}

fn pay_by_paypal(amt: u64) {
    println!("paypal {}", amt);
}

impl PaymentMode {
    fn pay(&self, amount: u64) {
        match self {
            PaymentMode::Debit => pay_by_debit(amount),
            PaymentMode::Credit => pay_by_credit(amount),
            PaymentMode::Paypal => pay_by_paypal(amount),
        }
    }
}

fn get_payment_mode() -> PaymentMode {
    PaymentMode::Debit
}

fn fn2() {
    let m = get_payment_mode();
    m.pay(215);
}