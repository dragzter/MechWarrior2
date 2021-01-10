package game

import (
	"fmt"
)

func ReadInput(s string) {

	switch s {
	case "test":
		fmt.Println("Scanner is working.")
	case "menu":
		AdvanceGame()
	case "help":
		fmt.Print("You can type:\n'who' to show your name\n'menu' to enter game menu\n'quit' or 'exit' to quit game\n")
	case "who":
		mainPlayer.Summary()
	case "quit":
		QuitGame()
	case "exit":
		QuitGame()
	default:
		GA.AnnounceRootMenu()
	}
}
