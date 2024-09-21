package main

import (
	"errors"

	"github.com/charmbracelet/huh"
)

func (m model) builderView() string {
	if m.formComplete {
		if m.builderStep == 1 {
			return "Step 1 complete! Press enter to proceed to step 2."
		} else if m.builderStep == 2 {
			return "Character created! Press enter to view character sheet."
		}
	}
	return m.form.View()
}

func createInitialForm(char *Character) *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Value(&char.Name).
				Key("name").
				Title("What is your character's name?").
				Placeholder("Kreech Blackwood"),
		),
		huh.NewGroup(
			huh.NewSelect[Classes]().
				Title("Choose your class").
				Key("class").
				Options(huh.NewOptions(Barbarian, Bard, Cleric, Druid, Fighter, Monk, Paladin, Ranger, Rouge, Sorcerer, Warlock, Wizard)...).
				Value(&char.Class.ClassType),
		),
	)
}

func createClassSpecificForm(char *Character) *huh.Form {
	switch char.Class.ClassType {
	case Barbarian:
		return huh.NewForm(
			huh.NewGroup(
				huh.NewMultiSelect[string]().
					Title("Choose 2 skill proficiencies").
					Key("skillProficiencies").
					Options(huh.NewOptions("Animal Handling", "Athletics", "Intimidation", "Nature", "Perception", "Survival")...).
					Limit(2).
					Validate(func(str []string) error {
						if len(str) != 2 {
							return errors.New("You must pick 2.")
						}
						return nil
					}).
					Value(&char.Class.SkillProficiencies),
				huh.NewMultiSelect[string]().
					Title("Choose 2 weapon masteries").
					Key("weaponMasteries").
					Options(huh.NewOptions("Battleaxe (Topple)", "Blowgun (Vex)", "Club (Slow)", "Dagger (Nick)", "Dart (Vex)", "Flail (Sap)", "Glaive (Graze)", "Greataxe (Cleave)", "Greatclub (Push)", "Greatsword (Graze)", "Halberd (Cleave)", "Hand Crossbow (Vex)", "Handaxe (Vex)", "Heavy Crossbow (Push)", "Javelin (Slow)", "Lance (Topple)", "Light Crossbow (Slow)", "Light Hammer (Nick)", "Longbow (Slow)", "Longsword (Sap)", "Mace (Sap)", "Maul (Topple)", "Morningstar (Sap)", "Musket (Slow)", "Pike (Push)", "Pistol (Vex)", "Quarterstaff (Topple)", "Rapier (Vex)", "Scimitar (Nick)", "Shortbow (Vex)", "Shortsword (Vex)", "Sickle (Nick)", "Sling (Slow)", "Spear (Sap)", "Trident (Topple)", "War Pick (Sap)", "Warhammer (Push)", "Whip (Slow)")...).
					Limit(2).
					Validate(func(str []string) error {
						if len(str) != 2 {
							return errors.New("You must pick 2.")
						}
						return nil
					}).
					Value(&char.Class.WeaponMasteries),
			),
		)
	case Bard:
		return huh.NewForm(
			huh.NewGroup(
				huh.NewMultiSelect[string]().
					Title("Choose 3 skill proficiencies").
					Key("skillProficiencies").
					Options(huh.NewOptions("Athletics", "Intimidation", "Survival", "Nature")...).
					Limit(3).
					Validate(func(str []string) error {
						if len(str) != 3 {
							return errors.New("You must pick 3.")
						}
						return nil
					}).
					Value(&char.Class.SkillProficiencies),
				huh.NewMultiSelect[string]().
					Title("Choose 3 instruments").
					Key("instruments").
					Options(huh.NewOptions("Lute", "Flute", "Drum", "Lyre")...).
					Limit(3).
					Validate(func(str []string) error {
						if len(str) != 3 {
							return errors.New("You must pick 3.")
						}
						return nil
					}).
					Value(&char.Class.Instruments),
			),
		)
	default:
		return nil
	}
}
