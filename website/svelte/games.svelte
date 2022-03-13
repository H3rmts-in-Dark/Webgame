<script lang="ts">
	import {Game} from "../ts/game"
	import Button from "@smui/button";

	let games: Game[] = [
		new Game(123141241412, "test"),
		new Game(41323256243634, "test2"),
		new Game(7803563453463, "test3"),
	]

	let join = (game: Game) => {
		console.log(`joining game`)
		console.table(game)
	}
</script>

<div style="display: flex; flex-direction: column; align-items: center">
	<h1>Games</h1>
	<div id="games">
		{#each games as game}
			<div class="game">
				<h2>{game.name}</h2>
				<Button variant="raised" color="secondary" on:click={() => {join(game)}} class="button-shaped-notch">
					Join
				</Button>
			</div>
		{/each}
	</div>
</div>

<style lang="sass">
	@import "../css/glob"

	.game
		background: $primary
		padding: 10px
		margin: 10px
		color: $on-primary
		border: solid $on-primary 6px

		:border // TODO fix
			style: solid
			width: 6px
			color: $on-primary

		display: flex
		flex-direction: column
		align-items: center

	#games
		margin: 15px
		display: grid
		grid-gap: 20px

		grid-template-columns: repeat(auto-fill, minmax(380px, 1fr))


</style>