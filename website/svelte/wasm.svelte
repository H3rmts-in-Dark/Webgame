<script lang="ts">
	import {onMount} from 'svelte';
	import Test from "../ts/wasmtest";

	onMount(Test)
</script>

<div>
	<button id="play-pause"></button>
	<br>
	<label>
		Width:
		<input id='width-input' type="number" min="1" value="40">
	</label>
	<label>
		Height:
		<input id='height-input' type="number" min="1" value="40">
	</label>
	<br>
	<div id="fps"></div>
	<br>
	<canvas id="canvas"></canvas>
</div>


<!--<style lang="sass">

</style>-->
