using backend.dto;
using backend.Entities;
using Microsoft.AspNetCore.Mvc;

namespace backend.Controllers;

[ApiController]
[ServerHeader]
[Route("games")]
public class CreateController : ControllerBase {
	private readonly IDatabase _database;

	public CreateController(IDatabase database) {
		_database = database;
	}

	[HttpPost("create")]
	public async Task<GameDto> Create(CreateGameDto create) {
		var game = Game.FromDto(create);
		await _database.CreateGame(game);
		return game.ToDto();
	}

	[HttpGet("all")]
	public async Task<IEnumerable<GameDto>> All() {
		return (await _database.GetGames()).Select(item => item.ToDto());
	}

	[HttpGet("{id}")]
	public async Task<GameDto> Get(Guid id) {
		return (await _database.GetGame(id)).ToDto();
	}
}