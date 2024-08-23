package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table      table.Model
	counter    int
	input      textinput.Model
	conditions []string
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.input.Focused() {
			switch msg.String() {
			case "enter":
				m.handleInput()
				m.input.Blur()
				m.input.SetValue("")
			case "esc":
				m.input.Blur()
				m.input.SetValue("")
			default:
				m.input, cmd = m.input.Update(msg)
				cmds = append(cmds, cmd)
			}
		} else {
			switch msg.String() {
			case "n":
				selectedIndex := m.table.Cursor()

				if selectedIndex >= len(m.table.Rows())-1 {
					m.counter++
					m.table.GotoTop()
				} else {
					m.table.MoveDown(1)
				}
			case "p":
				selectedIndex := m.table.Cursor()

				if selectedIndex == 0 {
					m.counter--
					m.table.GotoBottom()
				} else {
					m.table.MoveUp(1)
				}
			case "ctrl+c":
				return m, tea.Quit
			case "enter":
				if !m.input.Focused() {
					m.input.Focus()
				}
			}
		}
	}
	m.table, cmd = m.table.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m *model) handleInput() {
	if len(m.input.Value()) == 0 {
		return
	}
	healOrDamage, condition := strconv.Atoi(m.input.Value())
	if condition != nil {
		m.applyCondition(m.input.Value())
	}
	m.applyHealingOrDamage(healOrDamage)
}

func (m *model) applyCondition(condition string) {
	selectedIndex := m.table.Cursor()
	currentRow := m.table.Rows()[selectedIndex]

	if strings.Contains(currentRow[4], condition) {
		currentRow[4] = strings.Replace(currentRow[4], condition, "", -1)
		m.table.SetRows(m.table.Rows())
	} else {

		currentRow[4] = currentRow[4] + " " + condition
		m.table.SetRows(m.table.Rows())
	}
}

func RemveAtIndex(s []table.Row, index int) []table.Row {
	return append(s[:index], s[index+1:]...)
}

func (m *model) applyHealingOrDamage(healOrDamage int) {
	selectedIndex := m.table.Cursor()
	currentRow := m.table.Rows()[selectedIndex]

	currentHP, _ := strconv.Atoi(currentRow[3])
	newHP := currentHP + healOrDamage

	if newHP <= 0 && currentRow[5] != "y" {
		m.table.SetRows(RemveAtIndex(m.table.Rows(), selectedIndex))
	}

	// Update the row with the new HP
	currentRow[3] = strconv.Itoa(newHP)

	m.table.SetRows(m.table.Rows())
}

func (m model) View() string {
	tableView := baseStyle.Render(m.table.View())
	counterView := fmt.Sprintf("\n\nRound: %d", m.counter)

	var inputView string
	if m.input.Focused() {
		inputView = "\n" + m.input.View()
	}

	return tableView + counterView + inputView
}

func main() {
	columns := []table.Column{
		{Title: "Initiative", Width: 10},
		{Title: "Name", Width: 10},
		{Title: "AC", Width: 4},
		{Title: "HP", Width: 4},
		{Title: "Conditions", Width: 20},
		{Title: "Player", Width: 4},
	}

	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
	defer file.Close()

	// Read the CSV file
	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %v", err)
		os.Exit(1)
	}

	sort.Slice(rows, func(i, j int) bool {
		a, _ := strconv.Atoi(rows[i][0])
		b, _ := strconv.Atoi(rows[j][0])
		return a > b
	})

	tableRows := make([]table.Row, len(rows))
	for i, row := range rows {
		tableRows[i] = table.Row(row)
	}

	var m model

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(tableRows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	ta := textinput.New()
	ta.Placeholder = "Enter the amount for damage or healing..."
	ta.Blur()
	ta.CharLimit = 6
	ta.Width = 60

	m.table = t
	m.counter = 0
	m.input = ta
	m.conditions = []string{}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
