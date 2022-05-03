<script context="module">
	export function load({error, status}) {
		return {
			props: {
				title: status + ':' + ' ' + error.message
			}
		};
	}
</script>

<script>
	export let title;
</script>

<svelte:head>
	<title>{title}</title>
</svelte:head>

<h1>{title}</h1>


<style lang="scss">
	@use "src/css/vars";

	h1 {
		padding-top: vars.$navbar-height;
		text-align: center;
	}

</style>

