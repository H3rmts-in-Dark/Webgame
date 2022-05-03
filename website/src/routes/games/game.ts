import type {CreateGame} from "src/ts/dto/createGame"
import type {Game} from "src/ts/dto/game"
import {getServerAddress, getWebsocketAddress} from "../../ts/addresses";
import {sleep} from "../../ts/util";


async function getGamesFromServer(): Promise<Game[]> {
	let data = await fetch(`${getServerAddress()}/games/all`).then((games) => games.json())
	console.debug(data)
	return data.map((game) => {
		return game as Game
	})
}

// add check of exists
async function getGameFromServer(id: string): Promise<Game> {
	let data = await fetch(`${getServerAddress()}/games/${id}`).then((games) => games.json())
	console.debug(data)
	return data as Game
}

async function createGameOnServer(game: CreateGame) {
	console.debug("creating", game)
	let data = await fetch(`${getServerAddress()}/games/create`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify(game)
	}).then((game) => game.json())
	return data as Game
}

// add check of exists
async function checkAvailable(): Promise<boolean> {
	await sleep(1000)
	return true
}


function buildWebsocket(game: Game): WebSocket {
	let websocket: WebSocket = null
	try {
		websocket = new WebSocket(`${getWebsocketAddress()}/${game.id}`);
		console.log("Connection built");
	} catch(err) {
		console.log("Connection invalid", err);
		return
	}

	websocket.onerror = function(error) {
		console.log("WebSocket Error");
		console.error(error)
	};
	return websocket
}


export {getGamesFromServer, createGameOnServer, getGameFromServer, checkAvailable, buildWebsocket}
export type {Game, CreateGame}

