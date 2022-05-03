<script lang="ts">
	import Navbar from '$lib/navbar/navbar.svelte'
	import "../css/global.scss"
	import "../css/smui.css"

	// https://github.com/vitejs/vite/discussions/3143
	// clears console on hmr
	if(import.meta.hot) {
		import.meta.hot.on(
				"vite:beforeUpdate",
				() => console.clear()
		);
	}
</script>

<div id="svelte-root">
	<header>
		<Navbar/>
	</header>

	<main>
		<slot/>
	</main>

	<footer>
	</footer>
</div>