<script lang="ts">
	import {page} from '$app/stores';
	import type {Game} from "./game";
	import {getGameFromServer} from "./game";
	import {onDestroy} from "svelte";
	import Button from "@smui/button";
	import Textfield from "@smui/textfield";
	import {buildWebsocket} from "../../ts/websocket.ts";

	let game: Promise<Game> = getGameFromServer($page.params.game)

	let websocket: WebSocket = undefined;
	let connected: Boolean = false

	async function Websocket() {
		websocket = buildWebsocket(await game,
				() => {
					console.debug("opened");
					connected = true
				}, mess => {
					console.debug("message");
					if(mess == "Start")
						initialising = false
				}, err => {
					console.debug("error");
				}, () => {
					console.debug("lost");
					connected = false
				}
		)
	}

	onDestroy(() => {
		if(websocket != undefined)
			websocket.close()
	})

	let send = "fuf"

	let initialising = true

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

	{#if initialising}
		<h2>{game.id}, {game.limit}, {game.name}</h2>

		<Button variant="outlined" color="primary" on:click={Websocket}>
			Connect
		</Button>
		<Textfield class="shaped-outlined" variant="outlined" bind:value={send} label="Name"/>
		{#if connected}
			<Button variant="outlined" color="primary" on:click={() => {console.time("ws");websocket.send(send)}}>
				Start
			</Button>
		{:else }
			<Button disabled variant="outlined" color="primary">
				Start
			</Button>
		{/if}
	{/if}
{/await}