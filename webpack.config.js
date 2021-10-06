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
		index: "./ts/index.ts",
	},
	output: {
		path: out,
		filename: "pack.js",
	},
	plugins: [
		new HtmlWebpackPlugin({
			filename: "index.html",
			title: "Webgame",
			template: "html/index.ejs",
			favicon: "resources/favicon.ico",
			mode: mode,
			inject: "body",
		}),
		new WasmPackPlugin({
			crateDirectory: src,
			outDir: "pkg",
			outName: "WasmPack",
			forceMode: mode, // shortens wasm file significantly
		}),
		new MiniCssExtractPlugin({
			filename: "index.css",
		})
	],
	resolve: {
		alias: {  // shorten imports from ../html/test.html to HTML/test.html
			// IDE gets confused with functions in TS imports, but can resolve the file so only css and html
			CSS: path.resolve(__dirname, "css"),
			HTML: path.resolve(__dirname, "html"),
//			TS: path.resolve(__dirname, "ts")
		}
	},
	experiments: {
		asyncWebAssembly: true,
	},
	module: {
		rules: [
			{
				test: /\.ts$/i,
				use: [
					{
						loader: "ts-loader",
						options: {
							transpileOnly: mode === Modes.Development,   // improves speed
						}
					}
				]
			},
			{
				test: /\.css$/i,
				use: [
					// normal: mode === Modes.Development ?
					// but weird things happen, because some JS functions get
					// executed before css is applied by style-loader so
					// MiniCssExtractPlugin is used always
					mode === Modes.Development ?
						 {
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
				],
			},
			{
				test: /\.html$/i,
				use: [
					{
						loader: "html-loader",
						options: {
							minimize: mode === Modes.Production,
						}
					}
				]
			},
		],
	},
};
