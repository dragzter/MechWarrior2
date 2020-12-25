package game

// Items
var laserWelder = Item{
	Name:        "M47 Welding Laser",
	Description: "Adds 5 to Mech Armor",
	Property: ItemEffect{
		Name:         "High Frequency Laser",
		Description:  "Welds armor cracks.",
		EffectAmount: 5,
	},
}

var repairKit = Item{
	Name:        "Arctus Industries Repair Kit",
	Description: "Restores 10 HP to Mech.",
	Property: ItemEffect{
		Name:         "Assorted parts.",
		Description:  "Theres all kinds of stuff inthere.",
		EffectAmount: 10,
	},
}

var deflector = Item{
	Name:        "Sting Industries Model JL Deflector",
	Description: "The Deflector will absorb all damage taken form a single salvo before it depletes.",
	Property: ItemEffect{
		Name:         "Deflection of enemy fire.",
		Description:  "Makes Mech immune for one round of combat.",
		EffectAmount: 1,
	},
}

// Weapons
var chainBladeA = Weapon{
	Name:        "M9 Vulcan Chain Blade",
	Description: "Vulcan Industries Model A-CF9 Carbon forged chain blade, adopted the shortened designation; M9 for the Military.",
	Damage:      10,
	Accuracy:    90,
	Attacks:     1,
	Ammunition:  20,
	Ready:       true,
}

var autoCannonA = Weapon{
	Name:        "M3 Autocannon",
	Description: "Standard shoulder-mounted.  Fires 50cal rounds at a rate of 300 rounds per minute.",
	Damage:      6,
	Accuracy:    70,
	Attacks:     2,
	Ammunition:  30,
	Ready:       true,
}

var autoCannonB = Weapon{
	Name:        "M3-VSE Variant Autocannon",
	Description: "The VE (Vertical Sensory Engagement) Variant can engage targets faster and fire far more accurately.",
	Damage:      6,
	Accuracy:    92,
	Attacks:     2,
	Ammunition:  40,
	Ready:       true,
}

var autoCannonC = Weapon{
	Name:        "M3-HE Variant Autocannon",
	Description: "HE - High explosive.  Heavier rounds means the HE is less accurate than the VSE variant and carries fewer rounds, but does much more damage.",
	Damage:      24,
	Accuracy:    75,
	Attacks:     2,
	Ammunition:  20,
	Ready:       true,
}

var missilePodA = Weapon{
	Name:        "M90 Lionshead Missile Pod",
	Description: "Chest or shoulder mounted missile battery.",
	Damage:      30,
	Accuracy:    60,
	Attacks:     2,
	Ammunition:  4,
	Ready:       true,
}

/**
=================
| Mech Variants |
=================

Armor is a direct sutraction of damage done.

Example:
Your mech has 10 armor,
Your mech is hit for 20 damage
Your armor absorbs 10 damage, the rest is applied to your remaining hitpoints.

If your armor exceeds the damage you sustained, you take no damage.

Armor Levels:
1. Light
	3 - 10
2. Medium
	11 - 16
3. Heavy
	17 - 24
4. Super Heavy
	25 - 100+


*/

var starterMech = &Mech{
	Name:        "P66 Havoc Combat System (King Crab)",
	Codename:    "Havoc King Crab",
	ID:          TokenGenerator(),
	Armor:       6,
	Hitpoints:   450,
	WeaponSlots: 1,
	Weapons:     map[string]Weapon{},
}

var starterMechB = &Mech{
	Name:        "Stemson S200 Air and Ground Combat Vehicle",
	Codename:    "Jaguar",
	ID:          TokenGenerator(),
	Armor:       5,
	Hitpoints:   400,
	WeaponSlots: 2,
	Weapons:     map[string]Weapon{},
}

var starterMechC = &Mech{
	Name:        "AET-M1 Combat Walker by Hyperdine",
	Codename:    "Jaguar",
	ID:          TokenGenerator(),
	Armor:       8,
	Hitpoints:   380,
	WeaponSlots: 1,
	Weapons:     map[string]Weapon{},
}
