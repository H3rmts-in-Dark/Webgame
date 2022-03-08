namespace backend.Settings;

public class MongoDbSettings {
	public string? ConnectionString => $"mongodb://{Username}:{Password}@{Host}:{Port}";
	public string? Db               { get; set; }
	public string? Collection       { get; set; }
	public string? Username         { get; set; }
	public string? Password         { get; set; }
	public string? Host             { get; set; }
	public int?    Port             { get; set; }
}