<script lang="ts">
	import {count} from '../ts/stores';

	export var init = '0'
	count.init(init)
</script>

<div>
	<h1>The count is {$count}</h1>

	<button on:click={count.increment}>+</button>
	<button on:click={count.decrement}>-</button>
	<button on:click={count.reset}>reset</button>
</div>
