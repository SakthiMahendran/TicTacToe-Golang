package gameboard

func NewGameBoard() GameBoard {
	gb := GameBoard{}
	return gb
}

type GameBoard struct {
	board [3][3]byte
}

func (gb *GameBoard) IsEmpty(row, col int) bool {
	return gb.board[row][col] == 0
}

func (gb *GameBoard) PutChar(char byte, x, y int) {
	if gb.board[x][y] == 0 {
		gb.board[x][y] = char
	}
}

func (gb *GameBoard) GetChar(row, col int) byte {
	return gb.board[row][col]
}

func (gb *GameBoard) IsOver() [][]byte {

	row := gb.checkRows()
	if row != nil {
		return row
	}

	col := gb.checkColumns()
	if col != nil {
		return col
	}

	dia := gb.checkDiagonals()
	if dia != nil {
		return dia
	}

	if gb.isBoardFilled() {
		return [][]byte{}
	}

	return nil
}

func (gb *GameBoard) Reset() {
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			if gb.board[i][j] != 0 {
				gb.board[i][j] = 0
			}
		}
	}
}

func (gb *GameBoard) isBoardFilled() bool {
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			if gb.board[i][j] == 0 {
				return false
			}
		}
	}

	return true
}

func (gb *GameBoard) checkRows() [][]byte {
	for i := byte(0); i <= 2; i++ {
		row := gb.checkRow(i)

		if row != nil {
			return row
		}
	}

	return nil
}

func (gb *GameBoard) checkColumns() [][]byte {
	for i := byte(0); i <= 2; i++ {
		col := gb.checkColumn(i)

		if col != nil {
			return col
		}
	}

	return nil
}

func (gb *GameBoard) checkDiagonals() [][]byte {
	if gb.board[0][0] != 0 && gb.board[0][0] == gb.board[1][1] && gb.board[0][0] == gb.board[2][2] {
		return [][]byte{{0, 0}, {1, 1}, {2, 2}}
	} else if gb.board[0][2] != 0 && gb.board[0][2] == gb.board[1][1] && gb.board[0][2] == gb.board[2][0] {
		return [][]byte{{0, 2}, {1, 1}, {2, 0}}
	}

	return nil
}

func (gb *GameBoard) checkRow(row byte) [][]byte {
	if gb.board[row][0] != 0 && gb.board[row][0] == gb.board[row][1] && gb.board[row][0] == gb.board[row][2] {
		return [][]byte{{row, 0}, {row, 1}, {row, 2}}
	}

	return nil
}

func (gb *GameBoard) checkColumn(col byte) [][]byte {
	if gb.board[0][col] != 0 && gb.board[0][col] == gb.board[1][col] && gb.board[0][col] == gb.board[2][col] {
		return [][]byte{{0, col}, {1, col}, {2, col}}
	}

	return nil
}
