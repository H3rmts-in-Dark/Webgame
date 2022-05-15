// public record GameDto(Guid Id, ushort Limit, ushort Players, string Name);

type Game = {
	id: number
	limit: number
	players: number
	name: string
}

export type {Game}