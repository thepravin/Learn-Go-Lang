package models

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	docStyle = lipgloss.NewStyle().Margin(1, 2)
)

type Model struct {
	newFileInput             textinput.Model
	notetextarea             textarea.Model
	isCreateFileInputVisible bool
	RootDir                  string
	currentFile              *os.File // sotre the file pointer
	list                     list.Model
	isListShowing            bool
}

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	switch msg := msg.(type) {

	// list
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v-18)

	// It is a key press?
	case tea.KeyMsg:

		// what was the actual key press
		switch msg.String() {

		// These keys should exit the program
		case "ctrl+c":
			return m, tea.Quit
		case "ctrl+n":
			// fmt.Println("Key : ", msg)
			m.isCreateFileInputVisible = true
			return m, nil

		case "ctrl+l":
			noteList := ListFiles(m.RootDir)
			m.list.SetItems(noteList)
			m.isListShowing = true
			return m, nil

		case "esc":
			if m.isCreateFileInputVisible {
				m.isCreateFileInputVisible = false
			}
			if m.currentFile != nil {
				m.notetextarea.SetValue("")
				m.currentFile = nil
			}
			if m.isListShowing {
				if m.list.FilterState() == list.Filtering {
					break
				}
				m.isListShowing = false
			}

			return m, nil

		case "ctrl+s":
			// textarea value -> write it in that file and close it

			if m.currentFile == nil { // If NO file is open, do nothing
				break
			}

			if err := m.currentFile.Truncate(0); err != nil {
				fmt.Println("Can not save the file :(")
				return m, nil
			}

			if _, err := m.currentFile.Seek(0, 0); err != nil {
				fmt.Println("Can not save the file :(")
				return m, nil
			}

			// write into file
			if _, err := m.currentFile.WriteString(m.notetextarea.Value()); err != nil {
				fmt.Println("Can not save the file :(")
				return m, nil
			}

			if err := m.currentFile.Close(); err != nil {
				fmt.Println("Can not close the file :(")
				return m, nil
			}

			m.currentFile = nil
			m.notetextarea.SetValue("")

			return m, nil

		case "enter":
			// if file is open
			if m.currentFile != nil {
				break
			}

			// if list open
			if m.isListShowing {
				item, ok := m.list.SelectedItem().(item)
				if ok {
					filePath := fmt.Sprintf("%s/%s", m.RootDir, item.title)

					content, err := os.ReadFile(filePath)
					if err != nil {
						log.Printf("Error reading file : %v", err)
						return m, nil
					}
					m.notetextarea.SetValue(string(content))

					file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
					if err != nil {
						log.Printf("Error reading file : %v", err)
						return m, nil
					}
					m.currentFile = file // assign file pointer
					m.isListShowing = false

				}

				return m, nil

			}

			// todo : create file
			fileName := m.newFileInput.Value()
			if fileName != "" {
				filePath := fmt.Sprintf("%s/%s.md", m.RootDir, fileName)

				// check file already prsent or not
				if _, err := os.Stat(filePath); err == nil {
					// file exist
					return m, nil
				}

				file, err := os.Create(filePath)
				if err != nil {
					log.Fatal("Error : %v", err)
				}
				m.currentFile = file
				m.isCreateFileInputVisible = false
				m.newFileInput.SetValue("")
			}
			return m, nil
		}

	}

	if m.isCreateFileInputVisible {
		m.newFileInput, cmd = m.newFileInput.Update(msg)
	}

	if m.currentFile != nil {
		m.notetextarea, cmd = m.notetextarea.Update(msg)
	}

	if m.isListShowing {
		m.list, cmd = m.list.Update(msg)
	}

	return m, cmd
}

func (m Model) View() string {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("16")).
		Background(lipgloss.Color("205")).
		PaddingLeft(2).
		PaddingRight(2)

	welcome := style.Render("WELCOME TO TOTION 📝")

	help := "Ctrl+N: new file | Ctrl+L: list | Ctrl+S: save | Esc: back | Ctrl+Q: quit"

	view := ""
	if m.isCreateFileInputVisible {
		view = m.newFileInput.View()
	}

	if m.currentFile != nil {
		view = m.notetextarea.View()
	}

	if m.isListShowing {
		view = m.list.View()
	}

	return fmt.Sprintf("\n%s\n\n%s\n\n%s", welcome, view, help)
}

// helper function
func ModelInitializationBridge(fileName textinput.Model, isCreateFileInputVisible bool, rootDir string, notesTextArea textarea.Model, finalList list.Model) Model {
	return Model{
		newFileInput:             fileName,
		isCreateFileInputVisible: isCreateFileInputVisible,
		RootDir:                  rootDir,
		notetextarea:             notesTextArea,
		list:                     finalList,
	}
}

// ListFiles returns a list of files in the given directory
func ListFiles(rootDir string) []list.Item {
	items := make([]list.Item, 0)

	entries, err := os.ReadDir(rootDir)
	if err != nil {
		log.Fatal("Error reading notes")
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				continue
			}

			modifiedTime := info.ModTime().Format("2006-01-02 15:04")

			items = append(items, item{
				title: entry.Name(),
				desc:  fmt.Sprintf("Modified: %s", modifiedTime),
			})
		}
	}

	return items
}
