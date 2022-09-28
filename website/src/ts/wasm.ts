import init from "../wasm";

async function run() {
	console.group("WASM started")
	console.time("WASM")
	await init()
	console.groupEnd()
	console.log("WASM finished")
	console.timeEnd("WASM")
}

export {run}