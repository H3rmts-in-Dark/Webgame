<script lang="ts">
	import {onMount} from "svelte";

	export var link

	onMount(() => {
		document.getElementById('link').click()
	})

	// this is the simplest way to route in js
</script>

<a id="link" sveltekit:prefetch href="{link}">Link</a>

<style>
	#link {
		display: none;
	}
</style>