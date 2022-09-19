import init from "../wasm";

async function run() {
	console.log("WASM started")
	await init()
	console.log("WASM finished")
}

export {run}