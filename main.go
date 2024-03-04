package main

import (
	"custompro98/circuits/pkg/gates"
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	gloss "github.com/charmbracelet/lipgloss"
)

var (
	gateStyle = gloss.NewStyle().
			Width(15).
			Height(5).
			Align(gloss.Center, gloss.Top).
			BorderStyle(gloss.RoundedBorder())

	focusedGateStyle = gloss.NewStyle().
				Width(15).
				Height(5).
				Align(gloss.Center, gloss.Top).
				BorderStyle(gloss.RoundedBorder()).
				BorderForeground(gloss.Color("21"))

	arrowStyle = gloss.NewStyle().
			Width(5).
			Height(5).
			Align(gloss.Center, gloss.Center).
			BorderStyle(gloss.HiddenBorder())

	helpStyle = gloss.NewStyle().Foreground(gloss.Color("241"))
)

type model struct {
	selected int
	gates    []gates.Gate
}

func initialModel() model {
	in := gates.New(gates.Input)
	nand := gates.New(gates.Nand).WithInput(in).WithInput(in)
	out := gates.New(gates.Output).WithInput(nand)

	return model{
		gates: []gates.Gate{in, nand, out},
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			if m.selected+1 < len(m.gates) {
				m.selected = m.selected + 1
			} else {
				m.selected = 0
			}

			break
		case "t":
			m.gates[m.selected].Toggle()
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	var s string
	var gateStrings []string

	for i, gate := range m.gates {
		if i == m.selected {
			gateStrings = append(gateStrings, focusedGateStyle.Render(gate.String()))
		} else {
			gateStrings = append(gateStrings, gateStyle.Render(gate.String()))
		}

		if i != len(m.gates)-1 {
			// do not append an arrow to the last gate
			gateStrings = append(gateStrings, arrowStyle.Render("-->"))
		}
	}

	s += gloss.JoinHorizontal(gloss.Top, gateStrings...)

	help := []string{"tab: focus next"}

	if m.gates[m.selected].Type() == gates.Input.String() {
		help = append(help, "t: toggle bit state")
	}

	help = append(help, "q: exit")

	s += helpStyle.Render(fmt.Sprintf("\n%s\n", strings.Join(help, " • ")))

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
