package main

import "github.com/charmbracelet/huh"

type model struct {
	state        string
	characters   []Character
	newCharacter Character
	selectedChar int
	form         *huh.Form
	formComplete bool
	builderStep  int
}

type AbilityScore struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
}

type Character struct {
	Name         string
	Lvl          int
	Race         string
	Class        Class
	AbilityScore AbilityScore
	Background   string
	Hp           int
}

type Class struct {
	ClassType          Classes
	SkillProficiencies []string
	WeaponMasteries    []string
	Instruments        []string
	// Add other fields as needed
}

type Classes string

const (
	Barbarian Classes = "Barbarian"
	Bard      Classes = "Bard"
	Cleric    Classes = "Cleric"
	Druid     Classes = "Druid"
	Fighter   Classes = "Fighter"
	Monk      Classes = "Monk"
	Paladin   Classes = "Paladin"
	Ranger    Classes = "Ranger"
	Rouge     Classes = "Rouge"
	Sorcerer  Classes = "Sorcerer"
	Warlock   Classes = "Warlock"
	Wizard    Classes = "Wizard"
)

type ClassOptions struct {
	SkillProficiencies        []string
	SkillProficienciesChoices int
	WeaponMasteries           []string
	WeaponChoices             int
	Languages                 []string
	LanguagesChoices          int
	Instruments               []string
	InstrumentChoices         int
	ToolsAndInstruments       []string
	ToolsAndInstrumentChoices int
	DivineOrder               []string
	DivineOrderChoices        int
	PrimalOrder               []string
	PrimalOrderChoices        int
	FightingStyle             []string
	FightingStyleChoices      int
	Expertise                 []string
	ExpertiseChoices          int
	EldrichInvocations        []string
	EldrichInvocationsChoices int
}

type OptionGroup struct {
	Options []string
	Choices int
}

var commonSkills = []string{
	"Acrobatics", "Animal Handling", "Arcana", "Athletics",
	"Deception", "History", "Insight", "Intimidation",
	"Investigation", "Medicine", "Nature", "Perception",
	"Performance", "Persuasion", "Religion", "Sleight of Hand",
	"Stealth", "Survival",
}

var allWeapons = []string{
	"Battleaxe (Topple)", "Blowgun (Vex)", "Club (Slow)", "Dagger (Nick)",
	"Dart (Vex)", "Flail (Sap)", "Glaive (Graze)", "Greataxe (Cleave)",
	"Greatclub (Push)", "Greatsword (Graze)", "Halberd (Cleave)",
	"Hand Crossbow (Vex)", "Handaxe (Vex)", "Heavy Crossbow (Push)",
	"Javelin (Slow)", "Lance (Topple)", "Light Crossbow (Slow)",
	"Light Hammer (Nick)", "Longbow (Slow)", "Longsword (Sap)", "Mace (Sap)",
	"Maul (Topple)", "Morningstar (Sap)", "Musket (Slow)", "Pike (Push)",
	"Pistol (Vex)", "Quarterstaff (Topple)", "Rapier (Vex)", "Scimitar (Nick)",
	"Shortbow (Vex)", "Shortsword (Vex)", "Sickle (Nick)", "Sling (Slow)",
	"Spear (Sap)", "Trident (Topple)", "War Pick (Sap)", "Warhammer (Push)", "Whip (Slow)",
}

var allInstruments = []string{
	"Bagpipes", "Drum", "Flute", "Lute", "Lyre", "Horn", "Pan flute", "Shawm",
	"Viol",
}

var allLanguages = []string{
	"Common", "Dwarvish", "Elvish", "Giant", "Gnomish",
	"Goblin", "Halfling", "Orc", "Undercommon",
}

// Define other common options as needed

