extern crate web_sys;

use std::fmt;

use wasm_bindgen::prelude::*;
use web_sys::console;


#[cfg(feature = "wee_alloc")]
#[global_allocator]
static ALLOC: wee_alloc::WeeAlloc = wee_alloc::WeeAlloc::INIT;

// joinked from
// https://rustwasm.github.io/docs/book/game-of-life/hello-world.html
// https://github.com/rustwasm/rust-webpack-template/tree/master/template

// curl https://rustwasm.github.io/wasm-pack/installer/init.sh -sSf | sh

// tf is this?
macro_rules! log {
    ( $( $t:tt )* ) => {
        web_sys::console::log_1(&format!( $( $t )* ).into());
    }
}

#[wasm_bindgen]
#[repr(u8)]
#[derive(Clone, Copy, PartialEq)]
pub enum Cell {
	Dead = 0,
	Alive = 1,
}

#[wasm_bindgen]
pub struct Universe {
	width: u32,
	height: u32,
	cells: Vec<Cell>,
}


impl Cell {
	fn toggle(&mut self) {
		*self = match *self {
			Cell::Dead => Cell::Alive,
			Cell::Alive => Cell::Dead,
		};
	}
}

impl Universe {
	fn get_index(&self, row: u32, column: u32) -> usize {
		(row * self.width + column) as usize
	}

	fn live_neighbor_count(&self, row: u32, column: u32) -> u8 {
		let mut count = 0;
		for delta_row in [self.height - 1, 0, 1].iter().cloned() {
			for delta_col in [self.width - 1, 0, 1].iter().cloned() {
				if delta_row == 0 && delta_col == 0 {
					continue;
				}

				let neighbor_row = (row + delta_row) % self.height;
				let neighbor_col = (column + delta_col) % self.width;
				let idx = self.get_index(neighbor_row, neighbor_col);
				count += self.cells[idx] as u8;
			}
		}
		return count;
	}
}

#[wasm_bindgen]
impl Universe {
	pub fn tick(&mut self) {
		let timer = Timer::new("Universe::tick");

		let mut next = self.cells.clone();

		for row in 0..self.height {
			for col in 0..self.width {
				let idx = self.get_index(row, col);
				let cell = self.cells[idx];
				let live_neighbors = self.live_neighbor_count(row, col);

				let next_cell = match (cell, live_neighbors) {
					// Rule 1: Any live cell with fewer than two live neighbours
					// dies, as if caused by underpopulation.
					(Cell::Alive, x) if x < 2 => Cell::Dead,
					// Rule 2: Any live cell with two or three live neighbours
					// lives on to the next generation.
					(Cell::Alive, 2) | (Cell::Alive, 3) => Cell::Alive,
					// Rule 3: Any live cell with more than three live
					// neighbours dies, as if by overpopulation.
					(Cell::Alive, x) if x > 3 => Cell::Dead,
					// Rule 4: Any dead cell with exactly three live neighbours
					// becomes a live cell, as if by reproduction.
					(Cell::Dead, 3) => Cell::Alive,
					// All other cells remain in the same state.
					(otherwise, _) => otherwise,
				};

				next[idx] = next_cell;
			}
		}

		self.cells = next;
		timer.end()
	}

	pub fn toggle_cell(&mut self, row: u32, column: u32) {
		let idx = self.get_index(row, column);
		self.cells[idx].toggle();
	}

	pub fn render(&self) -> String {
		self.to_string()
	}

	pub fn width(&self) -> u32 {
		self.width
	}

	pub fn height(&self) -> u32 {
		self.height
	}

	pub fn cells(&self) -> *const Cell {
		self.cells.as_ptr()
	}
}

impl fmt::Display for Universe {
	fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
		for line in self.cells.as_slice().chunks(self.width as usize) {
			for &cell in line {
				write!(f, "{}", if cell == Cell::Dead { '◻' } else { '◼' })?;
			}
			write!(f, "\n")?;
		}

		return Ok(());
	}
}

#[wasm_bindgen]
impl Universe {
	pub fn new(width: u32, height: u32) -> Universe {
		let cells = (0..width * height).map(|i| {
			if i % 2 == 0 || i % 7 == 0 {
				Cell::Alive
			} else {
				Cell::Dead
			}
		}).collect();

		Universe { width, height, cells }
	}
}

pub struct Timer<'a> {
	name: &'a str,
}

impl<'a> Timer<'a> {
	pub fn new(name: &'a str) -> Timer<'a> {
		console::time_with_label(name);
		Timer { name }
	}

	fn end(&self) {
		console::time_end_with_label(self.name);
	}
}