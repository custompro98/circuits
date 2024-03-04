package gates

import (
	"custompro98/circuits/pkg/core"
	"fmt"
	"strings"
)

type GateOutput core.Bit
type GateInput core.Bit
type GateType int

const (
	Input GateType = iota
	Nand
	And
	Not
	Or
	Nor
	Xor
	Xnor
	Output
)

func (gt GateType) String() string {
	switch gt {
	case Input:
		return "Input"
	case Nand:
		return "Nand"
	case And:
		return "And"
	case Not:
		return "Not"
	case Or:
		return "Or"
	case Nor:
		return "Nor"
	case Xor:
		return "Xor"
	case Xnor:
		return "Xnor"
	case Output:
		return "Output"
	}

	return fmt.Sprintf("Could not translate gate %v", int(gt))
}

type Gate interface {
	WithInput(Gate) Gate
	ClearInput() Gate

	Input() []GateInput
	Output() GateOutput
	Type() string
	String() string

	// Only valid if Type == "Input"
	Toggle() Gate
}

type gate struct {
	gateType GateType

	input  []Gate
	output GateOutput
}

func New(t GateType) Gate {
	g := &gate{
		gateType: t,

		input: make([]Gate, 0),

		output: GateOutput(core.BitInvalid),
	}

	if t == Input {
		g.output = GateOutput(core.BitOff)
	}

	return g
}

func (g *gate) WithInput(in Gate) Gate {
	g.input = append(g.input, in)

	return g
}

func (g *gate) ClearInput() Gate {
	g.input = make([]Gate, 0)

	return g
}

func (g *gate) Input() []GateInput {
	if g.gateType == Input || g.gateType == Output {
		return []GateInput{}
	}

	input := make([]GateInput, 0)

	for _, v := range g.input {
		input = append(input, GateInput(v.Output()))
	}

	return input
}

func (g *gate) Output() GateOutput {
	switch g.gateType {
	case Input:
		break
	case Nand:
		g.output = GateOutput(nand(core.Bit(g.input[0].Output()), core.Bit(g.input[1].Output())))
		break
	case And:
		g.output = GateOutput(and(core.Bit(g.input[0].Output()), core.Bit(g.input[1].Output())))
		break
	case Not:
		g.output = GateOutput(not(core.Bit(g.input[0].Output())))
		break
	case Or:
		g.output = GateOutput(or(core.Bit(g.input[0].Output()), core.Bit(g.input[1].Output())))
		break
	case Nor:
		g.output = GateOutput(nor(core.Bit(g.input[0].Output()), core.Bit(g.input[1].Output())))
		break
	case Xor:
		g.output = GateOutput(xor(core.Bit(g.input[0].Output()), core.Bit(g.input[1].Output())))
		break
	case Xnor:
		g.output = GateOutput(xnor(core.Bit(g.input[0].Output()), core.Bit(g.input[1].Output())))
		break
	case Output:
		g.output = GateOutput(g.input[0].Output())
		break
	}

	return g.output
}

func (g *gate) Type() string {
	return g.gateType.String()
}

func (g *gate) Toggle() Gate {
	if g.gateType != Input {
		return g
	}

	g.output = GateOutput(not(core.Bit(g.output)))

	return g
}

func (g *gate) String() string {
	s := []string{fmt.Sprintf("%s\n", g.Type())}

	if g.gateType != Input && g.gateType != Output {
		s = append(s, fmt.Sprintf("Input: %v", g.Input()))
	} else {
		s = append(s, "")
	}

	s = append(s, fmt.Sprintf("Output: %v", g.Output()))

	return strings.Join(s, "\n")
}
