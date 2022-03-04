using backend.dto;

namespace backend.Entities;

public record Game(Guid Id, string Code, uint Limit) {
	public static Game FromDto(CreateGameDto create) {
		var (code, limit) = create;
		return new Game(Guid.NewGuid(), code, limit);
	}

	public GameDto ToDto() {
		return new GameDto(Id, Code, Limit);
	}
}