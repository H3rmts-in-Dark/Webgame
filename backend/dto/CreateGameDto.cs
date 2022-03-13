namespace backend.dto;

public record CreateGameDto(bool Visible, string Code, ushort Limit, string Name);