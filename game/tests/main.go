package main

// Exploring ideas with inventory system
// Rough Draft

import "fmt"

type effect struct {
	desc         string
	effectAmount int
}

type item struct {
	name     string
	property effect
}

type inventory struct {
	items map[string]item
}

type player struct {
	damage int
	hp     int
	name   string
}

var inv = inventory{
	items: map[string]item{},
}

func main() {

	autoCannon := item{
		name: "Autocannon",
		property: effect{
			desc:         "Adds 20 to player damage",
			effectAmount: 20,
		},
	}

	repairKit := item{
		name: "Repair Kit",
		property: effect{
			desc:         "Restores health to player",
			effectAmount: 45,
		},
	}

	p1 := player{
		hp:     30,
		name:   "Harold",
		damage: 40,
	}

	p1.increaseDamage(autoCannon)

	inv.addItem(autoCannon)
	inv.addItem(repairKit)
	printInv()
	p1.restoreLife(repairKit)
	inv.removeItem(repairKit)
	printInv()

}

func (p *player) increaseDamage(i item) {
	p.damage += i.property.effectAmount
}

func (p *player) restoreLife(i item) {
	p.hp += i.property.effectAmount
}

func (inv *inventory) addItem(it item) {
	inv.items[it.name] = it
	fmt.Println(inv.items)
}

func (inv *inventory) removeItem(it item) {
	_, ok := inv.items[it.name]
	if ok {
		delete(inv.items, it.name)
	}
}

func printInv() {
	fmt.Println(inv)
}
