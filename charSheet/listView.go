package main

import "fmt"

func (m model) listView() string {
	s := "Character List:\n\n"
	for i, char := range m.characters {
		cursor := " "
		if m.selectedChar == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, char.Name)
	}
	cursor := " "
	if m.selectedChar == len(m.characters) {
		cursor = ">"
	}
	s += fmt.Sprintf("%s Create New Character\n", cursor)
	s += "\nPress enter to select, q to quit\n"
	return s
}
