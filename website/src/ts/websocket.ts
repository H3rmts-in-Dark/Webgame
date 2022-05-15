import type {Game} from "./dto/game";
import {getWebsocketAddress} from "./addresses";
import {load} from "./wasm";

function buildWebsocket(game: Game, onopen: () => void, onmessage: (mess: string) => void, onerror: (err: Event) => void, onclose): WebSocket {
	let websocket: WebSocket = null
	try {
		websocket = new WebSocket(`${getWebsocketAddress()}/ws/${game.id}`);
		console.log("Connection built");
	} catch(err) {
		console.log("Connection invalid", err);
		return
	}
	websocket.onmessage = async (mess) => {
		console.debug(mess)
		console.timeEnd("ws")
		await process(mess.data)
		onmessage(mess.data)
	}
	websocket.onopen = (ev: Event) => {
		onopen()
	}
	websocket.onerror = (error: Event) => {
		console.log("WebSocket Error");
		console.error(error)
		onerror(error)
	};
	websocket.onclose = (close: CloseEvent) => {
		onclose()
	}
	return websocket
}

async function process(data: string) {
	switch(data) {
		case "Start":
			console.log("INIT")
			load().then(() => {
				console.log("Loaded")
			})
			break
		default:
			console.log("process failed")
	}
}

export {
	buildWebsocket
}