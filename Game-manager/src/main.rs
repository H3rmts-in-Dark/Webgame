use std::collections::HashMap;
use std::convert::Infallible;
use std::sync::Arc;

use dashmap::DashMap;
use warp::{Filter, Rejection, Reply};

mod ws;

#[derive(Debug, Clone)]
pub struct Game {
	pub connected_clients: u8,
}

type Games = Arc<DashMap<String, Game>>;

#[tokio::main]
async fn main() {
	let games: Games = Arc::new(DashMap::new());

	games.insert(String::from("40283437-12bf-4a85-92fe-d9391223259d"), Game { connected_clients: 0 });// hardcoded from database

	let ws_route = warp::path("ws")
			.and(warp::ws())   // 1. param (ws)
			.and(warp::path::param())  // 2. param (game_id)
			.and(with_clients(games.clone()))  // 3. param (games)
			.and_then(ws_handler);

	let routes = ws_route.with(warp::cors().allow_origins(vec!["http://localhost:3000", "http://localhost:3001"]));

	warp::serve(routes).run(([127, 0, 0, 1], 6969)).await;
}

// ?
fn with_clients(clients: Games) -> impl Filter<Extract=(Games, ), Error=Infallible> + Clone {
	warp::any().map(move || clients.clone())
}

async fn ws_handler(ws: warp::ws::Ws, game_id: String, games: Games) -> Result<impl Reply, Rejection> {
	if games.get(&game_id).is_some() {
		Ok(ws.on_upgrade(move |socket| {
			ws::client_connection(socket, game_id, games)
		}))
	} else {
		println!("game not found {}", game_id);
		Err(warp::reject::not_found())
	}
}