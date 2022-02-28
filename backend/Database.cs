using backend.Entities;
using backend.Settings;
using MongoDB.Bson;
using MongoDB.Driver;

namespace backend;

public class Database : IDatabase {
	private readonly IMongoCollection<Game> _gamesCollection;

	public Database(DbSettings settings) {
		var client = new MongoClient(settings.Mongo.ConnectionString);
		_gamesCollection = client.GetDatabase(settings.Mongo.Db).GetCollection<Game>(settings.Mongo.Collection);
	}

	public async Task<List<Game>> GetGames() {
		return await _gamesCollection.Find(new BsonDocumentFilterDefinition<Game>(new BsonDocument())).ToListAsync();
	}

	public async Task CreateGame(Game game) {
		await _gamesCollection.InsertOneAsync(game);
	}
}