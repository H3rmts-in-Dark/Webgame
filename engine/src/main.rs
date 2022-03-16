use std::io::{Read, Write};
use std::os::unix::net::{UnixListener, UnixStream};
use std::path::Path;
use std::thread;

fn handle_client(stream: UnixStream) {
	for data in stream.bytes() {
		match data {
			Err(err) => {
				println!("Error reading from stream {:?}", err);
			}
			Ok(dat) => {
				println!("Reading from stream {} ({})", dat as char, dat);
			}
		}
	}
}

fn handle_client2(mut stream: UnixStream) {
	let result = stream.write_all(b"test");
	match result {
		Ok(..) => {
			println!("ok");
		}
		Err(err) => {
			println!("{}", err);
		}
	}
}

fn main() {
	let path = Path::new("/tmp/engine.sock");

	let listener = UnixListener::bind(path).expect("error creating socket");

	for stream in listener.incoming() {
		match stream {
			Err(err) => {
				println!("Error connecting to socket {:?}", err);
			}
			Ok(stream) => {
				println!("connection to socket");
				let dup = stream.try_clone().unwrap();
				thread::spawn(|| {
					handle_client(stream);
				});
				handle_client2(dup);
			}
		}
	}

	drop(path);
}

fn drop(path: &Path) {
	let _ = std::fs::remove_file(path).unwrap();
}
