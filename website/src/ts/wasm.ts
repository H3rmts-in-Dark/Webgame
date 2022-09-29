import init from "../wasm";
import {run as runWASM} from "../wasm";

async function run() {
	console.log("WASM loading")
	console.time("Load")
	await init()
	console.log("WASM loaded")
	console.timeEnd("Load")

	console.group("WASM started")
	console.time("WASM")
	try {
		await runWASM()
		console.groupEnd()
	} catch(e) {
		console.groupEnd()
		console.error("WASM Error:", e)
	} finally {
		console.log("WASM finished")
		console.timeEnd("WASM")
	}
}

export {run}