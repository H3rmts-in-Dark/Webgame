<script lang="ts">
	import Button from "@smui/button";

	let todos = [
		{done: true, text: 'finish Svelte tutorial'},
		{done: false, text: 'build an app'},
		{done: false, text: 'world domination'}
	];

	function add() {
		todos = todos.concat({done: false, text: ''});
	}

	function clear() {
		todos = todos.filter(t => !t.done);
	}

	$: remaining = todos.filter(t => !t.done).length;
</script>

<div>

	<h1>Todos</h1>

	{#each todos as todo}
		<div class:done={todo.done} class:notdone={!todo.done}>
			<input type=checkbox bind:checked={todo.done}>
			<input placeholder="What needs to be done?" bind:value={todo.text}>
		</div>
	{/each}

	<p>{remaining} remaining</p>

	<Button variant="outlined" color="primary" on:click={add} class="button-shaped-notch">
		Add new
	</Button>

	<Button variant="outlined" color="secondary" on:click={clear} class="button-shaped-round">
		Clear completed
	</Button>
</div>

<style lang="sass">
	.done
		opacity: 0.2
		transition: 0.3s
	.notdone
		opacity: 1
		transition: 0.3s

</style>