<script lang="ts">
	import Title from "../lib/Title.svelte";
	import {run} from "../ts/wasm";
</script>

<svelte:head>
	<title>WASM Test</title>
</svelte:head>

<Title title="WASM TEST"></Title>

<button on:click={run}>
	START
</button>

<div id="1234">
	box
</div>

<style lang="scss">

</style>