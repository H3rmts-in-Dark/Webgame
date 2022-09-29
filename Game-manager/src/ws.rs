use futures::{SinkExt, StreamExt};
use futures::stream::{SplitSink, SplitStream};
use uuid::Uuid;
use warp::ws::{Message, WebSocket};

use crate::Games;

pub async fn client_connection(ws: WebSocket, game_id: String, games: Games) {
	let uuid = &Uuid::new_v4().to_string()[24..36]; // get 12 long random to identify connection
	let (mut sender, mut reciver): (SplitSink<WebSocket, Message>, SplitStream<WebSocket>) = ws.split();

	match games.get_mut(&game_id) {
		Some(mut game) => {
			game.connected_clients += 1;
			println!("{} connected to {}, connected clients: {}", uuid, game_id, game.connected_clients);
		}
		None => {
			eprintln!("{} game not found {}", uuid, game_id);
			return;
		}
	}


	while let Some(result) = reciver.next().await {
		let msg = match result {
			Ok(msg) => msg,
			Err(e) => {
				eprintln!("{} error receiving ws message: {}", uuid, e);
				break;
			}
		};

		println!("{} received message: {:?}", uuid, msg);

		if msg.is_close() {
			println!("{} received close", uuid);
			break;
		}

		let mess = "Start";

		match sender.send(Message::text(mess)).await {
			Ok(_) => {
				println!("{} sent message: {}", uuid, mess)
			}
			Err(e) => {
				eprintln!("{} error sending ws message: {}", uuid, e);
				break;
			}
		}
	}

	match games.get_mut(&game_id) {
		Some(mut game) => {
			game.connected_clients -= 1;
			println!("{} disconnected from {}, connected clients: {}", uuid, game_id, game.connected_clients);
		}
		None => {
			eprintln!("{} game not found {}", uuid, game_id);
			return;
		}
	}
}