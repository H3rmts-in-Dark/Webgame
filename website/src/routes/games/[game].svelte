<script context="module" lang="ts">
	import {getGame} from "./game";

	export async function load({params}) {
		return {
			stuff: await getGame(params.game)
		};
	}
</script>

<script lang="ts">
	import {page} from '$app/stores';
	import type {Game} from "./game";

	let game: Game = $page.stuff as Game
	console.table(game)
</script>

<svelte:head>
	<title>Game</title>
</svelte:head>

<div style="display: flex; flex-direction: column; align-items: center">
	<h0>Game {game.name}</h0>
</div>

<h2>{game.id}, {game.limit}, {game.name}</h2>