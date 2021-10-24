// @ts-ignore
import html from 'HTML/login.html'

import "CSS/login.sass";


let login: HTMLElement

function createLoginDiv() {
	login = document.createElement('div')
	document.body.appendChild(login) // must happen before outerHTML is set because it requires a parent node
	login.outerHTML = html
}

function getParams(): Map<string, string> {
	let params = new Map<string, string>()
	if (location.search)
		location.search.substr(1).split("&").forEach(function (item) {
			const s = item.split("=")
			params.set(s[0], decodeURIComponent(s[1]))
		})
	return params
}

function checkLoggedIn(): string | undefined {
	let ip = getParams().get("IP")
	return (ip == undefined || ip.length == 0) ? undefined : ip
}

function logout() {
	window.open(location.origin, "_self");
}

function addEvents() {
	document.getElementById('leaveButton').onclick = logout
}

export {
	checkLoggedIn,
	addEvents,
	createLoginDiv
}