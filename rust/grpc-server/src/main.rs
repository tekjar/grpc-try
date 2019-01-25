extern crate grpcio;
extern crate futures;
extern crate protobuf;
#[macro_use]
extern crate log;

use pretty_env_logger;

mod server;

fn main() {
    pretty_env_logger::init();
    info!("Hello world");

    server::GreeterService::start()
}
