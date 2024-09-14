package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m model) sheetView() string {
	char := m.characters[m.selectedChar]
	s := fmt.Sprintf("Character Sheet for %s:\n\n", char.Name)
	s += fmt.Sprintf("Race: %s\n", char.Race)
	s += fmt.Sprintf("Class: %s\n", char.Class)
	s += fmt.Sprintf("Background: %s\n", char.Background)
	s += fmt.Sprintf("HP: %d\n", char.Hp)
	s += "\nAbility Scores:\n"
	s += fmt.Sprintf("STR: %d  DEX: %d  CON: %d\n", char.AbilityScore.Strength, char.AbilityScore.Dexterity, char.AbilityScore.Constitution)
	s += fmt.Sprintf("INT: %d  WIS: %d  CHA: %d\n", char.AbilityScore.Intelligence, char.AbilityScore.Wisdom, char.AbilityScore.Charisma)
	s += "\nPress enter to return to list"
	return lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1).Render(s)
}
