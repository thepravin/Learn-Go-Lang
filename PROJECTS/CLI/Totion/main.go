package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/thepravin/totion/models"
)

var (
	rootDir     string
	cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
)

// this is go's by default init() method which call first when program start
func init() {
	homDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error getting hot directory", err)
	}
	rootDir = fmt.Sprintf("%s/.totion", homDir)
}

func initializeModel() models.Model {
	// create a folder for storing file at the time of initialization
	err := os.MkdirAll(rootDir, 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	// Initialize new file input
	ti := textinput.New()
	ti.Placeholder = "What would you like to call it?"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 40
	ti.Cursor.Style = cursorStyle
	ti.PromptStyle = cursorStyle
	ti.TextStyle = cursorStyle

	// Initialize note textarea
	ta := textarea.New()
	ta.Placeholder = "Once upon a time..."
	ta.Cursor.Style = cursorStyle
	ta.ShowLineNumbers = false
	ta.Focus()

	// list
	noteList := models.ListFiles(rootDir)

	finalList := list.New(noteList, list.NewDefaultDelegate(), 0, 0)
	finalList.Title = "All notes 📋"
	finalList.Styles.Title = lipgloss.NewStyle().
		Foreground(lipgloss.Color("16")).
		Background(lipgloss.Color("254")).
		Padding(0, 1)

	return models.ModelInitializationBridge(ti, false, rootDir, ta, finalList)
}

func main() {
	p := tea.NewProgram(initializeModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error : ", err.Error())
		os.Exit(1)
	}
}
