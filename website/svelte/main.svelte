<script lang="ts">
	import Test from "../ts/wasm";
	import Dialog from './dialog.svelte';
	import TODO from './todo.svelte'
	import Scale from './scale.svelte'
	import MenuDropDown from '../img/menu.svg'
	
	console.log(MenuDropDown)
	
	import {Button, Checkbox, Icon, Menu, Menuitem, Textfield} from "svelte-mui";
</script>

<style>
    h1 {
        color: aqua;
    }
</style>

<Textfield bind:value filled label="Name" message="Enter your name"/>

<h1>Hello {value}!</h1>

<Checkbox bind:checked>Checkbox</Checkbox>

<p>Checkbox is <strong>{checked ? 'checked' : 'unchecked'}</strong></p>

<Button outlined shaped color="Red" on:click={() => { checked = !checked }}>
	Inverse
</Button>


<Dialog visible={true} username={"test"} password={"test"} onclose={(username, passwd) => {console.log(username, passwd)}}/>

<div style="display: flex; gap: 50px; flex-direction: column">
	<TODO/>
	<Scale/>
	
	
	<Menu origin="top left">
		<div slot="activator">
			<Button color="primary" outlined ripple={false} style="padding-right: 4px;">
				<span>Menu</span>
				<Icon path={MenuDropDown}/>
			</Button>
		</div>
		
		<Menuitem on:click={()=>{console.log("Refresh")}}>Refresh</Menuitem>
		<Menuitem on:click={()=>{console.log("Send feedback")}}>Send feedback</Menuitem>
		<hr/>
		<Menuitem on:click={()=>{console.log("Settings")}} disabled>Settings</Menuitem>
		<Menuitem on:click={()=>{console.log("Help")}}>Help</Menuitem>
		<Menuitem on:click={()=>{console.log("Sign In")}}>Sign In</Menuitem>
	</Menu>
</div>


