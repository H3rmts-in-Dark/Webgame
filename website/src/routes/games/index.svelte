<script lang="ts">
	import type {CreateGame, Game} from "./game";
	import {check, createGameOnServer, getGamesFromServer} from "./game"
	import Button from "@smui/button";

	import {mdiLoading} from "@mdi/js";
	import SvgIcon from "$lib/SvgIcon.svelte";
	import Open from "$lib/open.svelte"
	import FormField from '@smui/form-field';
	import Textfield from "@smui/textfield";
	import Checkbox from '@smui/checkbox';
	import Slider from "$lib/Slider.svelte";
	import {sleep} from "../../ts/util";
	import {onMount} from "svelte";
	import Title from "../../lib/Title.svelte";
	import {CheckCodes} from "../../ts/dto/checkCodes";

	type GameDisplay = {
		game: Game
		buttonDisplay: string
		code: string
	}

	let creatingGame = false
	let newGame: CreateGame = {code: "", limit: 4, name: "new Game", visible: false}
	let log: number | string | undefined = undefined

	// Array if loaded, null if loading, string if error
	let games: Array<GameDisplay> | null | string = null

	async function loadGames() {
		creatingGame = false
		games = null
		games = await getGamesFromServer()
				.then(async (games: Array<Game>) => {
					await sleep(50) // show loading text to let user know something happened
					return games.map((game: Game) => {
						return {game: game, buttonDisplay: "Open", code: ''}
					})
				})
				.catch((err: Error) => {
					return err.message
				})
	}

	async function joinGame(gameDisplay: GameDisplay) {
		console.log(`joining game`)


		gameDisplay.buttonDisplay = "Load"

		games = games // forces rerender

		let ok = await check(gameDisplay.game.id, gameDisplay.code)

		if(ok == CheckCodes.Ok)
			gameDisplay.buttonDisplay = "__OpenLink"
		else {
			switch(ok) {
				case CheckCodes.AlreadyPlaying:
					gameDisplay.buttonDisplay = "Already Playing"
					break
				case CheckCodes.CodeWrong:
					gameDisplay.buttonDisplay = "Wrong Code"
					break
				case CheckCodes.MaxPlayersReached:
					gameDisplay.buttonDisplay = "Max Players Reached"
					break
			}
		}
		games = games // forces rerender
	}

	async function createGame() {
		log = await createGameOnServer(newGame).then((game: Game) => {
			console.log("created game:", game)
			return game.id
		}).catch((err: Error) => {
			return "error: " + err.message
		})
	}

	async function toggleCreateGame() {
		creatingGame = !creatingGame
		if(!creatingGame) {
			await loadGames()
		}
	}

	onMount(() => {
		loadGames()
	})
</script>

<svelte:head>
	<title>Games</title>
</svelte:head>

<Title title="Games"></Title>

<div id="buttons_bar">
	<div>
		<Button variant="outlined" color="primary" on:click={toggleCreateGame}>
			{creatingGame ? "List" : "Create"}
		</Button>
	</div>
	<div>
		<Button variant="outlined" color="primary">
			Hidden
		</Button>
		<Button variant="outlined" color="primary" on:click={loadGames}>
			Scan
		</Button>
	</div>
</div>
{#if creatingGame}
	<div id="createGame">
		<h1 class="title">Create new Game</h1>
		<div class="field">
			<Textfield style="width: 45%;" class="shaped-outlined" variant="outlined" bind:value={newGame.name}
						  label="Name"/>
		</div>
		<div class="field">
			<Slider bind:data={newGame.limit} name="Max Players: "></Slider>
		</div>
		<div class="field">
			<FormField>
				<Checkbox bind:checked={newGame.visible} touch/>
				<h3 slot="label">Visible</h3>
			</FormField>
			<Textfield style="width: 35%;" class="shaped-outlined" variant="outlined" bind:value={newGame.code}
						  label="Code"/>
		</div>
		<div class="field">
			<Button variant="raised" color="secondary" on:click={createGame}>
				Create
			</Button>
			{#if log}
				{#if log.startsWith("error: ") }
					<h3>{log}</h3>
				{:else }
					<Open link={`games/${log}`}/>
				{/if}
			{/if}
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
						{game.game.players} / {game.game.limit}
					</h3>
					{#if game.game.code}
						<Textfield variant="outlined" bind:value={game.code} label="Code"/>
					{/if}
					{#if game.buttonDisplay === "Load"}
						<Button variant="raised" color="secondary">
							<SvgIcon cls="rotate" svg={mdiLoading}/>
						</Button>
					{:else if game.buttonDisplay === "__OpenLink"}
						<Open link={`games/${game.game.id}`}/>
					{:else}
						<Button variant="raised" color="secondary" on:click={() => {joinGame(game)}}>
							{game.buttonDisplay}
						</Button>
					{/if}
				</div>
			{/each}
		{/if}
	</div>
{/if}

<style lang="scss">
	@use "src/css/vars";

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
		width: 50%;
		background: vars.$primary-surface;
		color: vars.$on-primary-surface;
		margin-left: calc(25% - 10px - 6px);
		margin-right: calc(25% - 10px - 6px);

		border: {
			style: solid;
			width: 6px;
			color: vars.$on-primary-surface;
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
		background: vars.$primary-surface;
		padding: 10px;
		margin: 10px;
		color: vars.$on-primary-surface;

		border: {
			style: solid;
			width: 0.4em;
			color: vars.$on-primary-surface;
		}

		display: flex;
		gap: 1rem;
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