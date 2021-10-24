// @ts-ignore
import * as wasm from "../pkg/wasm.js";

// @ts-ignore
import {memory} from "../pkg/wasm_bg.wasm";

// @ts-ignore
import html from 'HTML/wasmtest.html'

import "CSS/wasmtest.sass";


let wasmTestElement: HTMLElement


const CELL_SIZE = 5; // px
const GRID_COLOR = "#CCCCCC";
const DEAD_COLOR = "#FFFFFF";
const ALIVE_COLOR = "#000000";

function createWasmDiv() {
	wasmTestElement = document.createElement('div')
	document.body.appendChild(wasmTestElement) // must happen before outerHTML is set because it requires a parent node
	wasmTestElement.outerHTML = html
}

let paused: boolean = false
let animationId: number = null;

function wasmTest() {
	createWasmDiv()
	
	const playPauseButton = document.getElementById("play-pause");
	
	const play = () => {
		playPauseButton.textContent = "⏸";
		animationId = requestAnimationFrame(renderLoop)
		paused = false
	};
	
	const pause = () => {
		playPauseButton.textContent = "▶";
		cancelAnimationFrame(animationId);  // removes the registered animation request with its ID
		paused = true
	};
	
	playPauseButton.addEventListener("click", () => {
		if (paused) {
			play();
		} else {
			pause();
		}
	});
	
	const canvas: HTMLCanvasElement = document.getElementById("canvas") as HTMLCanvasElement;
	const ctx = canvas.getContext('2d');
	
	const fps = new FPS(document.getElementById("fps"))
	
	canvas.addEventListener("click", event => {
		console.log("canvas clicked")
		const boundingRect = canvas.getBoundingClientRect();
		
		const scaleX = canvas.width / boundingRect.width;
		const scaleY = canvas.height / boundingRect.height;
		
		const canvasLeft = (event.clientX - boundingRect.left) * scaleX;
		const canvasTop = (event.clientY - boundingRect.top) * scaleY;
		
		const row = Math.min(Math.floor(canvasTop / (CELL_SIZE + 1)), height - 1);
		const col = Math.min(Math.floor(canvasLeft / (CELL_SIZE + 1)), width - 1);
		
		universe.toggle_cell(row, col);
		
		console.log(scaleX, scaleY, canvasLeft, canvasTop, row, col)
		
		drawGrid(ctx, height, width);
		drawCells(ctx, universe, height, width);
	});
	
	const renderLoop = () => {
		fps.render();
		universe.tick();
		
		drawGrid(ctx, height, width);
		drawCells(ctx, universe, height, width);
		
		animationId = requestAnimationFrame(renderLoop);  // registers animation request and stores its ID
	};
	
	let widthButton = document.getElementById('width-input') as HTMLInputElement
	widthButton.oninput = () => {
		console.log(`width change ${widthButton.value}`)
		cancelAnimationFrame(animationId);
		
		initWASM(parseInt(widthButton.value), parseInt(heightButton.value))
		
		drawGrid(ctx, height, width);
		drawCells(ctx, universe, height, width);
		if (!paused) {
			animationId = requestAnimationFrame(renderLoop)
		}
	}
	let heightButton = document.getElementById('height-input') as HTMLInputElement
	heightButton.oninput = () => {
		console.log(`height change ${heightButton.value}`)
		cancelAnimationFrame(animationId);
		
		initWASM(parseInt(widthButton.value), parseInt(heightButton.value))
		
		drawGrid(ctx, height, width);
		drawCells(ctx, universe, height, width);
		if (!paused) {
			animationId = requestAnimationFrame(renderLoop)
		}
	}
	
	initWASM(parseInt(widthButton.value), parseInt(heightButton.value))
	
	fps.render();
	drawGrid(ctx, height, width);
	drawCells(ctx, universe, height, width);
	pause()
}

let universe: wasm.Universe
let width: number
let height: number

function initWASM(_width: number, _height: number) {
	console.log(`new init ${_width} ${_height}`)
	
	// Construct the universe, and get its width and height.
	universe = wasm.Universe.new(_width, _height);
	width = universe.width();
	height = universe.height();
	
	const canvas: HTMLCanvasElement = document.getElementById("canvas") as HTMLCanvasElement;
	
	// Give the canvas room for all of our cells and a 1px border
	// around each of them.
	canvas.height = (CELL_SIZE + 1) * height + 1;
	canvas.width = (CELL_SIZE + 1) * width + 1;
}


const drawGrid = (ctx: CanvasRenderingContext2D, height: number, width: number) => {
	ctx.beginPath();
	ctx.strokeStyle = GRID_COLOR;
	
	// Vertical lines.
	for (let i = 0; i <= width; i++) {
		ctx.moveTo(i * (CELL_SIZE + 1) + 1, 0);
		ctx.lineTo(i * (CELL_SIZE + 1) + 1, (CELL_SIZE + 1) * height + 1);
	}
	
	// Horizontal lines.
	for (let j = 0; j <= height; j++) {
		ctx.moveTo(0, j * (CELL_SIZE + 1) + 1);
		ctx.lineTo((CELL_SIZE + 1) * width + 1, j * (CELL_SIZE + 1) + 1);
	}
	
	ctx.stroke();
};

function getIndex(row, width, column) {
	return row * width + column;
}

function drawCells(ctx: CanvasRenderingContext2D, universe: wasm.Universe, height: number, width: number) {
	const cellsPtr = universe.cells();
	const cells = new Uint8Array(memory.buffer, cellsPtr, width * height);
	
	ctx.beginPath();
	
	for (let row = 0; row < height; row++) {
		for (let col = 0; col < width; col++) {
			const idx = getIndex(row, width, col);
			
			ctx.fillStyle = cells[idx] === wasm.Cell.Dead
				 ? DEAD_COLOR
				 : ALIVE_COLOR;
			
			ctx.fillRect(
				 col * (CELL_SIZE + 1) + 1,
				 row * (CELL_SIZE + 1) + 1,
				 CELL_SIZE,
				 CELL_SIZE
			);
		}
	}
	
	ctx.stroke();
}

class FPS {
	private fps: HTMLElement;
	private readonly frames: number[];
	private lastFrameTimeStamp: DOMHighResTimeStamp;
	
	constructor(element: HTMLElement) {
		this.fps = element;
		this.frames = [];
		this.lastFrameTimeStamp = performance.now();
	}
	
	render() {
		// Convert the delta time since the last frame render into a measure
		// of frames per second.
		const now = performance.now();
		const delta = now - this.lastFrameTimeStamp;
		this.lastFrameTimeStamp = now;
		const fps = 1 / delta * 1000;
		
		// Save only the latest 100 timings.
		this.frames.push(fps);
		if (this.frames.length > 100) {
			this.frames.shift();
		}
		
		// Find the max, min, and mean of our 100 latest timings.
		let min = -1;
		let max = -1;
		let sum = 0;
		for (let i = 0; i < this.frames.length; i++) {
			sum += this.frames[i];
			min = Math.min(this.frames[i], min);
			max = Math.max(this.frames[i], max);
		}
		let mean = sum / this.frames.length;
		
		// Render the statistics.
		this.fps.innerText = `Frames per Second: \n latest = ${Math.round(fps)} \n avg of last 100 = ${Math.round(mean)} \n min of last 100 = ${Math.round(min)} \n max of last 100 = ${Math.round(max)}`
	}
}

export default wasmTest