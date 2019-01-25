#[macro_use]
extern crate log;

use futures::stream::Stream;
use futures::Future;
use std::net::SocketAddr;
use tokio_codec::LinesCodec;
use tokio::runtime::current_thread;
use tokio::net::{UdpSocket, UdpFramed};
use pretty_env_logger;

fn framed() -> impl Future<Item = String, Error = io::Error> {
    let addr: SocketAddr = "127.0.0.1:0".parse().unwrap();
    
    let a = UdpSocket::bind(&addr).unwrap();
    let (sink, stream) = UdpFramed::new(a, LinesCodec::new()).split();

    stream.map(|v| {
        info!("Incoming: {:?}", v);
        let () = v;
    })
    .forward(sink)
}

fn main() {
    pretty_env_logger::init();

}
