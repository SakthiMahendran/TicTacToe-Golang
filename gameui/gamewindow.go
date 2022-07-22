package gameui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/SakthiMahendran/TicTacToe/gameboard"
)

func NewGameWindow() GameWindow {
	gw := GameWindow{}

	gw.game = app.New()
	gw.window = gw.game.NewWindow("TicTacToe")

	return gw
}

type GameWindow struct {
	gameBoard *gameboard.GameBoard
	game      fyne.App
	window    fyne.Window
	btns      [9]*widget.Button
	lbl       *widget.Label
	player    string
}

func (gw *GameWindow) RenderNewWindow() {
	gw.player = "X"
	gw.gameBoard = &gameboard.GameBoard{}
	gw.lbl = widget.NewLabel(gw.player + "'s turn")

	btnGrid := container.NewGridWithColumns(3)

	for i := 0; i < 9; i++ {
		gw.btns[i] = widget.NewButton("", gw.newTappedFunc(i))
		btnGrid.Add(gw.btns[i])
	}

	panel := container.NewVSplit(gw.lbl, btnGrid)
	panel.SetOffset(0)

	gw.window.Resize(fyne.NewSize(600, 600))

	gw.window.SetContent(panel)
}

func (gw *GameWindow) newTappedFunc(i int) func() {
	btnIndex := i

	return func() {
		gw.btns[btnIndex].SetText(gw.player)
		gw.btns[btnIndex].Disable()

		row, col := gw.trans(btnIndex)

		gw.gameBoard.PutChar(gw.player[0], row, col)

		gStat := gw.gameBoard.IsOver()

		if len(gStat) == 0 && gStat != nil {
			gw.updateLabel("Draw")
			gw.diableAllBtns()
			return
		} else if gStat != nil {
			gw.updateLabel(gw.player + " Won!!!")
			gw.diableAllBtns()
			return
		}

		gw.togglePlayer()
		gw.updateLabel(gw.player + "'s turn")

	}
}

func (gw *GameWindow) togglePlayer() {
	if gw.player == "X" {
		gw.player = "O"
	} else if gw.player == "O" {
		gw.player = "X"
	}
}

func (gw *GameWindow) trans(i int) (int, int) {
	var row, col int

	if i < 3 {
		row = 0
		col = i
	} else if i < 6 {
		row = 1
		col = i - 3
	} else if i < 9 {
		row = 2
		col = i - 6
	}

	return row, col
}

func (gw *GameWindow) diableAllBtns() {
	for _, btn := range gw.btns {
		btn.Disable()
	}
}

func (gw *GameWindow) updateLabel(txt string) {
	gw.lbl.SetText(txt)
}

func (gw *GameWindow) ShowAndRun() {
	gw.window.ShowAndRun()
}
