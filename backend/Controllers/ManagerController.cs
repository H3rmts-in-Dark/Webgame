using backend.dto;
using Microsoft.AspNetCore.Mvc;

namespace backend.Controllers;

public class ManagerController : ControllerBase {
	private readonly IDatabase _database;

	public ManagerController(IDatabase database) {
		_database = database;
	}

	[HttpGet("all")]
	public async Task<IEnumerable<GameDto>> All() {
		return (await _database.GetGames(false)).Select(item => item.ToDto());
	}

	[HttpGet("{id:guid}")]
	public async Task<GameDto> Get(Guid id) {
		return (await _database.GetGame(id)).ToDto();
	}

	[HttpPost("{id:guid}/addPlayer")]
	public async Task AddPlayer(Guid id) {
		await _database.AddPlayer(id);
	}

	[HttpPost("{id:guid}/removePlayer")]
	public async Task SubtractPlayer(Guid id) {
		await _database.SubtractPlayer(id);
	}
}