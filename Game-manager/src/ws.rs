use futures::{SinkExt, StreamExt};
use uuid::Uuid;
use warp::ws::{Message, WebSocket};

use crate::Games;

pub async fn client_connection(ws: WebSocket, game_id: String, games: Games) {
	let uuid = &Uuid::new_v4().to_string()[24..36]; // get 12 long random
	let (mut sender, mut reciver) = ws.split();

	games.get_mut(&game_id).unwrap().connected_clients += 1;

	match games.get(&game_id) {
		Some(game) => {
			println!("{} connected to {:?}", uuid, game.value())
		}
		None => {
			eprintln!("Error getting {} game from list", game_id);
			return;
		}
	}


	while let Some(result) = reciver.next().await {
		let msg = match result {
			Ok(msg) => msg,
			Err(e) => {
				eprintln!("error receiving ws message for game_id: {}: {}", game_id, e);
				break;
			}
		};

		println!("received message from {}: {:?}", uuid, msg);
		if msg.is_close() {
			break;
		}

		let mess = "Start";
		let res = sender.send(Message::text(mess)).await;
		if res.is_err() {
			eprintln!("error sending ws message for game_id: {}: {}", game_id, res.err().unwrap());
		} else {
			println!("sending {} to {}", mess, uuid);
		}
	}


	games.get_mut(&game_id).unwrap().connected_clients -= 1;

	match games.get(&game_id) {
		Some(game) => {
			println!("{} disconnected from {:?}", uuid, game.value())
		}
		None => {
			eprintln!("Error getting {} game from list", game_id);
			return;
		}
	}
}