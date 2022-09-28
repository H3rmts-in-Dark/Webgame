<script lang="ts">
	import {page} from '$app/stores';
	import type {Game} from "./game";
	import {getGameFromServer} from "./game";
	import Title from "../../lib/Title.svelte";

	let game: Promise<Game> = getGameFromServer($page.params.game)
</script>

<svelte:head>
	<title>Game</title>
</svelte:head>

{#await game}
	<Title title="Loading"></Title>
{:then game }
	<Title title={game.name}></Title>
	<div id="game">

	</div>
{/await}

<style lang="scss">
	#connect {
		display: flex;
		flex-direction: column;
		width: auto;
		gap: 3em;
		align-items: center;

		margin-top: 2em;
	}
</style>