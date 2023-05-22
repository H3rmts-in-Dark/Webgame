<script lang="ts">
	export let title
</script>

<div id="title">
	<h1>{title}</h1>
</div>

<style lang="scss">
	#title {
		display: flex;
		flex-direction: column;
		align-items: center
	}

	h1 {
		font-size: 4em;
		margin-block-start: 0.5em;
		margin-block-end: 0.5em;
		font-weight: bold;
	}
</style>