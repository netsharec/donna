package lape

import (
        `bytes`
)

type Game struct {
	board	[64]Piece
	players	[2]*Player
        current int
}

func (g *Game)Initialize() *Game {
        g.players[0] = new(Player).Initialize(0)
        g.players[1] = new(Player).Initialize(1)
        g.current = 0

        return g
}

func (g *Game)Set(row, col int, piece Piece) *Game {
        g.board[Index(row, col)] = piece

        return g
}

func (g *Game)Get(row, col int) Piece {
        return g.board[Index(row, col)]
}

func (g *Game)SetInitialPosition() *Game {

        for color := 0;  color < 2; color ++ {
                g.Set(color * 7, 0, Rook(color))
                g.Set(color * 7, 1, Knight(color))
                g.Set(color * 7, 2, Bishop(color))
                g.Set(color * 7, 3, Queen(color))
                g.Set(color * 7, 4, King(color))
                g.Set(color * 7, 5, Bishop(color))
                g.Set(color * 7, 6, Knight(color))
                g.Set(color * 7, 7, Rook(color))
                for col := 0; col <= 7; col++ {
                        g.Set(color * 5 + 1, col, Pawn(color))
                }
        }

        return g
}

func (g *Game)ToString() string {
	buffer := bytes.NewBufferString("  a b c d e f g h\n")
	for row := 7;  row >= 0;  row-- {
		buffer.WriteByte('1' + byte(row))
		for col := 0;  col <= 7; col++ {
			index := Index(row, col)
			buffer.WriteByte(' ')
			if piece := g.board[index]; piece != 0 {
				buffer.WriteString(piece.ToString())
			} else {
				buffer.WriteString("\u22C5")
			}
		}
		buffer.WriteByte('\n')
	}
	return buffer.String()
}
