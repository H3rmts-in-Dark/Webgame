namespace backend.dto;

public record GameDto(Guid Id, ushort Limit, ushort Players, string Name);