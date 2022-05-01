use std::collections::HashMap;
use std::sync::Arc;

use dashmap::DashMap;
use futures::{SinkExt, StreamExt};
use tokio::sync::RwLock;
use uuid::Uuid;
use warp::ws::{Message, WebSocket};

use crate::{Game, Games};

pub async fn client_connection(mut ws: WebSocket, id: String, games: Games) {
	let uuid = &Uuid::new_v4().to_string()[0..8];
	games.get_mut(&id).unwrap().connected_clients += 1;

	if let game = games.get(&id).unwrap() {
		println!("{} connected to {:?}", uuid, game.value());
	}

	while let Some(result) = ws.next().await {
		let msg = match result {
			Ok(msg) => msg,
			Err(e) => {
				eprintln!("error receiving ws message for id: {}: {}", id, e);
				break;
			}
		};
//		client_msg(&id, msg).await;
		println!("received message from {}: {:?}", uuid, msg);
		ws.send(Message::text(msg.to_str().unwrap().to_owned() + " fuf")).await;
	}


	games.get_mut(&id).unwrap().connected_clients -= 1;
	println!("{} disconnected", uuid);
}

async fn client_msg(id: &str, msg: Message) {
	println!("received message from {}: {:?}", id, msg);
	msg.as_bytes();
}