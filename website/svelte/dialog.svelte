<!-- https://svelte-mui.vercel.app/ -->

<script lang="ts">
	export let visible = false;
	export let username = '';
	export let password = '';
	
	export let onclose = () => {
	}
	
	$: valid = username !== '' && password !== ''
	
	import {Dialog, Textfield, Button} from 'svelte-mui';
</script>

<style>
    .footer {
        text-align: center;
        margin-bottom: 1rem;
        font-size: 13px;
    }

    .footer a {
        color: #f50057;
        padding-left: 1rem;
    }
</style>

<Dialog width="290" bind:visible
        on:close={() => {onclose(username, password)}}>
	<div slot="title"> Welcome!</div>
	
	<Textfield
			  name="username"
			  autocomplete="off"
			  required
			  bind:value={username}
			  label="username"
			  message="Your account name"
	/>
	<Textfield
			  type="password"
			  name="password"
			  autocomplete="off"
			  required
			  bind:value={password}
			  label="password"
			  message="Your password"
	/>
	
	<div slot="actions" class="actions center">
		<Button color="secondary">Create New</Button>
		<Button color="primary" disabled="{!valid}" on:click={() => {visible = false; onclose(username, password)}}>Sign In</Button>
	</div>
	
	<div slot="footer" class="footer">
		Forgot Password
		<a class="disabled">Reset Password</a>
	</div>
</Dialog>