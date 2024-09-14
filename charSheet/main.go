package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func initialModel() model {
	m := model{
		state:        "list",
		characters:   []Character{}, // You can pre-populate this with saved characters
		newCharacter: Character{},
		selectedChar: 0,
		formComplete: false,
	}
	m.form = createCharacterForm(&m.newCharacter)
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			switch m.state {
			case "list":
				if m.selectedChar == len(m.characters) {
					m.state = "builder"
					m.formComplete = false
					m.newCharacter = Character{} // Reset new character
					m.form = createCharacterForm(&m.newCharacter)
					return m, m.form.Init()
				} else {
					m.state = "sheet"
				}
			case "builder":
				if m.formComplete {
					m.newCharacter.Name = m.form.GetString("name")
					m.newCharacter.Race = m.form.GetString("race")
					m.newCharacter.Class = m.form.GetString("class")
					m.newCharacter.Background = m.form.GetString("background")
					m.newCharacter.AbilityScore.Strength = m.form.GetInt("strength")
					m.newCharacter.AbilityScore.Dexterity = m.form.GetInt("dexterity")
					m.newCharacter.AbilityScore.Constitution = m.form.GetInt("constitution")
					m.newCharacter.AbilityScore.Intelligence = m.form.GetInt("intelligence")
					m.newCharacter.AbilityScore.Wisdom = m.form.GetInt("wisdom")
					m.newCharacter.AbilityScore.Charisma = m.form.GetInt("charisma")
					m.characters = append(m.characters, m.newCharacter)
					m.selectedChar = len(m.characters) - 1
					m.state = "sheet"
				}
			case "sheet":
				m.state = "list"
			}
		case "esc":
			if m.state == "builder" {
				m.state = "list"
				return m, nil
			}
		case "up", "down":
			if m.state == "list" {
				m.selectedChar += map[string]int{"up": -1, "down": 1}[msg.String()]
				if m.selectedChar < 0 {
					m.selectedChar = len(m.characters)
				} else if m.selectedChar > len(m.characters) {
					m.selectedChar = 0
				}
			}
		}
	}

	if m.state == "builder" {
		form, cmd := m.form.Update(msg)
		if f, ok := form.(*huh.Form); ok {
			m.form = f
		}
		if m.form.State == huh.StateCompleted {
			m.formComplete = true
		}
		return m, cmd
	}

	return m, cmd
}

func (m model) View() string {
	switch m.state {
	case "list":
		return m.listView()
	case "builder":
		return m.builderView()
	case "sheet":
		return m.sheetView()
	default:
		return "Unknown state"
	}
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
