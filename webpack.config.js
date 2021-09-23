const path = require("path");

const WasmPackPlugin = require("@wasm-tool/wasm-pack-plugin");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");

const out = path.resolve(__dirname, "site");
const src = path.resolve(__dirname, "webassembly");

const Modes = {
	Development: "development",
	Production: "production",
};

const mode = Modes.Development;

module.exports = {
	mode: mode, // renames everything to shit
	devtool: "eval-cheap-module-source-map",
	entry: {
		index: "./js/index.ts",
	},
	output: {
		path: out,
		filename: "pack.js",
	},
	plugins: [
		new HtmlWebpackPlugin({
			filename: "index.html",
			title: "Webgame",
			mode: mode,
			favicon: "resources/favicon.ico",
			template: "html/login.html"
		}),
		new WasmPackPlugin({
			crateDirectory: src,
			outDir: "pkg",
			outName: "WasmPack",
			forceMode: mode, // shortens wasm file significantly
		}),
		new MiniCssExtractPlugin({
			filename: "index.css",
		}),
	],
	resolve: {
		extensions: [".ts", ".js"],
	},
	experiments: {
		asyncWebAssembly: true,
	},
	module: {
		rules: [
			{
				test: /\.ts$/,
				use: "ts-loader",
			},
			{
				test: /\.css$/,
				use: [
					MiniCssExtractPlugin.loader,
					{
						loader: "css-loader",
						options: {
							sourceMap: true,
						},
					},
				],
			}
		],
	},
};
