import {defineConfig} from 'vite'
import {svelte} from '@sveltejs/vite-plugin-svelte'
// import wasmPack from 'vite-plugin-wasm-pack'

// https://vitejs.dev/config/
export default defineConfig({
	root: './',
	build: {
		target:['edge90','chrome90','firefox90','safari15'],
		minify: false
	},
	plugins: [
		svelte(),
		// wasmPack(['../webgame'])
	]
})
