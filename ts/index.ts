// @ts-ignore
import html from 'HTML/index.html'

import "CSS/main.css";


// @ts-ignore
import * as login from './login.ts'

// @ts-ignore
import runwasm from './wasmtest.ts'


(() => {
	console.debug(login)
	let ip = login.checkLoggedin()
	if (ip) {
		console.log(`passed login with ip:${ip}`)
		
		login.addEvents()
		runwasm()
	} else {
		console.log("opening login")
		login.createLoginDiv()
	}
})()