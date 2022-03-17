type Game = {
	id: number
	limit: string
	name: string
}

async function loadGames(): Promise<Game[]> {
	let data = await fetch(`${'https://' + location.host.split(':')[0]}:7044/games/all`)
	let json = await data.json()
	console.debug(json)
	return json.map((game) => {
		return game as Game
	})
}

export {loadGames}
