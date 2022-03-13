using backend.dto;

namespace backend.Entities;

public class Game {
	public Guid   Id;
	public bool   Visible;
	public string Code;
	public ushort Limit;
	public string Name;

	public Game(Guid id, bool visible, string code, ushort limit, string name) {
		Id      = id;
		Visible = visible;
		Code    = code;
		Limit   = limit;
		Name    = name;
	}

	public static Game FromDto(CreateGameDto create) {
		var (visible, code, limit, name) = create;
		return new Game(Guid.NewGuid(), visible, code, limit, name);
	}

	public GameDto ToDto() {
		return new GameDto(Id, Limit, Name);
	}
}