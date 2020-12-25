package game

import (
	"bufio"
	"fmt"
	"os"
)

// Init Actors
var mainPlayer = &Player{
	HasMech: false,
}

var MenuList = map[string]string{
	"Armory":   "Select Mech",
	"CheckInv": "Check Inventory",
	"Use":      "Use Item",
	"Repair":   "Repair Mech",
	"MainMenu": "Main Menu",
	"Quit":     "Quit Game",
}

var GameMenu = MainMenuOptions{
	List: []string{
		"Main Menu",
		"Quit Game",
	},
}

var MechChoice = MainMenuOptions{
	List: []string{
		starterMech.Name,
		starterMechB.Name,
		starterMechC.Name,
	},
}

var GA = GameAnnouncer{}

// Play - Starts the game scanner loop
func Play() {
	scanner := bufio.NewScanner(os.Stdin)
	mechwarrior :=
		`
		          _    __    __                _            
	  /\/\   ___  ___| |__/ / /\ \ \__ _ _ __ _ __(_) ___  _ __ 
	 /    \ / _ \/ __| '_ \ \/  \/ / _  | '__| '__| |/ _ \| '__|
	/ /\/\ \  __/ (__| | | \  /\  / (_| | |  | |  | | (_) | |   
	\/    \/\___|\___|_| |_|\/  \/ \__,_|_|  |_|  |_|\___/|_|   
	`
	fmt.Println(mechwarrior)
	fmt.Println("Welcome to MechWarrior, Press Enter to begin.")
	response := SingleSelectPrompt([]string{"Begin Game"}, "Press Enter.")

	if response == "Begin Game" {

		if mainPlayer.Name == "" {
			mainPlayer.SetName()
		} else {
			fmt.Println("\nYou have already set your name,", mainPlayer.Name)
		}

		for scanner.Scan() {
			ReadInput(scanner.Text())
		}
	}

}
