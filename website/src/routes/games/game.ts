import type {CreateGame} from "src/ts/dto/createGame"
import type {Game} from "src/ts/dto/game"
import {getServerAddress} from "../../ts/addresses";
import {sleep} from "../../ts/util";


async function getGamesFromServer(): Promise<Game[]> {
	let data = await fetch(`${getServerAddress()}/games/all`).then((games) => games.json())
	console.debug(data)
	return data.map((game) => {
		return game as Game
	})
}

// add check if exists
async function getGameFromServer(id: string): Promise<Game> {
	let data = await fetch(`${getServerAddress()}/games/${id}`).then((games) => games.json())
	console.debug(data)
	await sleep(300) // TODO just test
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
	await sleep(2500)
	return true
}


export {getGamesFromServer, createGameOnServer, getGameFromServer, checkAvailable}
export type {Game, CreateGame}

