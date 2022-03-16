use std::io::prelude::*;
use std::os::unix::net::UnixStream;

fn main() -> std::io::Result<()> {
	let mut stream = UnixStream::connect("/tmp/engine.sock")?;
	stream.write_all(b"hello world")?;
	stream.write_all(b"hello world Test")?;
	for data in stream.bytes() {
		match data {
			Err(err) => {
				println!("Error reading from stream {:?}", err);
			}
			Ok(dat) => {
				println!("Reading from stream {}", dat as char);
			}
		}
	}

	Ok(())
}