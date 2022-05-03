<script lang="ts">
	import IconButton from "@smui/icon-button";
	import SvgIcon from "$lib/SvgIcon.svelte";

	export var svg = ''
	export var size = 'button'
	export var cls = ''

	export var disabled: string = "False"

	export var click = () => {
	}
</script>

{#if disabled === "True"}
	<IconButton class="material-icons {cls}" size={size} color="secondary" on:click={click} disabled>
		<SvgIcon svg="{svg}"/>
	</IconButton>
{:else}
	<IconButton class="material-icons {cls}" size={size} color="secondary" on:click={click}>
		<SvgIcon svg="{svg}"/>
	</IconButton>
{/if}