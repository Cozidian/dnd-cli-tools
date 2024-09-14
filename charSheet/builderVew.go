package main

import (
	"os"
	"strconv"

	"github.com/charmbracelet/huh"
)

func (m model) builderView() string {
	if m.formComplete {
		return "Character created! Press enter to view character sheet."
	}
	return m.form.View()
}

func createCharacterForm(char *Character) *huh.Form {
	accessible, _ := strconv.ParseBool(os.Getenv("ACCESSIBLE"))
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Value(&char.Name).
				Key("name").
				Title("Name your character?").
				Placeholder("Kreech Blackwood"),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Options(huh.NewOptions("Aasimar", "Dragonborn", "Dwarf", "Elf", "Gnome", "Goliath", "Halfling", "Human", "Orc", "Tiefling")...).
				Title("Choose your race").
				Key("race").
				Value(&char.Race),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose your class").
				Key("class").
				Options(huh.NewOptions("Barbarian", "Bard", "Cleric", "Druid", "Fighter", "Monk", "Paladin", "Ranger", "Rogue", "Sorcerer", "Warlock", "Wizard")...).
				Value(&char.Class),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose your background").
				Key("background").
				Options(huh.NewOptions("Acolyte", "Artisan", "Charlatan", "Criminal", "Entertainer", "Farmer", "Guard", "Guide", "Hermit", "Merchant", "Noble", "Sage", "Sailor", "Scribe", "Soldier", "Wayfarer")...).
				Value(&char.Background),
		),
		huh.NewGroup(
			huh.NewSelect[int]().
				Title("Strength").
				Options(huh.NewOptions(8, 10, 12, 13, 14, 15)...).
				Key("strength").
				Value(&char.AbilityScore.Strength),
			huh.NewSelect[int]().
				Title("Dexterity").
				Options(huh.NewOptions(8, 10, 12, 13, 14, 15)...).
				Key("dexterity").
				Value(&char.AbilityScore.Dexterity),
			huh.NewSelect[int]().
				Title("Constitution").
				Options(huh.NewOptions(8, 10, 12, 13, 14, 15)...).
				Key("constitution").
				Value(&char.AbilityScore.Constitution),
			huh.NewSelect[int]().
				Title("Intelligence").
				Options(huh.NewOptions(8, 10, 12, 13, 14, 15)...).
				Key("intelligence").
				Value(&char.AbilityScore.Intelligence),
			huh.NewSelect[int]().
				Title("Wisdom").
				Options(huh.NewOptions(8, 10, 12, 13, 14, 15)...).
				Key("wisdom").
				Value(&char.AbilityScore.Wisdom),
			huh.NewSelect[int]().
				Title("Charisma").
				Options(huh.NewOptions(8, 10, 12, 13, 14, 15)...).
				Key("charisma").
				Value(&char.AbilityScore.Charisma),
		),
	).WithAccessible(accessible)
}
