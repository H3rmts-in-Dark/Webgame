using backend.Entities;
using Microsoft.AspNetCore.Mvc;

namespace backend.Controllers;

[ApiController]
[ServerHeader]
[Route("create")]
public class CreateController : ControllerBase {
	private readonly IDatabase _database;

	public CreateController(IDatabase database) {
		_database = database;
	}

	[HttpPost("create")]
	public async Task<ActionResult<Game>> Create() {
		var item = new Game(Guid.NewGuid(), "test");
		await _database.CreateGame(item);
		return item;
	}

	[HttpPost("get")]
	public async Task<ActionResult<List<Game>>> Get() {
		return await _database.GetGames();
	}
}