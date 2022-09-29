extern crate console_error_panic_hook;
extern crate core;

use futures_util::sink::SinkExt;
use futures_util::stream::{SplitSink, SplitStream};
use futures_util::StreamExt;
use wasm_bindgen::prelude::*;
use ws_stream_wasm::{WsErr, WsMessage, WsMeta, WsStream};

#[wasm_bindgen]
extern "C" {
	#[wasm_bindgen(js_namespace = console)]
	fn log(s: &str);
	#[wasm_bindgen(js_namespace = console)]
	fn error(s: &str);
}

#[cfg(feature = "wee_alloc")]
#[global_allocator]
static ALLOC: wee_alloc::WeeAlloc = wee_alloc::WeeAlloc::INIT;

macro_rules! console_log {
    ($($t:tt)*) => (log(&format_args!($($t)*).to_string()))
}
macro_rules! console_err {
    ($($t:tt)*) => (error(&format_args!($($t)*).to_string()))
}


#[wasm_bindgen]
pub async fn run() -> Result<(), JsValue> {
	console_log!("Hello from {}!", "rust");
	set_panic_hook();

	const T: &str = "40283437-12bf-4a85-92fe-d9391223259d";
	let (ws, mut sender, mut receiver) = match open(format!("ws://localhost:6969/ws/{}", T)).await {
		Ok((ws, sender, receiver)) => (ws, sender, receiver),
		Err(_) => {
			return Err(JsValue::from_str("Failed to open websocket"));
		}
	};

	match send(&mut sender, "TESTTEST").await {
		Ok(_) => {}
		Err(_) => {
			return Err(JsValue::from_str("Failed to send message"));
		}
	}

	while let Some(msg) = receiver.next().await {
		console_log!("received message: {:?}", msg);
	}

	console_log!("closed connection");

	Ok(())
}

async fn open(url: String) -> Result<(WsMeta, SplitSink<WsStream, WsMessage>, SplitStream<WsStream>), WsErr> {
	match WsMeta::connect(url, None).await {
		Ok((ws, wsio)) => {
			console_log!("connected to websocket");
			let (sender, receiver): (SplitSink<WsStream, WsMessage>, SplitStream<WsStream>) = wsio.split();
			Ok((ws, sender, receiver))
		}
		Err(e) => {
			console_err!("Error connecting to websocket: {:?}", e);
			Err(e)
		}
	}
}

async fn send(wsio: &mut SplitSink<WsStream, WsMessage>, mess: &str) -> Result<(), WsErr> {
	match wsio.send(WsMessage::Text(mess.to_string())).await {
		Ok(_) => {
			Ok(())
		}
		Err(e) => {
			match e {
				WsErr::ConnectionNotOpen => {
					console_err!("Error sending ws message: Connection Not Open SendFailed: {:?}", e);
					Err(e)
				}
				_ => {
					console_err!("Error sending ws message: {:?}", e);
					Err(e)
				}
			}
		}
	}
}

pub fn set_panic_hook() {
	// When the `console_error_panic_hook` feature is enabled, we can call the
	// `set_panic_hook` function at least once during initialization, and then
	// we will get better error messages if our code ever panics.
	//
	// For more details see
	// https://github.com/rustwasm/console_error_panic_hook#readme
	#[cfg(feature = "console_error_panic_hook")]
	console_error_panic_hook::set_once();
}

//
//fn start_websocket() -> Result<(), JsValue> {
//	// Connect to an echo server
//	let ws = WebSocket::new("wss://echo.websocket.events")?;
////	ws.set_onopen();
////	ws.set_onerror();
////	ws.set_onclose();
//
//	// For small binary messages, like CBOR, Arraybuffer is more efficient than Blob handling
//	ws.set_binary_type(web_sys::BinaryType::Arraybuffer);
//
//	// create callback
//	let cloned_ws = ws.clone();
//	let onmessage_callback = Closure::<dyn FnMut(_)>::new(move |e: MessageEvent| {
//		// Handle difference Text/Binary,...
//		if let Ok(abuf) = e.data().dyn_into::<js_sys::ArrayBuffer>() {
//			console_log!("message event, received arraybuffer: {:?}", abuf);
//			let array = js_sys::Uint8Array::new(&abuf);
//			let len = array.byte_length() as usize;
//			console_log!("Arraybuffer received {}bytes: {:?}", len, array.to_vec());
//			// here you can for example use Serde Deserialize decode the message
//			// for demo purposes we switch back to Blob-type and send off another binary message
//			cloned_ws.set_binary_type(web_sys::BinaryType::Blob);
//			match cloned_ws.send_with_u8_array(&vec![5, 6, 7, 8]) {
//				Ok(_) => console_log!("binary message successfully sent"),
//				Err(err) => console_log!("error sending message: {:?}", err),
//			}
//		} else if let Ok(blob) = e.data().dyn_into::<web_sys::Blob>() {
//			console_log!("message event, received blob: {:?}", blob);
//			// better alternative to juggling with FileReader is to use https://crates.io/crates/gloo-file
//			let fr = web_sys::FileReader::new().unwrap();
//			let fr_c = fr.clone();
//			// create onLoadEnd callback
//			let onloadend_cb = Closure::<dyn FnMut(_)>::new(move |_e: web_sys::ProgressEvent| {
//				let array = js_sys::Uint8Array::new(&fr_c.result().unwrap());
//				let len = array.byte_length() as usize;
//				console_log!("Blob received {}bytes: {:?}", len, array.to_vec());
//				// here you can for example use the received image/png data
//			});
//			fr.set_onloadend(Some(onloadend_cb.as_ref().unchecked_ref()));
//			fr.read_as_array_buffer(&blob).expect("blob not readable");
//			onloadend_cb.forget();
//		} else if let Ok(txt) = e.data().dyn_into::<js_sys::JsString>() {
//			console_log!("message event, received Text: {:?}", txt);
//		} else {
//			console_log!("message event, received Unknown: {:?}", e.data());
//		}
//	});
//	// set message event handler on WebSocket
//	ws.set_onmessage(Some(onmessage_callback.as_ref().unchecked_ref()));
//	// forget the callback to keep it alive
//	onmessage_callback.forget();
//
//	let onerror_callback = Closure::<dyn FnMut(_)>::new(move |e: ErrorEvent| {
//		console_log!("error event: {:?}", e);
//	});
//	ws.set_onerror(Some(onerror_callback.as_ref().unchecked_ref()));
//	onerror_callback.forget();
//
//	let cloned_ws = ws.clone();
//	let onopen_callback = Closure::<dyn FnMut()>::new(move || {
//		console_log!("socket opened");
//		match cloned_ws.send_with_str("ping") {
//			Ok(_) => console_log!("message successfully sent"),
//			Err(err) => console_log!("error sending message: {:?}", err),
//		}
//		// send off binary message
//		match cloned_ws.send_with_u8_array(&vec![0, 1, 2, 3]) {
//			Ok(_) => console_log!("binary message successfully sent"),
//			Err(err) => console_log!("error sending message: {:?}", err),
//		}
//	});
//	ws.set_onopen(Some(onopen_callback.as_ref().unchecked_ref()));
//	onopen_callback.forget();
//
//	Ok(())
//}