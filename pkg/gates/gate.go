package gates

import (
	"custompro98/circuits/pkg/core"
	"fmt"
)

type GateOutput core.Bit
type GateInput GateOutput
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

type Gate struct {
	Type GateType

	InputA *GateOutput
	InputB *GateOutput
}

func (g *Gate) Output() []GateOutput {
	o := []GateOutput{}
	switch g.Type {
	case Input:
		o = append(o, GateOutput(core.BitOn))
		break
	case Nand:
		o = append(o, GateOutput(nand(core.Bit(*g.InputA), core.Bit(*g.InputB))))
		break
	case And:
		o = append(o, GateOutput(and(core.Bit(*g.InputA), core.Bit(*g.InputB))))
		break
	case Not:
		o = append(o, GateOutput(not(core.Bit(*g.InputA))))
		break
	case Or:
		o = append(o, GateOutput(or(core.Bit(*g.InputA), core.Bit(*g.InputB))))
		break
	case Nor:
		o = append(o, GateOutput(nor(core.Bit(*g.InputA), core.Bit(*g.InputB))))
		break
	case Xor:
		o = append(o, GateOutput(xor(core.Bit(*g.InputA), core.Bit(*g.InputB))))
		break
	case Xnor:
		o = append(o, GateOutput(xnor(core.Bit(*g.InputA), core.Bit(*g.InputB))))
		break
	case Output:
		o = append(o, GateOutput(*g.InputA))
		break
	}

	return o
}
