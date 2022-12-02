mod config;
mod routes;
mod models;

fn main() {
    println!("main");
    config::print_config();
    routes::health_route::print_health_route();
    routes::user_route::print_user_route();
}