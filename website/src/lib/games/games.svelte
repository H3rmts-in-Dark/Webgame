<script lang="ts">
	import {loadGames} from "./game"
	import type {Game} from "./game";
	import Button from "@smui/button";

	import {mdiLoading} from "@mdi/js";
	import SvgIcon from "../SvgIcon.svelte";

	type GameDisplay = {
		game: Game
		buttonDisplay: string
	}

	let games: Array<GameDisplay> | null | string
	let load = async () => {
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

	let join = (gameDisplay: GameDisplay) => {
		console.log(`joining game`)
		gameDisplay.buttonDisplay = "Load"
		console.table(gameDisplay)
		console.table(gameDisplay)
		// forces update
		games = games
	}
</script>

<div style="display: flex; flex-direction: column; align-items: center">
	<h1>Games</h1>
	<div id="scan">
		<Button variant="raised" color="primary" on:click={load}>
			Scan
		</Button>
	</div>
	<div id="games">
		{#if games == null}
			<h2>Loading</h2>
		{:else if typeof games == "string"}
			<h2>Error {games}</h2>
		{:else}
			{#each games as game}
				<div class="game">
					<h2>{game.game.name}</h2>
					{#if game.buttonDisplay === "Load"}
						<Button variant="raised" color="secondary">
							<SvgIcon cls="rotate" svg={mdiLoading}/>
						</Button>
					{:else}
						<Button variant="raised" color="secondary" on:click={() => {join(game)}}>
							{game.buttonDisplay}
						</Button>
					{/if}
				</div>
			{/each}
		{/if}
	</div>
</div>

<style lang="scss">
	@import "src/css/vars";

	#scan {
		position: absolute;
		padding-top: 20px;
		right: 20px;
	}

	.game {
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
	}

	#games {
		margin: 15px;
		display: grid;
		grid-gap: 20px;
		width: calc(100% - 30px);

		grid-template-columns: repeat(auto-fill, minmax(300px, 3fr));
	}
</style>