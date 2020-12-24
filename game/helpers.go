package game

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/manifoldco/promptui"
)

func MakeTextPrompt(lb string) string {

	// lb = "What is our name?"
	prompt := promptui.Prompt{
		Label: lb,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	return result
}

// MakeSelectPrompt - Wrapper for promtpui select
func MakeSelectPrompt(si []string, lb string) string {
	prompt := promptui.Select{
		Label: mainPlayer.Name + ", " + lb, // lb "What would you like to do?"
		Items: si,                          // List of items so make a menu from
	}
	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "Error"
	}

	return result
}

func MainMenu() {
	list := []string{
		"Board Mech",
		"Check Inventory",
		"Quit",
		"Go Back",
	}
	result := MakeSelectPrompt(list, "What would you like to do?")

	if result == "Error" {
		fmt.Println("Invalid selection, menu exited.")
	} else {
		fmt.Println("You have selected", result)
	}

	if result == "Board Mech" {
		CreateMech(&mech)
	} else if result == "Quit" {
		QuitGame()
	} else if result == "Go Back" {
		AnnounceRootMenu()
	}
}

func AdvanceGame() {
	MainMenu()
}

func PerformScan() {
	fmt.Println("Scanning...")
}

func QuitGame() {

	result := MakeTextPrompt("Quit Game? (y,n)")

	if result == "y" {
		os.Exit(1)
	} else if result == "n" {
		AdvanceGame()
	} else {
		fmt.Println("Invalid Input")
		QuitGame()
	}
}

func AnnounceRootMenu() {
	fmt.Println("\nAt Main screen\nType 'menu' to enter game menu. Type 'help' to see all options")
}

func CreateMech(m *Mech) {
	makeId := rand.Intn(1000)
	m.Id = makeId
	fmt.Println("Created Mech")
}
