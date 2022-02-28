using backend.Entities;

namespace backend;

public interface IDatabase {
	Task<List<Game>> GetGames();

	Task CreateGame(Game game);
}