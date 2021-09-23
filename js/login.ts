import '../css/login.css';

function getparams(): {} {
	var qd = {};
	if (location.search) location.search.substr(1).split("&").forEach(function (item) {
		var s = item.split("="),
			k = s[0],
			v = s[1] && decodeURIComponent(s[1]); //  null-coalescing / short-circuit
		//(k in qd) ? qd[k].push(v) : qd[k] = [v]
		(qd[k] = qd[k] || []).push(v) // null-coalescing / short-circuit
	})
	return qd
}

function checkloggedin(): string | undefined {
	return getparams()["IP"]
}

function visible(visible: boolean) {
	
	let loginshadow = document.getElementById('loginshadow')
	loginshadow.style.display = visible ? "block":"none" 
}


export {
	checkloggedin,
	visible
}