package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell"
)

func main() {
	screen, err := tcell.NewScreen()

	if err != nil {
		log.Fatal("%+v", err)
	}

	if err := screen.Init(); err != nil {
		log.Fatal("%+v", err)
	}

	// Set style for game
	style := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorYellow)
	screen.SetStyle(style)

	snakeBody := SnakeBody{
		X:      5,
		Y:      4,
		Xspeed: 1,
		Yspeed: 0,
	}

	game := Game{
		Screen:    screen,
		snakeBody: snakeBody,
	}

	go game.Run()

	for {
		switch event := game.Screen.PollEvent().(type) {
		case *tcell.EventResize:
			game.Screen.Sync()

		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				game.ScreenFini()
				os.Exit(0)
			} else if event.Key() == tcell.KeyUp {
				game.snakeBody.ChangeDir(-1, 0)
			} else if event.Key() == tcell.Keydown {
				game.snakeBody.ChangeDir(1, 0)
			} else if event.Key() == tcell.KeyLeft {
				game.snakeBody.ChangeDir(0, -1)
			} else if event.Key() == tcell.KeyRight {
				game.snakeBody.ChangeDir(0, 1)
			}
		}
	}
}
