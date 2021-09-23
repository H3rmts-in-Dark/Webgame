import '../css/main.css';

import * as login from './login'

async function main() {
	// @ts-ignore
	const wasm = await import("../webassembly/pkg");
	console.log("wasm: ")
	console.log(wasm);
	
	wasm.greet()
}

//window.onload = main

window.onload = () => {
	let ip = login.checkloggedin()
	if (ip) {
		console.log(`passed login with ip:${ip}`)
		login.visible(false)

		main()
	} else {
		console.log("opening login")
		login.visible(true)
	}
}