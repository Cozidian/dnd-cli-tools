package main

import "github.com/charmbracelet/huh"

type model struct {
	state        string
	characters   []Character
	newCharacter Character
	selectedChar int
	form         *huh.Form
	formComplete bool
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
	Class        string
	AbilityScore AbilityScore
	Background   string
	Hp           int
}
