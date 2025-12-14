package models

import tea "github.com/charmbracelet/bubbletea"

type MESSAGE struct {
	MSG string
}

func (m MESSAGE) Init() tea.Cmd {
	return nil
}

func (m MESSAGE) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil
}

func NewMessage(initialMsg string) MESSAGE {
	return MESSAGE{
		MSG: initialMsg,
	}
}
