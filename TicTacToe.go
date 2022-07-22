package main

import (
	"github.com/SakthiMahendran/TicTacToe/gameui"
)

func main() {
	gw := gameui.NewGameWindow()

	gw.RenderNewWindow()
	gw.ShowAndRun()
}

