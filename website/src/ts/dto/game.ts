// public record GameDto(Guid Id, ushort Limit, ushort Players, string Name, bool Code);

type Game = {
	id: number
	limit: number
	players: number
	name: string
	code: boolean
}

export type {Game}