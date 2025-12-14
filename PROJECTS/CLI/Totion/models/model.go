package models

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	MSG string
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// It is a key press?
	case tea.KeyMsg:

		// what was the actual key press
		switch msg.String() {

		// These keys should exit the program
		case "ctrl+c", "q":
			return m, tea.Quit

		}

	}

	return m, nil
}

func (m Model) View() string {
	return m.MSG
}

func NewMessage(initialMsg string) Model {
	return Model{
		MSG: initialMsg,
	}
}
