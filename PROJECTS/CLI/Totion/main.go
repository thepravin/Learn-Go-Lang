package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/thepravin/totion/models"
)

func initializeMode() models.Model {
	return models.NewMessage("I am pravin...")
}

func main() {
	p := tea.NewProgram(initializeMode())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error : ", err.Error())
		os.Exit(1)
	}
}
