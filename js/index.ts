import '../css/main.css';

import * as im from "./tesfile";
import {testhello} from "./tesfile";

async function main() {
	// @ts-ignore
	const wasm = await import("../webassembly/pkg");
	console.log("wasm: ")
	console.log(wasm);
	
	wasm.greet()
	
	testhello()
	
	im.test3()
}

window.onload = main

im.test2()