package game

import "fmt"

// Player - The main player
type Player struct {
	Name string
	ID   int
}

// Weapon - can be equipped by Mech
type Weapon struct {
	Ready       bool
	Name        string
	Description string
	Damage      int
	Accuracy    int
	Attacks     int
	Ammunition  int
}

// Mech - Owned by player, has weapon slots, when mech is destroyed, game ends
type Mech struct {
	Name        string
	ID          string
	Armor       int
	Hitpoints   int
	WeaponSlots int
	Weapons     map[string]Weapon
}

// Inventory - Holds items, has an owner
type Inventory struct {
	Items    map[string]Item
	Owner    string
	Capacity int
}

// Item - Has properties made up of effects, can be stored in inventory
type Item struct {
	Name        string
	Description string
	Property    ItemEffect // Can only have 1 effect for now
}

// ItemEffect - Does something to a player or computer, must be part of item
type ItemEffect struct {
	Name         string
	Description  string
	EffectAmount int
}

type GameState struct {
	PlayerSet bool
	MechIsSet bool
}

/**
 	When the player performs a scan, he may encounter
	a random event.  Events will have codes, and based
	on the code, the event will triggered.  Events will also
	trigger when player advances.

	Possible events:
	1. Enemy found (should be most common, becomes more common as player advances)
	2. Item found
	3. Damage sustained from travel (less common)
	4. Random rpg (flavor) event. (You meet a person they ask for help, your mech takes damage or gets an item etc.)
*/

type EventList struct {
	PossibleEvents map[string]EventItem // an event will be chosen at random and the appropriate method will begin
}

type EventItem struct {
	// e.g. Code 34 - new enemy etc.
	Code int
	Name string
}

// Inventory Method
func (inv *Inventory) AddItem(it Item) {
	inv.Items[it.Name] = it
	fmt.Println(inv.Items)
}

func (inv *Inventory) RemoveItem(i Item) {
	_, ok := inv.Items[i.Name]
	if ok {
		delete(inv.Items, i.Name)
	}
}

// Player methods
func (p *Player) PrintName() {
	fmt.Println("Player name:", p.Name)
}

func (p *Player) SetName() {
	result := MakeTextPrompt("Name")
	p.Name = result

	fmt.Println("Player Name:", mainPlayer.Name+".", "\nLoading Menu...")
	AdvanceGame()
}
