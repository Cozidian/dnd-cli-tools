package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func printChar(char Character) {
	outerBox := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")).
		Padding(1, 2).
		Width(50)

	innerBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("212")).
		Padding(1, 1).
		MarginTop(1)

	title := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("211")).
		MarginBottom(1)

	keyword := lipgloss.NewStyle().
		Foreground(lipgloss.Color("212"))

	// Print character summary.
	{
		var sb strings.Builder
		fmt.Fprintf(&sb, "%s\n\n", title.Render("Character Sheet"))
		fmt.Fprintf(&sb, "Name: %s\n", keyword.Render(char.Name))
		fmt.Fprintf(&sb, "Race: %s\n", keyword.Render(char.Race))
		fmt.Fprintf(&sb, "Class: %s\n", keyword.Render(char.Class))
		fmt.Fprintf(&sb, "Background: %s\n", keyword.Render(char.Background))

		// Build ability scores
		abilityScores := []struct {
			name  string
			score int
		}{
			{"STR", char.AbilityScore.Strength},
			{"DEX", char.AbilityScore.Dexterity},
			{"CON", char.AbilityScore.Constitution},
			{"INT", char.AbilityScore.Intelligence},
			{"WIS", char.AbilityScore.Wisdom},
			{"CHA", char.AbilityScore.Charisma},
		}

		var abilityBoxes []string
		for _, ability := range abilityScores {
			box := innerBox.Render(fmt.Sprintf("%s\n\n%d", ability.name, ability.score))
			abilityBoxes = append(abilityBoxes, box)
		}

		// Combine ability score boxes
		abilityRow := lipgloss.JoinHorizontal(lipgloss.Top, abilityBoxes...)

		// Combine all elements
		innerContent := fmt.Sprintf("%s\n%s\n%s",
			sb.String(),
			title.Render("Ability Scores"),
			abilityRow)

		// Render the final nested box
		fmt.Println(outerBox.Render(innerContent))
	}
}
