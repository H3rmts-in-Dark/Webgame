<script lang="ts">
	import {loadGames, create, hidden, connect} from "./game"
	import type {Game} from "./game";
	import Button from "@smui/button";

	import {mdiLoading} from "@mdi/js";
	import SvgIcon from "$lib/SvgIcon.svelte";
	import Open from "$lib/open.svelte"
	import Textfield from "@smui/textfield";
	import Slider from "$lib/Slider.svelte";

	type GameDisplay = {
		game: Game
		buttonDisplay: string
	}

	let creatingGame = false
	let newGame: Game = null

	// Array if loaded, null if loading, string if error
	let games: Array<GameDisplay> | null | string = null

	async function load() {
		creatingGame = false
		games = null
		games = await loadGames()
				.then((games: Array<Game>) => {
					return games.map((game: Game) => {
						return {game: game, buttonDisplay: "Join"}
					})
				})
				.catch((err: Error) => {
					return err.message
				})
	}

	load()

	async function join(gameDisplay: GameDisplay) {
		console.log(`joining game`)
		gameDisplay.buttonDisplay = "Load"

		games = games // forces rerender

		await connect()

		gameDisplay.buttonDisplay = "__OpenLink"
		games = games // forces rerender
	}

	async function createGame() {
		creatingGame = !creatingGame
		if(creatingGame) {
			newGame = {limit: 0, name: ""}
		}
	}
</script>

<svelte:head>
	<title>Games</title>
</svelte:head>

<div style="display: flex; flex-direction: column; align-items: center">
	<h0>Games</h0>
	<div id="buttons_bar">
		<div>
			<Button variant="outlined" color="primary" on:click={createGame}>
				{creatingGame ? "List" : "Create"}
			</Button>
		</div>
		<div>
			<Button variant="outlined" color="primary" on:click={hidden}>
				Hidden
			</Button>
			<Button variant="outlined" color="primary" on:click={load}>
				Scan
			</Button>
		</div>
	</div>
	{#if creatingGame}
		<div id="createGame">
			<h1 class="title">Create new Game</h1>
			<div class="field">
				<Textfield style="width: 45%;" class="shaped-outlined" color="secondary" variant="outlined" bind:value={newGame.name}
							  label="Name"/>
				<Slider bind:data={newGame.limit} name="Value: "></Slider>
			</div>
			<div class="field">
				<Button variant="raised" color="secondary" on:click={() => {create(newGame)}}>
					Create
				</Button>
			</div>
		</div>
	{:else}
		<div id="games">
			{#if games == null}
				<h2>Loading</h2>
			{:else if typeof games == "string"}
				<h2>Error: {games}</h2>
			{:else}
				{#each games as game}
					<div class="game">
						<h2>{game.game.name}</h2>
						<h3 class="players">
							? / {game.game.limit}
						</h3>
						{#if game.buttonDisplay === "Load"}
							<Button variant="raised" color="secondary">
								<SvgIcon cls="rotate" svg={mdiLoading}/>
							</Button>
						{:else if game.buttonDisplay === "__OpenLink"}
							<Open link={`games/${game.game.id}`}/>
						{:else}
							<Button variant="raised" color="secondary" on:click={() => {join(game)}}>
								{game.buttonDisplay}
							</Button>
						{/if}
					</div>
				{/each}
			{/if}
		</div>
	{/if}
</div>

<style lang="scss">
	@import "src/css/vars";

	#buttons_bar {
		display: flex;
		width: calc(100% - 50px);
		margin-right: 25px;
		margin-left: 25px;

		justify-content: space-between;

		div {
			display: flex;
			gap: 15px;
		}
	}

	#games {
		margin: 15px;
		display: grid;
		grid-gap: 20px;
		width: calc(100% - 30px);

		grid-template-columns: repeat(auto-fill, minmax(300px, 3fr));
	}

	#createGame {
		margin-block: 3.5em;
		padding: 10px;
		padding-inline: 3em;
		width: calc(80%);

		border: {
			style: solid;
			width: 6px;
			color: $on-primary;
		}

		.field {
			align-items: center;
			justify-content: center;
			display: flex;
			margin-block: 1.2em;
			gap: 2em;
		}

		.title {
			text-align: center;
		}
	}

	.game {
		position: relative;
		background: $primary;
		padding: 10px;
		margin: 10px;
		color: $on-primary;

		border: {
			style: solid;
			width: 6px;
			color: $on-primary;
		}
		display: flex;
		flex-direction: column;
		align-items: center;

		.players {
			margin: 0;
			position: absolute;
			right: 7px;
			top: 5px;
		}
	}

</style>