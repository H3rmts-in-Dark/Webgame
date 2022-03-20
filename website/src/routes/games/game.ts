type Game = {
	id: number
	limit: string
	name: string
}

async function loadGames(): Promise<Game[]> {
	let data = await fetch(`https://${location.host.split(':')[0]}:7044/games/all`)
	let json = await data.json()
	console.debug(json)
	return json.map((game) => {
		return game as Game
	})
}

function create() {

}

function hidden() {

}

function connect() {
	return new Promise(resolve => setTimeout(resolve, 2000))
}

async function getGame(id: string): Promise<Game> {
	// let data = await fetch(`https://${location.host.split(':')[0]}:7044/games/${id}`)   // location not available on page reload
	let data = await fetch(`http://localhost:5252/games/${id}`)   // => hardcoded https caused probelms
	let json = await data.json()
	console.debug(json)
	return json as Game
}

export {loadGames, create, hidden, getGame, connect}
export type {Game}

