<script lang="ts">
	import {onMount} from "svelte";

	import Settings from "./settings.svelte"
	import Logout from "./logout.svelte"
	import Source from "./source.svelte"
	import {page} from "$app/stores";

	let hovered = false

	function handleMouseEnter() {
		hovered = true;
		document.getElementById('navbarfocus').focus() // some element inside navbar needs focus for close to work
	}

	onMount(() => {
		document.getElementById('svelte-root').addEventListener('focusout', (e) => {
			if(document.getElementById('navbar').contains(e.target as HTMLElement)) {
				if(!(document.getElementById('navbar').contains(e.relatedTarget as HTMLElement) || document.getElementById('navbar') == e.relatedTarget as HTMLElement)) {
					hovered = false
				}
			}
		})
		handleMouseEnter()
	})

</script>

<div id="navbar" class:mdc-top-app-bar--short-collapsed={!hovered} class="mdc-top-app-bar mdc-top-app-bar--short"
	  on:mouseenter={handleMouseEnter}>
	<div class="mdc-top-app-bar__row" style="height: 70px">
		<section class="mdc-top-app-bar__section mdc-top-app-bar__section--align-start">
			<img class="icon" class:roundedIcon={!hovered} src="/favicon.ico" alt="favicon">
			<h3 style="font-size: 2em">Webgame</h3>
		</section>
		<section class:hide={!hovered} class="mdc-top-app-bar__section align-middle">
			<div class:active={$page.url.pathname === '/'}><a sveltekit:prefetch href="/">Home</a></div>
			<div class:active={$page.url.pathname === '/games'}><a sveltekit:prefetch href="/games">Games</a></div>
			<div class:active={$page.url.pathname === '/about'}><a sveltekit:prefetch href="/about">About</a></div>
			<div class:active={$page.url.pathname === '/testWasm'}><a sveltekit:prefetch href="/testWasm">testWasm</a></div>
		</section>
		<section class:hide={!hovered} class="mdc-top-app-bar__section mdc-top-app-bar__section--align-end">
			<Source/>
			<Settings/>
			<Logout/>
		</section>
	</div>
	<input id="navbarfocus" class="hidden"/>
</div>

<div class="navbarOffset" class:navbarOffsetCompressed={!hovered}></div>

<style lang="scss">
	@use "src/css/vars";

	#navbar {
		overflow: hidden;
		height: 70px;

		.hide {
			opacity: 0;
			transition: 0.3s;
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
				transition: 0.4s;
			}

			.active {
				text-decoration: underline;
			}
		}
	}

	.navbarOffset {
		padding-top: vars.$navbar-height;
		transition: 0.3s;
	}

	.navbarOffsetCompressed {
		padding-top: 0;
		animation: linear moveOut 0.3s;
	}

	@keyframes moveOut {
		0% {
			padding-top: vars.$navbar-height;
		}
		40% {
			padding-top: vars.$navbar-height * 0.9;
		}
		100% {
			padding-top: 0;
		}
	}

	.mdc-top-app-bar--short-collapsed {
		width: vars.$navbar-height;
		height: vars.$navbar-height;
	}

	.mdc-top-app-bar {
		color: vars.$on-primary;
		background: vars.$primary;
	}

	a {
		color: inherit;
		text-decoration: inherit;
		font-size: 1.2em;
	}
</style>