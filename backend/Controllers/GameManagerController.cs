using backend.dto;
using Microsoft.AspNetCore.Mvc;

namespace backend.Controllers;

[ApiController]
[ServerHeader]
[Route("games")]
public class GameManagerController : ControllerBase {
	private readonly IDatabase _database;

	public GameManagerController(IDatabase database) {
		_database = database;
	}

	[HttpGet("{id:guid}")]
	public async Task<GameDto> Get(Guid id) {
		return (await _database.GetGame(id)).ToDto();
	}

	[HttpPost("{id:guid}/players/add")]
	public async Task AddPlayer(Guid id) {
		await _database.AddPlayer(id);
	}

	[HttpPost("{id:guid}/players/remove")]
	public async Task SubtractPlayer(Guid id) {
		await _database.SubtractPlayer(id);
	}

	[HttpPost("{id:guid}/players/set")]
	public async Task SetPlayer(Guid id, int? value) {
		await _database.SetPlayer(id, value ?? 0);
	}
}