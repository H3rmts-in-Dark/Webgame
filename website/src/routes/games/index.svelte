<script lang="ts">
	import {loadGames, create, hidden} from "./game"
	import type {Game} from "./game";
	import Button from "@smui/button";

	import {mdiLoading} from "@mdi/js";
	import SvgIcon from "$lib/SvgIcon.svelte";

	type GameDisplay = {
		game: Game
		buttonDisplay: string
	}

	let games: Array<GameDisplay> | null | string
	let load = async () => {
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

	let join = (gameDisplay: GameDisplay) => {
		console.log(`joining game`)
		gameDisplay.buttonDisplay = "Load"
		console.table(gameDisplay)
		// forces update
		games = games
	}
</script>

<svelte:head>
	<title>Games</title>
</svelte:head>

<div style="display: flex; flex-direction: column; align-items: center">
	<h0>Games</h0>
	<div id="buttons_bar">
		<div>
			<Button variant="outlined" color="primary" on:click={create}>
				Create
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
	<div id="games">
		{#if games == null}
			<h2>Loading</h2>
		{:else if typeof games == "string"}
			<h2>Error {games}</h2>
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

	#games {
		margin: 15px;
		display: grid;
		grid-gap: 20px;
		width: calc(100% - 30px);

		grid-template-columns: repeat(auto-fill, minmax(300px, 3fr));
	}
</style>