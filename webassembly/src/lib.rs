use wasm_bindgen::prelude::*;
use web_sys::console;

#[cfg(feature = "wee_alloc")]
#[global_allocator]
static ALLOC: wee_alloc::WeeAlloc = wee_alloc::WeeAlloc::INIT;

// joinked from
// https://rustwasm.github.io/docs/book/game-of-life/hello-world.html
// https://github.com/rustwasm/rust-webpack-template/tree/master/template

#[wasm_bindgen(start)]
pub fn main_js() -> Result<(), JsValue> {
    
    #[cfg(debug_assertions)]
    console_error_panic_hook::set_once();

    console::log_1(&JsValue::from_str("Hello world!"));

    Ok(())
}

#[wasm_bindgen]
extern {
    fn alert(s: &str);
}

#[wasm_bindgen]
pub fn greet() {
    alert("Hello, Custom created function that uses JS functions!");
}