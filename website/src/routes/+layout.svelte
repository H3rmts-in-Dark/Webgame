<script>
	import Navbar from '$lib/navbar/navbar.svelte'

	import "../css/global.scss"
	import "../css/vars.scss"
	import "../css/smui.css"
</script>

<div class="app" id="svelte-root">
	<header>
		<Navbar/>
	</header>

	<main>
		<slot/>
	</main>

	<footer>
	</footer>
</div>

<style>
</style>
