<script lang="ts">
	import {page} from '$app/stores';
	import type {Game} from "./game";
	import {getGameFromServer} from "./game";
	import {onDestroy} from "svelte";
	import Button from "@smui/button";
	import Textfield from "@smui/textfield";
	import {buildWebsocket} from "../../ts/websocket.ts";
	import Title from "../../lib/Title.svelte";

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
					console.error(err)
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
	<Title title="Loading"></Title>
{:then game }
	<Title title={game.name}></Title>

	{#if initialising}
		<div id="connect">
			{#if !connected}
				<Button variant="outlined" color="primary" on:click={Websocket}>
					Connect
				</Button>
			{/if}
			<div>
				<Textfield class="shaped-outlined" variant="outlined" bind:value={send} label="Name"/>
			</div>
			{#if connected}
				<Button variant="outlined" color="primary" on:click={() => {console.time("ws");websocket.send(send)}}>
					Start
				</Button>
			{:else }
				<Button disabled variant="outlined" color="primary">
					Start
				</Button>
			{/if}
		</div>
	{:else }
		*Game*
	{/if}
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