<script lang="ts">
	import {page} from '$app/stores';
	import type {Game} from "./game";
	import {getGame} from "./game";
	import {onDestroy} from "svelte";
	import Button from "@smui/button";
	import Textfield from "@smui/textfield";

	let game: Promise<Game> = getGame($page.params.game)

	let websocket: WebSocket = undefined;

	let start = undefined;

	async function ws() {
		let g = await game
		try {
			websocket = new WebSocket(`ws://localhost:6969/ws/${g.id}`);
			console.log("Connection built");
		} catch (err) {
			console.log("Connection invalid", err);
			return
		}

		websocket.onopen = function () {
			console.log("connection opened!");
		};

		websocket.onerror = function (error) {
			console.log("WebSocket Error: " + error);
		};

		websocket.onclose = function () {
			console.log("Connection lost");
		};

		websocket.onmessage = function (mess: MessageEvent) {
			let end = new Date().getTime()
			console.log(mess)
			recived = mess.data
			console.log(end - start)
		}
	}

	onDestroy(() => {
		if (websocket != undefined)
			websocket.close()
	})

	let send = "fuf"
	let recived = "--"

</script>

<svelte:head>
	<title>Game</title>
</svelte:head>

{#await game}
	<h0>Loading</h0>
{:then game }
	<div style="display: flex; flex-direction: column; align-items: center">
		<h0>Game {game.name}</h0>
	</div>

	<h2>{game.id}, {game.limit}, {game.name}</h2>

	<Button variant="outlined" color="primary" on:click={ws}>
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