extern crate console_error_panic_hook;
extern crate core;

use std::panic;

use futures_util::sink::SinkExt;
use wasm_bindgen::prelude::*;
use wasm_bindgen_futures::spawn_local;
use ws_stream_wasm::{WsMessage, WsMeta, WsStream};

#[wasm_bindgen]
extern "C" {
	#[wasm_bindgen(js_namespace = console)]
	fn log(s: &str);
	#[wasm_bindgen(js_namespace = console)]
	fn error(s: &str);
}

macro_rules! console_log {
    ($($t:tt)*) => (log(&format_args!($($t)*).to_string()))
}
macro_rules! console_err {
    ($($t:tt)*) => (error(&format_args!($($t)*).to_string()))
}

const T: &str = "40283437-12bf-4a85-92fe-d9391223259d";

#[wasm_bindgen(start)]
pub fn main() -> Result<(), JsValue> {
	panic::set_hook(Box::new(console_error_panic_hook::hook));

	console_log!("Hello from {}!", "rust");

	console_log!("Hello using web-sys");

	spawn_local(async {
		let (ws, mut wsio): (WsMeta, WsStream) = match WsMeta::connect(format!("ws://localhost:6969/ws/{}", T), None).await {
			Ok(conn) => conn,

			Err(e) => {
				console_err!("Error connecting to server: {:?}", e);
				panic!("Failed to connect to websocket")
			}
		};
		wsio.send(WsMessage::Text("TESTTEST".to_string())).await.expect("Failed to send message");
		ws.close().await.expect("Failed to close websocket");
	});

	Ok(())
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