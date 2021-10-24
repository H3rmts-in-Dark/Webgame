import "CSS/main.css";


// @ts-ignore
import * as login from './login.ts'

// @ts-ignore
import runWASM from './wasmtest.ts'


(() => {
	console.debug(login)
	let ip = login.checkLoggedIn()
	if (ip) {
		console.log(`passed login with ip:${ip}`)
		
		login.addEvents()
		runWASM()
	} else {
		console.log("opening login")
		login.createLoginDiv()
	}
})()