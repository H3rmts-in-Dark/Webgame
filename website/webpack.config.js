const path = require("path");

const WasmPackPlugin = require("@wasm-tool/wasm-pack-plugin");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");

const out = path.resolve(__dirname, "./site");
const wasmSrc = path.resolve(__dirname, "webassembly");
const wasmOut = path.resolve(__dirname, "pkg");

const Modes = {
	Development: "development",
	Production: "production",
};

const mode = Modes.Development;

module.exports = {
	mode: mode, // renames everything to shit
	devtool: "eval-cheap-module-source-map",
	entry: {
		index: "./ts/index.ts",
	},
	output: {
		path: out,
		filename: "pack.js",
	},
	plugins: [
		new HtmlWebpackPlugin({
			filename: "index.html",
			template: "html/index.html",
			favicon: "resources/favicon.ico",
			mode: mode,
			inject: "body",
		}),
		new WasmPackPlugin({
			crateDirectory: wasmSrc,
			outName: "wasm",
			outDir: wasmOut,
			forceMode: mode, // shortens wasm file significantly
		}),
		new MiniCssExtractPlugin({
			filename: "index.css",
		})
	],
	experiments: {
		asyncWebAssembly: true,
	},
	resolve: {
		extensions: ['.js', '.wasm', '.ts', '.css', '.svg']
	},
	module: {
		rules: [
			{
				test: /\.svg$/,
				use: {
					loader: 'svg-url-loader',
					options: {
						encoding: "utf-8",
					},
				},
			},
			{
				test: /\.svelte$/,
				use: {
					loader: 'svelte-loader',
					options: {
						emitCss: mode === Modes.Production,
						sourceMap: true,
						preprocess: require('svelte-preprocess')({

						})
					}
				},
			},
			{
				test: /\.ts$/i,
				use: {
					loader: "ts-loader",
					options: {
						transpileOnly: mode === Modes.Development,   // improves speed
					}
				}
			},
			{
				test: /\.(css|sass)$/,
				use: [
					mode === Modes.Development ? {
						loader: "style-loader",   // loads css with style tags
						options: {
							injectType: "autoStyleTag",
						}
					} : {
						loader: MiniCssExtractPlugin.loader,
						options: {}
					},
					{
						loader: "css-loader",
						options: {
							sourceMap: mode === Modes.Development,   // generates the ./css in devtools (origin files)
						},
					},
					{
						loader: "sass-loader",
						options: {
							sourceMap: mode === Modes.Development,   // generates the ./css in devtools (origin files)
						}
					}
				],
			},
			{
				test: /\.html$/i,
				use: {
					loader: "html-loader",
					options: {
						minimize: mode === Modes.Production,
					}
				}
			},
		],
	},
};
