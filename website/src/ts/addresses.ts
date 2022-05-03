function getServerAddress(): string {
	// return `https://${location.host.split(':')[0]}:7044`
	return `http://${location.host.split(':')[0]}:5252`
}

function getWebsocketAddress(): string {
	// return `wss://${location.host.split(':')[0]}:6969`
	return `ws://${location.host.split(':')[0]}:6969`
}

export {getServerAddress, getWebsocketAddress}