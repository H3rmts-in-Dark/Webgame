// public record CreateGameDto(bool Visible, string Code, ushort Limit, string Name);

type CreateGame = {
	visible: boolean
	code: string
	limit: number
	name: string
}

export type {CreateGame}