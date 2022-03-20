use std::collections::HashMap;
use std::sync::Arc;

use futures::SinkExt;
use tokio::stream::StreamExt;
use tokio::sync::RwLock;
use warp::ws::{Message, WebSocket};

use crate::Game;

pub async fn client_connection(mut ws: WebSocket, id: String, games: Arc<RwLock<HashMap<String, Game>>>) {
	let mut game = games.read().await.get(&id).unwrap().clone();
	game.connected_clients += 1;
	println!("{} connected to {:?}", id, game);
	games.write().await.insert(id.clone(), game);

	while let Some(result) = ws.next().await {
		let msg = match result {
			Ok(msg) => msg,
			Err(e) => {
				eprintln!("error receiving ws message for id: {}): {}", id, e);
				break;
			}
		};
//		client_msg(&id, msg).await;
		println!("received message from {}: {:?}", id, msg);
		ws.send(Message::text(msg.to_str().unwrap().to_owned() + " fuf")).await;
	}


	let mut game = games.read().await.get(&id).unwrap().clone();
	game.connected_clients -= 1;
	println!("{} disconnected", id);
	games.write().await.insert(id.clone(), game);
}

async fn client_msg(id: &str, msg: Message) {
	println!("received message from {}: {:?}", id, msg);
	msg.as_bytes();
}