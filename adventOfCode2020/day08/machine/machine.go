package machine

import (
	cmd "main/command"
)

type Machine struct {
	Accumulator      int
	Position         int
	VisitedPositions map[int]bool
}

// Execute ...
// Executes the command given for the machine
// This method updates the position of the next command
// and the accumulator value
func (m *Machine) Execute(c cmd.Command) {

	switch c.Operation {
	case "nop":
		m.Position++
	case "jmp":
		m.Position += c.Value
	case "acc":
		m.Accumulator += c.Value
		m.Position++
	}

	// fmt.Printf("Command: %s Value: %d Accumulator: %d NextPosition: %d\n", c.Operation, c.Value, m.Accumulator, m.Position)

}
