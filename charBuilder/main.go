
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// Define the different states of the application
type state int

const (
	stateMainMenu state = iota
	stateListCharacters
	stateCreateCharacter
	stateViewCharacter
)

// Character struct represents a DnD character
type Character struct {
	Name          string            `json:"name"`
	Race          string            `json:"race"`
	Class         string            `json:"class"`
	AbilityScores map[string]int    `json:"ability_scores"`
	Skills        []string          `json:"skills"`
	Equipment     []string          `json:"equipment"`
}

// Model represents the application's state
type Model struct {
	state           state
	characters      []Character
	selectedCharIdx int
	cursor          int
	character       Character
	step            int
	input           string
	err             error
}

// InitialModel returns the initial state of the application
func InitialModel() Model {
	return Model{
		state: stateMainMenu,
	}
}

// Init is the initial function for Bubble Tea
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the model accordingly
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.state {
	case stateMainMenu:
		return m.updateMainMenu(msg)
	case stateListCharacters:
		return m.updateListCharacters(msg)
	case stateCreateCharacter:
		return m.updateCreateCharacter(msg)
	case stateViewCharacter:
		return m.updateViewCharacter(msg)
	default:
		return m, nil
	}
}

// View renders the UI based on the current state
func (m Model) View() string {
	switch m.state {
	case stateMainMenu:
		return m.viewMainMenu()
	case stateListCharacters:
		return m.viewListCharacters()
	case stateCreateCharacter:
		return m.viewCreateCharacter()
	case stateViewCharacter:
		return m.viewViewCharacter()
	default:
		return "Unknown state"
	}
}

// Main Menu Update
func (m Model) updateMainMenu(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < 1 {
				m.cursor++
			}
		case "enter":
			if m.cursor == 0 {
				// List Characters
				m.state = stateListCharacters
				m.cursor = 0
				m.loadCharacters()
			} else if m.cursor == 1 {
				// Create Character
				m.state = stateCreateCharacter
				m.step = 0
				m.input = ""
				m.character = Character{
					AbilityScores: make(map[string]int),
				}
			}
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

// Main Menu View
func (m Model) viewMainMenu() string {
	s := "Welcome to the DnD Character Builder\n\n"
	menuOptions := []string{"List Characters", "Create New Character"}
	for i, option := range menuOptions {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, option)
	}
	s += "\nPress 'q' to quit."
	return s
}

// List Characters Update
func (m Model) updateListCharacters(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.characters)-1 {
				m.cursor++
			}
		case "enter":
			// View selected character
			m.state = stateViewCharacter
			m.selectedCharIdx = m.cursor
			m.cursor = 0
		case "b":
			// Back to main menu
			m.state = stateMainMenu
			m.cursor = 0
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

// List Characters View
func (m Model) viewListCharacters() string {
	if len(m.characters) == 0 {
		return "No characters found.\n\nPress 'b' to go back."
	}
	s := "Select a character to view\n\n"
	for i, char := range m.characters {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, char.Name)
	}
	s += "\nPress 'b' to go back."
	return s
}

