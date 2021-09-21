const path = require("path");

//const CopyPlugin = require("copy-webpack-plugin");
const WasmPackPlugin = require("@wasm-tool/wasm-pack-plugin");

const out = path.resolve(__dirname, "resources/sites");

const src = path.resolve(__dirname, "webassembly");

const mode = "development" // development | production

module.exports = {
	mode: mode, // renames everything to shit
	entry: {
		index: "./js/index.ts",
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
			outDir: "pkg",
			outName: "WasmPack",
			forceMode: mode // shortens wasm file significantly
		}),
	],
	resolve: {
		extensions: ['.tsx', '.ts', '.js'],
	},
	module: {
		rules: [
			{
				test: /\.ts$/,
				use: 'ts-loader'
			}, {
				test: /\.css$/,
				use: 'css-loader'
			}
		],
	}
};
