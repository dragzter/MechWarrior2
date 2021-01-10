package game

import (
	"crypto/rand"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

func MakeTextPrompt(lb string) string {
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
func MakeSelectPrompt(gm *MainMenuOptions, lb string) string {
	prompt := promptui.Select{
		Label: mainPlayer.Name + ", " + lb, // lb "What would you like to do?"
		Items: gm.List,                     // List of items so make a menu from
	}
	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "Error"
	}

	return result
}

func SingleSelectPrompt(s []string, lb string) string {
	prompt := promptui.Select{
		Label: mainPlayer.Name + ", " + lb, // lb "What would you like to do?"
		Items: s,
	}
	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "Error"
	}

	return result
}

func MainMenu() {
	if !mainPlayer.HasMech {
		GameMenu.List = append([]string{MenuList["Armory"]}, GameMenu.List...)
		mainPlayer.HasMech = true
	} else {
		GameMenu.List = []string{
			"Main Menu",
			"Quit Game",
		}
	}

	result := MakeSelectPrompt(&GameMenu, "What would you like to do?")

	if result == "Error" {
		fmt.Println("Invalid selection, menu exited.")
	} else {
		fmt.Println("You have selected", result)
	}

	if result == MenuList["Armory"] {
		CreateMech()
	} else if result == MenuList["Quit"] {
		QuitGame()
	} else if result == MenuList["MainMenu"] {
		GA.AnnounceRootMenu()
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

func CreateMech() {
	response := MakeSelectPrompt(&MechChoice, "Select to see more information")

	if response != "" {
		fmt.Println("You have selected:", response)
		DisplayMechInfo(response)
	}
}

func TokenGenerator() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func DisplayMechInfo(st string) {
	selectedMech := &Mech{}

	switch st {
	case starterMech.Name:
		fmt.Println("===============================")
		fmt.Println("\n"+st, "\n------------------", "\nArmor:", starterMech.Armor, "\nHitpoints:", starterMech.Hitpoints, "\nWeapon Slots:", starterMech.WeaponSlots)
		fmt.Println("===============================")
		selectedMech = starterMech
	case starterMechB.Name:
		fmt.Println("===============================")
		fmt.Println("\n"+st, "\n------------------", "\nArmor:", starterMechB.Armor, "\nHitpoints:", starterMechB.Hitpoints, "\nWeapon Slots:", starterMechB.WeaponSlots)
		fmt.Println("===============================")
		selectedMech = starterMechB
	case starterMechC.Name:
		fmt.Println("===============================")
		fmt.Println("\n"+st, "\n------------------", "\nArmor:", starterMechC.Armor, "\nHitpoints:", starterMechC.Hitpoints, "\nWeapon Slots:", starterMechC.WeaponSlots)
		fmt.Println("===============================")
		selectedMech = starterMechC
	}

	selectOptions := MainMenuOptions{
		List: []string{
			"Select Mech",
			"Go Back",
		},
	}

	response := MakeSelectPrompt(&selectOptions, "What would you like to do?")

	if response == "Select Mech" {
		mainPlayer.SetMech(*selectedMech)
		mainPlayer.WeaponizeMech(*autoCannonA)
		//mainPlayer.Summary()
		MainMenu()
	}

	if response == "Go Back" {
		CreateMech()
	}
}