// Create Character Update
func (m Model) updateCreateCharacter(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.step {
		case 0: // Name
			switch msg.Type {
			case tea.KeyRunes:
				m.input += msg.String()
			case tea.KeyBackspace:
				if len(m.input) > 0 {
					m.input = m.input[:len(m.input)-1]
				}
			case tea.KeyEnter:
				m.character.Name = m.input
				m.input = ""
				m.step++
			case tea.KeyCtrlC, tea.KeyEsc:
				m.state = stateMainMenu
				m.cursor = 0
			}
		case 1: // Race
			switch msg.Type {
			case tea.KeyRunes:
				m.input += msg.String()
			case tea.KeyBackspace:
				if len(m.input) > 0 {
					m.input = m.input[:len(m.input)-1]
				}
			case tea.KeyEnter:
				m.character.Race = m.input
				m.input = ""
				m.step++
			case tea.KeyCtrlC, tea.KeyEsc:
				m.state = stateMainMenu
				m.cursor = 0
			}
		case 2: // Class
			switch msg.Type {
			case tea.KeyRunes:
				m.input += msg.String()
			case tea.KeyBackspace:
				if len(m.input) > 0 {
					m.input = m.input[:len(m.input)-1]
				}
			case tea.KeyEnter:
				m.character.Class = m.input
				m.input = ""
				m.step++
			case tea.KeyCtrlC, tea.KeyEsc:
				m.state = stateMainMenu
				m.cursor = 0
			}
		case 3: // Ability Scores
			abilities := []string{"Strength", "Dexterity", "Constitution", "Intelligence", "Wisdom", "Charisma"}
			if m.cursor >= len(abilities) {
				// All abilities have been entered
				m.step++
				m.cursor = 0
				break
			}
			switch msg.Type {
			case tea.KeyRunes:
				m.input += msg.String()
			case tea.KeyBackspace:
				if len(m.input) > 0 {
					m.input = m.input[:len(m.input)-1]
				}
			case tea.KeyEnter:
				score := parseInt(m.input)
				if score >= 1 && score <= 20 {
					m.character.AbilityScores[abilities[m.cursor]] = score
					m.cursor++
					m.input = ""
				} else {
					m.err = fmt.Errorf("Ability score must be between 1 and 20")
				}
			case tea.KeyCtrlC, tea.KeyEsc:
				m.state = stateMainMenu
				m.cursor = 0
			}
		case 4: // Confirmation
			switch msg.String() {
			case "y", "Y":
				// Save character
				err := m.saveCharacter()
				if err != nil {
					m.err = err
				} else {
					m.state = stateMainMenu
					m.cursor = 0
				}
			case "n", "N":
				// Cancel and return to main menu
				m.state = stateMainMenu
				m.cursor = 0
			}
		}
	}
	return m, nil
}

// Create Character View
func (m Model) viewCreateCharacter() string {
	switch m.step {
	case 0:
		return fmt.Sprintf("Enter character name:\n%s\n\n(Press Enter to confirm)", m.input)
	case 1:
		return fmt.Sprintf("Enter character race:\n%s\n\n(Press Enter to confirm)", m.input)
	case 2:
		return fmt.Sprintf("Enter character class:\n%s\n\n(Press Enter to confirm)", m.input)
	case 3:
		abilities := []string{"Strength", "Dexterity", "Constitution", "Intelligence", "Wisdom", "Charisma"}
		if m.cursor >= len(abilities) {
			return "All ability scores entered. Press any key to continue."
		}
		prompt := fmt.Sprintf("Enter %s score (1-20):\n%s\n\n(Press Enter to confirm)", abilities[m.cursor], m.input)
		if m.err != nil {
			prompt += fmt.Sprintf("\n\nError: %s", m.err)
			m.err = nil
		}
		return prompt
	case 4:
		charData, _ := json.MarshalIndent(m.character, "", "  ")
		return fmt.Sprintf("Review your character:\n\n%s\n\nSave character? (y/n)", string(charData))
	default:
		return "Character creation complete!"
	}
}

// View Character Update
func (m Model) updateViewCharacter(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "b":
			// Back to list
			m.state = stateListCharacters
			m.cursor = 0
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

// View Character View
func (m Model) viewViewCharacter() string {
	char := m.characters[m.selectedCharIdx]
	charData, _ := json.MarshalIndent(char, "", "  ")
	return fmt.Sprintf("Character Details:\n\n%s\n\nPress 'b' to go back.", string(charData))
}

// Helper function to load characters from the chars/ directory
func (m *Model) loadCharacters() {
	m.characters = []Character{}
	files, err := ioutil.ReadDir("chars/")
	if err != nil {
		m.err = err
		return
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".json") {
			data, err := ioutil.ReadFile("chars/" + file.Name())
			if err != nil {
				m.err = err
				return
			}
			var char Character
			if err := json.Unmarshal(data, &char); err != nil {
				m.err = err
				return
			}
			m.characters = append(m.characters, char)
		}
	}
}

// Helper function to save character to a file
func (m *Model) saveCharacter() error {
	data, err := json.MarshalIndent(m.character, "", "  ")
	if err != nil {
		return err
	}
	if _, err := os.Stat("chars/"); os.IsNotExist(err) {
		err = os.Mkdir("chars/", 0755)
		if err != nil {
			return err
		}
	}
	filename := fmt.Sprintf("chars/%s.json", m.character.Name)
	return ioutil.WriteFile(filename, data, 0644)
}

// Helper function to parse integer from string
func parseInt(s string) int {
	var i int
	fmt.Sscanf(s, "%d", &i)
	return i
}

func main() {
	p := tea.NewProgram(InitialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}


