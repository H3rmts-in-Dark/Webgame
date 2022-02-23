<script lang="ts">
	import Test from "../ts/wasm";
	import TODO from './todo.svelte'
	import Scale from './scale.svelte'
	import Count from "./counter.svelte";
	import MenuDropDown from '../img/menu.svg'

	import Button from '@smui/button'
	import Textfield from '@smui/textfield';
	import IconButton, {Icon} from '@smui/icon-button';
	import Checkbox from '@smui/checkbox';
	import {Svg} from '@smui/common/elements';
	import CharacterCounter from "@smui/textfield/character-counter";


	export let name = "name missing";

	let value: string = name;

	let checked: boolean = true;

	console.log(MenuDropDown, MenuDropDown.length)

	let svg = MenuDropDown

	Test()
</script>

<div>
	<h1>Hello {value}!</h1>
	<Textfield bind:value class="shaped-outlined" label="Name" variant="outlined" prefix="$" input$maxlength={18}
				  style="width: 100%;"
				  helperLine$style="width: 100%;">
		<CharacterCounter slot="helper">0 / 18</CharacterCounter>
	</Textfield>
</div>


<Checkbox bind:checked indeterminate={checked === null} input$required/>

<p>Checkbox is <strong>{checked ? 'checked' : 'unchecked'}</strong></p>

<Button variant="raised" color="secondary" on:click={() => { checked = !checked }}>
	Inverse
</Button>

<Button variant="raised" color="primary" on:click={() => { checked = !checked }}>
	Inverse2
</Button>

<div style="display: flex; gap: 50px; flex-direction: column">
	<TODO/>
	<Scale/>


	<Count init='2'/>
	<br>
</div>


<div style="display: flex; align-items: center;">
	<IconButton>
		<Icon component={Svg} bind:svg viewBox="0 0 24 24">
		</Icon>
	</IconButton>
</div>
