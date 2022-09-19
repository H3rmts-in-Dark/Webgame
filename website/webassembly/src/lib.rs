use wasm_bindgen::prelude::*;
use web_sys::{Document, Element};

#[wasm_bindgen]
extern "C" {
	#[wasm_bindgen(js_namespace = console)]
	fn log(s: &str);
}

#[wasm_bindgen(start)]
pub fn main() -> Result<(), JsValue> {
	use web_sys::console;

	log(&format!("Hello from {}!", "rust"));

	console::log_1(&"Hello using web-sys".into());

	let js: JsValue = 4.into();
	console::log_2(&"Logging arbitrary values looks like".into(), &js);

	// Use `web_sys`'s global `window` function to get a handle on the global
	// window object.
	let window = web_sys::window().expect("no global `window` exists");
	let document: Document = window.document().expect("should have a document on window");

	// Manufacture the element we're going to append
	let val = document.create_element("p")?;
	val.set_inner_html("Hello from Rust!");

	let f: Option<Element> = document.get_element_by_id("123");
	f.unwrap().append_child(&val)?;

	Ok(())
}