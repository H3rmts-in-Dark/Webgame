type Game = {
	id: number
	limit: number
	name: string
}
type CreateGame = {
	limit: number
	name: string
	visible: boolean
	code: string
}

function getServerAddress(): string {
	// return `https://${location.host.split(':')[0]}:7044`
	return `http://${location.host.split(':')[0]}:5252`

}

async function loadGames(): Promise<Game[]> {
	let data = await fetch(`${getServerAddress()}/games/all`).then((games) => games.json())
	await sleep(500) // to see loading
	console.debug(data)
	return data.map((game) => {
		return game as Game
	})
}

async function create(game: CreateGame) {
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

function hidden() {

}

function connect() {
	return sleep(2000)
}

// add check of exists
async function getGame(id: string): Promise<Game> {
	let data = await fetch(`https://${location.host.split(':')[0]}:7044/games/${id}`)   // location not available on page reload
	// let data = await fetch(`http://localhost:5252/games/${id}`)   // => hardcoded https caused problems
	let json = await data.json()
	console.debug(json)
	return json as Game
}

function buildWebsocket(game: Game): WebSocket {
	let websocket: WebSocket = null
	try {
		websocket = new WebSocket(`ws://localhost:6969/ws/${game.id}`);
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

function sleep(delay: number) {
	return new Promise(resolve => setTimeout(resolve, delay))
}

export {loadGames, create, hidden, getGame, connect, buildWebsocket}
export type {Game, CreateGame}

