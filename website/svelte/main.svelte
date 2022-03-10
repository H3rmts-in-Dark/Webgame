<script lang="ts">
	import TODO from './todo.svelte'
	import Scale from './scale.svelte'
	import Count from "./counter.svelte";
	import MenuDropDown from '../img/menu.svg'
	import Wasm from "./wasm.svelte";
	import SvgIcon from "./SvgIcon.svelte";

	import Button from '@smui/button'
	import Textfield from '@smui/textfield';
	import IconButton, {Icon} from '@smui/icon-button';
	import Checkbox from '@smui/checkbox';
	import CharacterCounter from "@smui/textfield/character-counter";

	import {mdiTrainCarPassengerDoorOpen, mdiWrench} from "@mdi/js";
	import {Svg} from "@smui/common/elements";


	export let name = "name missing";

	let value: string = name;

	let checked: boolean = true;

	console.log(MenuDropDown, MenuDropDown.length)
</script>

<div>
	<h1>Hello {value}!</h1>
	<Textfield bind:value class="shaped-outlined" label="Name" variant="outlined" prefix=">" input$maxlength={18} style="width: 100%;"
				  helperLine$style="width: 100%;">
		<CharacterCounter slot="helper">0 / 18</CharacterCounter>
	</Textfield>
</div>


<Checkbox bind:checked indeterminate={checked === null} color="primary" input$required/>

<p>Checkbox is <strong>{checked ? 'checked' : 'unchecked'}</strong></p>

<Button variant="raised" color="primary" on:click={() => { checked = !checked }}>
	Inverse Prim
</Button>

<SvgIcon svg={mdiTrainCarPassengerDoorOpen} size="button"/>

<IconButton class="material-icons" size="button">
	<Icon component={Svg} viewBox="0 0 24 24">
		<path fill="currentColor" d={mdiTrainCarPassengerDoorOpen}/>
	</Icon>
</IconButton>

<Button variant="raised" color="secondary" on:click={() => { checked = !checked }}>
	Inverse2 Sec
</Button>

<div style="display: flex; gap: 50px; flex-direction: column">
	<TODO/>
	<Scale/>
	<Count init='2'/>
	<br>
</div>


<div style="display: flex; align-items: center;">
	<IconButton class="material-icons">build</IconButton>
	<IconButton>
		<Icon component={Svg} viewBox="0 0 24 24">
			<path fill="currentColor" d={mdiWrench}/>
		</Icon>
	</IconButton>
</div>

<Wasm/>

<style global lang="sass" src="../css/global.sass"></style>