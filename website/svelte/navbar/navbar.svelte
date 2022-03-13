<script lang="ts">
	import {onMount} from "svelte";

	import Settings from "./settings.svelte"
	import Logout from "./logout.svelte"
	import Source from "./source.svelte"

	let hovered = false

	function handleMouseEnter() {
		hovered = true;
		document.getElementById('navbarfocus').focus() // some element inside navbar needs focus for close to work
	}

	onMount(() => {
		document.getElementById('svelte-root').addEventListener('focusout', (e) => {
			if (document.getElementById('navbar').contains(e.target as HTMLElement)) {
				if (!(document.getElementById('navbar').contains(e.relatedTarget as HTMLElement) || document.getElementById('navbar') == e.relatedTarget as HTMLElement)) {
					hovered = false
				}
			}
		})
	})

</script>

<div id="navbar" class:mdc-top-app-bar--short-collapsed={!hovered} class="mdc-top-app-bar mdc-top-app-bar--short"
	  on:mouseenter={handleMouseEnter}>
	<input id="navbarfocus" class="hidden"/>
	<div class="mdc-top-app-bar__row" style="height: 70px">
		<section id="mainPage" class="mdc-top-app-bar__section mdc-top-app-bar__section--align-start">
			<img class="icon" class:roundedIcon={!hovered} src="../img/favicon.ico" alt="favicon">
			<h3 style="font-size: 2em">Webgame</h3>
		</section>
		<section class:hide={!hovered} class="mdc-top-app-bar__section mdc-top-app-bar__section--align-end">
			<Source/>
			<Settings/>
			<Logout/>
		</section>
	</div>
</div>

<style lang="sass">
	#navbar
		overflow: hidden
		height: 70px

		.hide
			opacity: 0
			transition: 0.3s

			gap: 10px

		section
			gap: 10px
			padding: 6px
			transition: 0.3s

			.icon
				height: 100%
				border-radius: 8px 8px 8px 8px
				transition: 0.1s

			.roundedIcon
				border-radius: 8px 8px 24px 8px
				transition: 0.5s
</style>