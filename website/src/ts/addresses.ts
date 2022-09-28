function getServerAddress(): string {
	// return `https://${location.host.split(':')[0]}:7044`
	return `http://${location.host.split(':')[0]}:5252`
}

export {getServerAddress}