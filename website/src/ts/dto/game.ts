// public record GameDto(Guid Id, ushort Limit, string Name);

type Game = {
	id: number
	limit: number
	name: string
}

export type {Game}