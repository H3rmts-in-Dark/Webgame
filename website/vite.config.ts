import {defineConfig} from 'vite'
import {svelte} from '@sveltejs/vite-plugin-svelte'

// https://vitejs.dev/config/

// noinspection JSUnusedGlobalSymbols
export default defineConfig({
	root: './',
	build: {
		target:['edge90','chrome90','firefox90','safari15'],
		minify: false
	},
	plugins: [
		svelte(),
	]
})
