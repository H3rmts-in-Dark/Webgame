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
		if(!client.GetDatabase(settings.Mongo.Db).RunCommandAsync((Command<BsonDocument>) "{ping:1}").Wait(1000)) {
			throw new Exception("Database unreachable");
		}
	}

	public async Task<List<Game>> GetGames() {
		return await _gamesCollection.Find(new BsonDocumentFilterDefinition<Game>(new BsonDocument())).ToListAsync();
	}

	public async Task<Game> GetGame(Guid id) {
		return await _gamesCollection.Find(g => g.Id == id).SingleOrDefaultAsync();
	}

	public async Task CreateGame(Game game) {
		await _gamesCollection.InsertOneAsync(game);
	}

	public async Task AddPlayer(Guid id) {
		await _gamesCollection.UpdateOneAsync(g => g.Id == id, Builders<Game>.Update.Inc(g => g.Players, 1));
	}

	public async Task SubtractPlayer(Guid id) {
		await _gamesCollection.UpdateOneAsync(g => g.Id == id, Builders<Game>.Update.Inc(g => g.Players, -1));
	}
}