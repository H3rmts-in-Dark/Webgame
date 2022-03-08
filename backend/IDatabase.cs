using backend.Entities;

namespace backend;

public interface IDatabase {
	Task<List<Game>> GetGames();

	Task<Game> GetGame(Guid id);

	Task CreateGame(Game game);
}