import init from '../webassembly/pkg';

const wasm = await init();

console.log(wasm)
console.log(wasm.greet)

function Test() {
	wasm.greet();
}

export default Test