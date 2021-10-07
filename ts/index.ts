// @ts-ignore
import html from 'HTML/index.html'

import "CSS/main.css";


// @ts-ignore
import * as login from './login.ts'

// @ts-ignore
import * as wasm from "../webassembly/pkg";

import {memory} from "../webassembly/pkg/wasm_bg.wasm";


const CELL_SIZE = 5; // px
const GRID_COLOR = "#CCCCCC";
const DEAD_COLOR = "#FFFFFF";
const ALIVE_COLOR = "#000000";


function wasmtest() {
	// Construct the universe, and get its width and height.
	const universe = wasm.Universe.new();
	const width = universe.width();
	const height = universe.height();
	
	let animationId = null;
	
	// Give the canvas room for all of our cells and a 1px border
	// around each of them.
	const canvas: HTMLCanvasElement = document.getElementById("canvas") as HTMLCanvasElement;
	canvas.height = (CELL_SIZE + 1) * height + 1;
	canvas.width = (CELL_SIZE + 1) * width + 1;
	
	const ctx = canvas.getContext('2d');
	
	const isPaused = () => {
		return animationId === null;
	};
	
	const playPauseButton = document.getElementById("play-pause");
	
	const play = () => {
		playPauseButton.textContent = "⏸";
		renderLoop();
	};
	
	const pause = () => {
		playPauseButton.textContent = "▶";
		cancelAnimationFrame(animationId);
		animationId = null;
	};
	
	playPauseButton.addEventListener("click", () => {
		if (isPaused()) {
			play();
		} else {
			pause();
		}
	});
	
	const renderLoop = () => {
		universe.tick();
		
		drawGrid(ctx, height, width);
		drawCells(ctx, universe, height, width);
		
		animationId = requestAnimationFrame(renderLoop);
	};
	
	play();
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

(() => {
	let ip = login.checkLoggedin()
	if (ip) {
		console.log(`passed login with ip:${ip}`)
		
		login.addEvents()
		wasmtest()
	} else {
		console.log("opening login")
		login.createLoginDiv()
	}
})()