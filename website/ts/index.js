import "../css/smui.css"

// @ts-ignore
import App from '../svelte/index.svelte';
import Test from "./wasmtest";

Test()

new App({
	target: document.getElementById('root'),
})

export default App
