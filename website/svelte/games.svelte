<script lang="ts">
	import {loadGames} from "../ts/game"
	import Button from "@smui/button";

	let games = loadGames()

	let join = (game) => {
		console.log(`joining game`)
		console.table(game)
	}
</script>

<div style="display: flex; flex-direction: column; align-items: center">
	<h1>Games</h1>
	<div id="games">
		{#await games}
			<h2>Loading</h2>
		{:then games}
			{#each games as game}
				<div class="game">
					<h2>{game.name}</h2>
					<Button variant="raised" color="secondary" on:click={() => {join(game)}} class="button-shaped-notch">
						Join
					</Button>
				</div>
			{/each}
		{:catch error}
			<h2>Error {error}</h2>
		{/await}
	</div>
</div>

<style lang="scss">
	@import "../css/vars";

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