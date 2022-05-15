using backend.Entities;

namespace backend;

public interface IDatabase {
	Task<List<Game>> GetGames(bool onlyVisible);

	Task<Game> GetGame(Guid id);

	Task CreateGame(Game game);

	Task AddPlayer(Guid id);

	Task SubtractPlayer(Guid id);
}