namespace backend.Settings;

public class MongoDbSettings {
	public  string? ConnectionString => $"mongodb://{Username}:{Password}@{Host}:{Port}";
	public  string? Db               { get; set; }
	public  string? Collection       { get; set; }
	private string? Username         { get; set; }
	private string? Password         { get; set; }
	private string? Host             { get; set; }
	private int?    Port             { get; set; }
}