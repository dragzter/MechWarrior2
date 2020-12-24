package game

import "fmt"

type Weapon struct {
	Ready      bool
	Name       string
	Damage     int
	Accuracy   float64
	Attacks    int
	Ammunition int
	Owner      string
}

type Player struct {
	Name string
	Id   int
}

type Mech struct {
	Name      string
	Id        int
	Armor     int
	Hitpoints int
	wner      string
}

type Inventory struct {
	Items    []string
	Count    int
	Capacity int
	wner     string
}

type GameState struct {
	PlayerSet bool
	MechIsSet bool
}

type EnemyScanner struct {
	ScanLength int
}

// Player methods
func (p *Player) PrintName() {
	fmt.Println("Player name:", p.Name)
}

func (p *Player) SetName() {
	result := MakeTextPrompt("Name")
	p.Name = result

	fmt.Println("nPlayer Name:", mainPlayer.Name+".", "\nLoading Menu...")
	AdvanceGame()
}
