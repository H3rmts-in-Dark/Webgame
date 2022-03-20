use std::collections::HashMap;
use std::convert::Infallible;
use std::sync::Arc;

use tokio::sync::RwLock;
use warp::{Filter, Rejection, Reply};

mod ws;

#[derive(Debug, Clone)]
pub struct Game {
	pub connected_clients: u8,
}

#[tokio::main]
async fn main() {
	let games: Arc<RwLock<HashMap<String, Game>>> = Arc::new(RwLock::new(HashMap::new()));

	games.write().await.insert(String::from("418ef12e-bfaf-4d2d-937b-8aac66988d0f"), Game { connected_clients: 0 }); // hardcoded from database

	let ws_route = warp::path("ws")
			.and(warp::ws())   // 1. param (ws)
			.and(warp::path::param())  // 2. param (game_id)
			.and(with_clients(games.clone()))  // 3. param (games)
			.and_then(ws_handler);

	let routes = ws_route.with(warp::cors().allow_origins(vec!["http://localhost:3000", "http://localhost:3001"]));

	warp::serve(routes).run(([127, 0, 0, 1], 6969)).await;
}

fn with_clients(clients: Arc<RwLock<HashMap<String, Game>>>) -> impl Filter<Extract=(Arc<RwLock<HashMap<String, Game>>>, ), Error=Infallible> + Clone {
	warp::any().map(move || clients.clone())
}

async fn ws_handler(ws: warp::ws::Ws, game_id: String, games: Arc<RwLock<HashMap<String, Game>>>) -> Result<impl Reply, Rejection> {
	if games.read().await.get(&game_id).is_some() {
		Ok(ws.on_upgrade(move |socket| {
			ws::client_connection(socket, game_id, games)
		}))
	} else {
		Err(warp::reject::not_found())
	}
}