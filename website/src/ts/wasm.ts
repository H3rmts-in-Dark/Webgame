import init from "../wasm/pkg";

async function load() {
	let wasm = await init()
	wasm.greet()
}

export {load}