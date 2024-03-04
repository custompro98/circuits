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
	out := gates.New(gates.Output).WithInput(in)

	return model{
		gates: []gates.Gate{in, out},
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
		case tea.KeyTab.String():
			if m.selected+1 < len(m.gates) {
				m.selected = m.selected + 1
			} else {
				m.selected = 0
			}

			break
		case tea.KeyShiftTab.String():
			if m.selected-1 >= 0 {
				m.selected = m.selected - 1
			} else {
				m.selected = len(m.gates) - 1
			}

			break
		case "t":
			if m.gates[m.selected].Type() != gates.Input.String() {
				// do not allow toggling a non input
				break
			}

			m.gates[m.selected].Toggle()
			break
		case "i":
			if m.selected == 0 {
				// do not allow a new gate before input
				break
			}

			ng := make([]gates.Gate, 0)

			for i, v := range m.gates {
				if i == m.selected {
					// insert new gate if we're looking at the current selection
					ng = append(
						ng,
						gates.
							New(gates.Nand).
							WithInput(m.gates[i-1]).
							WithInput(m.gates[i-1]),
					)

					// set the current selections input to the new gate
					v.ClearInput().
						WithInput(ng[len(ng)-1]).
						WithInput(ng[len(ng)-1])
				}

				// put old gate in place
				ng = append(ng, v)
			}

			m.gates = ng
			break
		case "d":
			if m.selected == 0 || m.selected == len(m.gates) {
				// do not allow removing the input or output
				break
			}

			ng := make([]gates.Gate, 0)

			for i, v := range m.gates {
				if i == m.selected {
					// remove current selection
					continue
				} else if i == m.selected+1 {
					// set the new current selection's input to previous selection's input
					v.ClearInput().
						WithInput(m.gates[m.selected-1]).
						WithInput(m.gates[m.selected-1])
				}

				// put old gate in place
				ng = append(ng, v)
			}

			m.gates = ng
			break
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

	if m.selected != 0 {
		help = append(help, "i: insert nand")
	}

	if m.selected != 0 && m.selected != len(m.gates)-1 {
		help = append(help, "d: delete selection")
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
