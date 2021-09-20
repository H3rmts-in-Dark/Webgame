const path = require("path");

//const CopyPlugin = require("copy-webpack-plugin");
const WasmPackPlugin = require("@wasm-tool/wasm-pack-plugin");

const out = path.resolve(__dirname, "resources/sites");

const src = path.resolve(__dirname, "webassembly");

module.exports = {
    mode: "development", // development | production
    entry: {
        index: "./js/index.js",
    },
    output: {
        path: out,
        filename: "pack.js",
    },
    devServer: {
        contentBase: out,
    },
    plugins: [
        new WasmPackPlugin({
            crateDirectory: src,
        }),
    ],
};