var classOptionsMap = map[Classes]ClassOptions{
	Barbarian: {
		SkillProficiencies:        []string{"Animal Handling", "Athletics", "Intimidation", "Nature", "Perception", "Survival"},
		SkillProficienciesChoices: 2,
		WeaponMasteries:           []string{},
		WeaponChoices:             2,
	},
	Bard: {
		SkillProficiencies:        []string{"Athletics", "Intimidation", "Survival", "Nature"},
		SkillProficienciesChoices: 3,
		Instruments:               []string{"Lute", "Flute", "Drum", "Lyre"},
		InstrumentChoices:         3,
	},
	Cleric: {
		SkillProficiencies:        []string{"Athletics", "Intimidation", "Survival", "Nature"},
		SkillProficienciesChoices: 2,
		DivineOrder:               []string{"Lute", "Flute", "Drum", "Lyre"},
		DivineOrderChoices:        1,
	},
	Druid: {
		SkillProficiencies:        []string{"Athletics", "Intimidation", "Survival", "Nature"},
		SkillProficienciesChoices: 2,
		PrimalOrder:               []string{"Lute", "Flute", "Drum", "Lyre"},
		PrimalOrderChoices:        1,
	},
	Fighter: {
		SkillProficiencies:        []string{"Athletics", "Intimidation", "Survival", "Nature"},
		SkillProficienciesChoices: 2,
		FightingStyle:             []string{"Athletics", "Intimidation", "Survival", "Nature"},
		FightingStyleChoices:      1,
		WeaponMasteries:           []string{"Battleaxe (Topple)", "Blowgun (Vex)", "Club (Slow)", "Dagger (Nick)", "Dart (Vex)", "Flail (Sap)", "Glaive (Graze)", "Greataxe (Cleave)", "Greatclub (Push)", "Greatsword (Graze)", "Halberd (Cleave)", "Hand Crossbow (Vex)", "Handaxe (Vex)", "Heavy Crossbow (Push)", "Javelin (Slow)", "Lance (Topple)", "Light Crossbow (Slow)", "Light Hammer (Nick)", "Longbow (Slow)", "Longsword (Sap)", "Mace (Sap)", "Maul (Topple)", "Morningstar (Sap)", "Musket (Slow)", "Pike (Push)", "Pistol (Vex)", "Quarterstaff (Topple)", "Rapier (Vex)", "Scimitar (Nick)", "Shortbow (Vex)", "Shortsword (Vex)", "Sickle (Nick)", "Sling (Slow)", "Spear (Sap)", "Trident (Topple)", "War Pick (Sap)", "Warhammer (Push)", "Whip (Slow)"},
		WeaponChoices:             3,
	},
	Monk: {
		SkillProficiencies:        []string{"Athletics", "Intimidation", "Survival", "Nature"},
		SkillProficienciesChoices: 2,
		ToolsAndInstruments:       []string{"Lute", "Flute", "Drum", "Lyre"},
		ToolsAndInstrumentChoices: 3,
	},
	Paladin: {
		SkillProficiencies:        []string{"Animal Handling", "Athletics", "Intimidation", "Nature", "Perception", "Survival"},
		SkillProficienciesChoices: 2,
		WeaponMasteries:           []string{"Battleaxe (Topple)", "Blowgun (Vex)", "Club (Slow)", "Dagger (Nick)", "Dart (Vex)", "Flail (Sap)", "Glaive (Graze)", "Greataxe (Cleave)", "Greatclub (Push)", "Greatsword (Graze)", "Halberd (Cleave)", "Hand Crossbow (Vex)", "Handaxe (Vex)", "Heavy Crossbow (Push)", "Javelin (Slow)", "Lance (Topple)", "Light Crossbow (Slow)", "Light Hammer (Nick)", "Longbow (Slow)", "Longsword (Sap)", "Mace (Sap)", "Maul (Topple)", "Morningstar (Sap)", "Musket (Slow)", "Pike (Push)", "Pistol (Vex)", "Quarterstaff (Topple)", "Rapier (Vex)", "Scimitar (Nick)", "Shortbow (Vex)", "Shortsword (Vex)", "Sickle (Nick)", "Sling (Slow)", "Spear (Sap)", "Trident (Topple)", "War Pick (Sap)", "Warhammer (Push)", "Whip (Slow)"},
		WeaponChoices:             2,
	},
	Ranger: {
		SkillProficiencies:        []string{"Animal Handling", "Athletics", "Intimidation", "Nature", "Perception", "Survival"},
		SkillProficienciesChoices: 2,
		WeaponMasteries:           []string{"Battleaxe (Topple)", "Blowgun (Vex)", "Club (Slow)", "Dagger (Nick)", "Dart (Vex)", "Flail (Sap)", "Glaive (Graze)", "Greataxe (Cleave)", "Greatclub (Push)", "Greatsword (Graze)", "Halberd (Cleave)", "Hand Crossbow (Vex)", "Handaxe (Vex)", "Heavy Crossbow (Push)", "Javelin (Slow)", "Lance (Topple)", "Light Crossbow (Slow)", "Light Hammer (Nick)", "Longbow (Slow)", "Longsword (Sap)", "Mace (Sap)", "Maul (Topple)", "Morningstar (Sap)", "Musket (Slow)", "Pike (Push)", "Pistol (Vex)", "Quarterstaff (Topple)", "Rapier (Vex)", "Scimitar (Nick)", "Shortbow (Vex)", "Shortsword (Vex)", "Sickle (Nick)", "Sling (Slow)", "Spear (Sap)", "Trident (Topple)", "War Pick (Sap)", "Warhammer (Push)", "Whip (Slow)"},
		WeaponChoices:             2,
	},
	Rouge: {
		SkillProficiencies:        []string{"Animal Handling", "Athletics", "Intimidation", "Nature", "Perception", "Survival"},
		SkillProficienciesChoices: 2,
		Expertise:                 []string{"Animal Handling", "Athletics", "Intimidation", "Nature", "Perception", "Survival"},
		ExpertiseChoices:          2,
		Languages:                 []string{"Animal Handling", "Athletics", "Intimidation", "Nature", "Perception", "Survival"},
		LanguagesChoices:          2,
		WeaponMasteries:           []string{"Battleaxe (Topple)", "Blowgun (Vex)", "Club (Slow)", "Dagger (Nick)", "Dart (Vex)", "Flail (Sap)", "Glaive (Graze)", "Greataxe (Cleave)", "Greatclub (Push)", "Greatsword (Graze)", "Halberd (Cleave)", "Hand Crossbow (Vex)", "Handaxe (Vex)", "Heavy Crossbow (Push)", "Javelin (Slow)", "Lance (Topple)", "Light Crossbow (Slow)", "Light Hammer (Nick)", "Longbow (Slow)", "Longsword (Sap)", "Mace (Sap)", "Maul (Topple)", "Morningstar (Sap)", "Musket (Slow)", "Pike (Push)", "Pistol (Vex)", "Quarterstaff (Topple)", "Rapier (Vex)", "Scimitar (Nick)", "Shortbow (Vex)", "Shortsword (Vex)", "Sickle (Nick)", "Sling (Slow)", "Spear (Sap)", "Trident (Topple)", "War Pick (Sap)", "Warhammer (Push)", "Whip (Slow)"},
		WeaponChoices:             2,
	},
	Sorcerer: {
		SkillProficiencies:        []string{"Animal Handling", "Athletics", "Intimidation", "Nature", "Perception", "Survival"},
		SkillProficienciesChoices: 2,
	},
	Warlock: {
		SkillProficiencies:        []string{"Animal Handling", "Athletics", "Intimidation", "Nature", "Perception", "Survival"},
		SkillProficienciesChoices: 2,
		EldrichInvocations:        []string{"Animal Handling", "Athletics", "Intimidation", "Nature", "Perception", "Survival"},
		EldrichInvocationsChoices: 2,
	},
	Wizard: {
		SkillProficiencies:        []string{"Animal Handling", "Athletics", "Intimidation", "Nature", "Perception", "Survival"},
		SkillProficienciesChoices: 2,
	},
}
