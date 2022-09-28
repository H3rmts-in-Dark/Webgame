import type {CreateGame} from "src/ts/dto/createGame"
import type {Game} from "src/ts/dto/game"
import {getServerAddress} from "../../ts/addresses";
import {sleep} from "../../ts/util";


async function getGamesFromServer(): Promise<Game[]> {
	return fetch(`${getServerAddress()}/games/all`).then((games) => {
		console.debug(games)
		return games.json() as unknown as Game[]
	}).catch((error) => {
		console.error("getGamesFromServer", error)
		throw error
	})
}

async function getGameFromServer(id: string): Promise<Game> {
	return fetch(`${getServerAddress()}/games/${id}`).then((game) => {
		console.debug(game)
		return game.json() as unknown as Game
	}).catch((error) => {
		console.error("getGameFromServer", error)
		throw error
	})
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

// add check if exists
async function checkAvailable(): Promise<boolean> {
	await sleep(2500)
	return true
}


export {getGamesFromServer, createGameOnServer, getGameFromServer, checkAvailable}
export type {Game, CreateGame}

