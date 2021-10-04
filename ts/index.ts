// @ts-ignore
import html from 'HTML/index.html'

import "CSS/main.css";


// @ts-ignore
import * as login from './login.ts'

// @ts-ignore
import * as wasm from "../webassembly/pkg";


function wasmtest() {
	console.log(wasm)
	wasm.greet()
}

window.onload = () => {console.log("test loaded")}

document.onload = () => {console.log("document loaded")}

(() => {
	let ip = login.checkLoggedin()
	if (ip) {
		console.log(`passed login with ip:${ip}`)
		
		login.addEvents()
		wasmtest()
	} else {
		console.log("opening login")
		login.createLoginDiv()
	}
})()