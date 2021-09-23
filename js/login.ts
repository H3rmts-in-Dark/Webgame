import '../css/login.css';

function getparams(): Map<string, string> {
	let params = new Map<string, string>()
	if (location.search)
		location.search.substr(1).split("&").forEach(function (item) {
			const s = item.split("=")
			params.set(s[0], decodeURIComponent(s[1]))
		})
	return params
}

function checkloggedin(): string | undefined {
	let ip = getparams().get("IP")
	return (ip == undefined || ip.length == 0) ? undefined : ip
}

function visible(visible: boolean) {
	document.getElementById('loginshadow').style.display = visible ? "block" : "none"
}

function logout() {
	window.open(location.origin, "_self");
}

function addevents() {
	document.getElementById('leavebutton').onclick = () => {
		logout()
	}
}

export {
	checkloggedin,
	visible,
	addevents
}