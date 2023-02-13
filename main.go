package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"orthocal/day"
	"orthocal/event"
	"os"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

var cursorStyle = lipgloss.NewStyle().
	Bold(true).
	Background(lipgloss.Color("2E31D4")).
	Foreground(lipgloss.Color("2E55GF"))

type model struct {
	days             [5][7]day.Day
	cursorX, cursorY int
	selected         map[int]struct{}
}

func initialModel() model {
	return model{
		// Our to-do list is a grocery list
		days:    BuildMonthDays(2),
		cursorX: 1,
		cursorY: 1,
		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func BuildMonthDays(month int) [5][7]day.Day {
	counter := 1
	var daysMatrix [5][7]day.Day
	for i := 0; i < 5; i++ {
		for j := 0; j < 7; j++ {
			fmt.Sprintf("%s, %s", i, j)
			daysMatrix[i][j] = day.NewDay(counter, "St. John", "Genesis 1", event.FastCheese)
			counter++
		}
	}
	return daysMatrix
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursorY > 0 {
				m.cursorY--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursorY < 4 {
				m.cursorY++
			}

		case "left", "h":
			if m.cursorX > 0 {
				m.cursorX--
			}
		case "right", "l":
			if m.cursorX < 7 {
				m.cursorX++
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "Mo Tu We Th Fr Sa Su\n\n"

	for i := 0; i < 5; i++ {
		for j := 0; j < 7; j++ {
			if i == m.cursorX && j == m.cursorY {
				shtring := fmt.Sprintf(" %d ", m.days[i][j].GetDay())
				s += cursorStyle.Render(shtring)
			}
			s += fmt.Sprintf(" %d ", m.days[i][j].GetDay())
		}
		s += "\n"

	}

	s += "\nPress q to quit.\n"

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("there's been an error: %v", err)
		os.Exit(1)
	}
}
