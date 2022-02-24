// @ts-ignore
import {writable} from 'svelte/store';

function createCount() {
	const {subscribe, set, update} = writable(0);

	return {
		subscribe,
		increment: () => update(n => n + 1),
		decrement: () => update(n => n - 1),
		reset: () => set(0),
		init: (initial) => set(initial)
	};
}

export const count = createCount();