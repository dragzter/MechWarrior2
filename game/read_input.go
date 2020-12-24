package game

import (
	"fmt"
)

func ReadInput(s string) {

	switch s {
	case "tes":
		fmt.Println("Scanner is working.")
	case "menu":
		AdvanceGame()
	case "help":
		fmt.Print("You can type:\n'who' to show your name\n'menu' to enter game menu\n'quit' or 'exit' to quit game\n")
	case "who":
		mainPlayer.PrintName()
	case "start":
		if mainPlayer.Name == "" {
			mainPlayer.SetName()
		} else {
			fmt.Println("\nYou have already set your name,", mainPlayer.Name)
		}
	case "quit":
		QuitGame()
	default:
		AnnounceRootMenu()
	}
}
