<script lang="ts">
	import Slider from "@smui/slider";

	export let data: number
	export let name: string
</script>

<div>
	<Slider style="flex-grow: 1;" bind:value={data}/>
	<pre class="status">{name}{data}</pre>
</div>

<style lang="scss">
	div {
		display: flex;
		flex-direction: column;
		flex-grow: 1;

		padding: 10px;

		pre {
			margin: 0;
			padding-left: 2em;
		}
	}
</style>