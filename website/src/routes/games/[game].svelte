<script lang="ts">
	import {page} from '$app/stores';
	import type {Game} from "./game";
	import {getGame} from "./game";
	import {onDestroy} from "svelte";
	import Button from "@smui/button";
	import Textfield from "@smui/textfield";
	import {buildWebsocket} from "./game.ts";

	let game: Promise<Game> = getGame($page.params.game)

	let websocket: WebSocket = undefined;

	async function Websocket() {
		websocket = buildWebsocket(await game)
		websocket.onmessage = function(mess: MessageEvent) {
			let end = new Date().getTime()
			console.debug(mess)
			recived = mess.data
			console.debug(end - start, "ms")
		}
	}

	let start = undefined;

	onDestroy(() => {
		if(websocket != undefined)
			websocket.close()
	})

	let send = "fuf"
	let recived = "--"

</script>

<svelte:head>
	<title>Game {game.name}</title>
</svelte:head>

{#await game}
	<h0>Loading</h0>
{:then game }
	<div style="display: flex; flex-direction: column; align-items: center">
		<h0>Game {game.name}</h0>
	</div>

	<h2>{game.id}, {game.limit}, {game.name}</h2>

	<Button variant="outlined" color="primary" on:click={Websocket}>
		Connect
	</Button>
	<Textfield class="shaped-outlined" variant="outlined" bind:value={send} label="SEnd"/>
	<Button variant="outlined" color="primary" on:click={() => {start = new Date().getTime();websocket.send(send)}}>
		Send
	</Button>
	<h2>
		{recived}
	</h2>
{/await}