import {getServerAddress} from "./addresses";
import type {CheckCodes} from "./dto/checkCodes";
import type {Game} from "./dto/game";
import type {CreateGame} from "./dto/createGame";

async function getGamesFromServer(): Promise<Game[]> {
	console.log("getGamesFromServer")
	return fetch(`${getServerAddress()}/games/all`).then(async (response) => {
		let json = await response.json()
		console.debug(response)
		console.log(json)
		return json as unknown as Game[]
	}).catch((error) => {
		console.error("getGamesFromServer", error)
		throw error
	})
}

async function getGameFromServer(id: string): Promise<Game> {
	console.log("getGameFromServer", id)
	return fetch(`${getServerAddress()}/games/${id}`).then(async (response) => {
		let json = await response.json()
		console.debug(response)
		console.log(json)
		return json as unknown as Game
	}).catch((error) => {
		console.error("getGameFromServer", error)
		throw error
	})
}

async function createGameOnServer(game: CreateGame) {
	console.log("createGameOnServer", game)
	return await fetch(`${getServerAddress()}/games/create`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify(game)
	}).then(async (response) => {
		let json = await response.json()
		console.debug(response)
		console.log(json)
		return json as unknown as Game
	}).catch((error) => {
		console.error("createGameOnServer", error)
		throw error
	})
}

async function check(id: number, code: string): Promise<CheckCodes> {
	console.debug("check", id)
	return await fetch(`${getServerAddress()}/games/${id}/check?code=${encodeURIComponent(code) || ''}`).then(async (response) => {
		let json = await response.json()
		console.debug(response)
		console.log(json)
		return json as CheckCodes
	}).catch((error) => {
		console.error("check", error)
		return error
	})
}


export {getGamesFromServer, createGameOnServer, getGameFromServer, check}
