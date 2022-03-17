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
	<div class="mdc-top-app-bar__row" style="height: 70px">
		<section class="mdc-top-app-bar__section mdc-top-app-bar__section--align-start">
			<img class="icon" class:roundedIcon={!hovered} src="static/favicon.ico" alt="favicon">
			<h3 style="font-size: 2em">Webgame</h3>
		</section>
		<section class:hide={!hovered} class="mdc-top-app-bar__section align-middle">
			<h2>Home</h2>
			<h2>Games</h2>
		</section>
		<section class:hide={!hovered} class="mdc-top-app-bar__section mdc-top-app-bar__section--align-end">
			<Source/>
			<Settings/>
			<Logout/>
		</section>
	</div>
	<input id="navbarfocus" class="hidden"/>
</div>

<div id="navbaroffset"></div>

<style lang="scss">
	@import "src/css/vars";

	#navbar {
		overflow: hidden;
		height: 70px;

		.hide {
			opacity: 0;
			transition: 0.3s;

			gap: 10px;
		}

		.align-middle {
			justify-content: center;
			order: 1
		}

		section {
			display: flex;
			gap: 10px;
			padding: 6px;
			transition: 0.3s;

			.icon {
				height: 100%;
				border-radius: 8px 8px 8px 8px;
				transition: 0.1s;
			}

			.roundedIcon {
				border-radius: 8px 8px 24px 8px;
				transition: 0.5s;
			}
		}
	}

	#navbaroffset {
		padding-top: $navbar-height;
	}

	.mdc-top-app-bar--short-collapsed {
		width: $navbar-height;
		height: $navbar-height;
	}

	.mdc-top-app-bar {
		color: $on-primary;
		background: $primary;
	}
</style>