package game

import (
	"bufio"
	"fmt"
	"os"
)

// Init Actors
var mainPlayer = Player{}
var mech = Mech{}

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
	fmt.Println("Welcome to MechWarrior, to begin, please type 'start' to enter your name")
	for scanner.Scan() {
		ReadInput(scanner.Text())
	}
}
