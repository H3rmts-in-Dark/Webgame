// import "CSS/main.sass";
import "../css/smui.css"
import "../node_modules/svelte-material-ui/bare.css"

// @ts-ignore
import App from '../svelte/main.svelte';

new App({
	target: document.getElementById('root'),
})