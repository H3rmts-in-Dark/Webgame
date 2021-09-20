async function main() {
    const wasmload = import("../webassembly/pkg");
    const wasm = await wasmload;

    console.log("python -m http.server");
    console.log(wasm);

    wasm.greet()
}

main();
