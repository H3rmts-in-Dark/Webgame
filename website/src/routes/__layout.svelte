<script lang="ts">
	import Navbar from '$lib/navbar/navbar.svelte'
	import "../css/global.scss"
	import "../css/smui.css"
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