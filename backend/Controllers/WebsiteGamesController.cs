using backend.dto;
using backend.Entities;
using Microsoft.AspNetCore.Mvc;

namespace backend.Controllers;

[ApiController]
[ServerHeader]
[Route("website/games")]
public class WebsiteGamesController : ControllerBase {
	private readonly IDatabase _database;

	public WebsiteGamesController(IDatabase database) {
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
		return (await _database.GetGames(true)).Select(item => item.ToDto());
	}

	[HttpGet("{id:guid}")]
	public async Task<GameDto> Get(Guid id) {
		return (await _database.GetGame(id)).ToDto();
	}

	[HttpGet("{id:guid}/check")]
	public async Task<CheckCodes> Check(Guid id, string? code) {
		return (await _database.GetGame(id)).Check(code ?? "");
	}
}