package main

import (
	"fmt"
	"io/ioutil"
	c "main/command"
	m "main/machine"
	"strconv"
	"strings"
)

func main() {

	// commands := getCommandsFromInputFile("./test.txt")
	commands := getCommandsFromInputFile("./input.txt")

	jmpOrNop := make(map[int]bool)

	for i, command := range commands {
		if command.Operation == "nop" || command.Operation == "jmp" {
			jmpOrNop[i] = true
		}
	}

	for k := range jmpOrNop {
		switchCommand(&commands[k])
		hasPassed, v := executeIteration(commands)
		switchCommand(&commands[k])

		if hasPassed {
			fmt.Println("Value: " + strconv.Itoa(v))
			break
		}
	}

}

func switchCommand(command *c.Command) {
	if command.Operation == "jmp" {
		command.Operation = "nop"
	} else if command.Operation == "nop" {
		command.Operation = "jmp"
	}
}

func executeIteration(commands []c.Command) (bool, int) {

	machine := &m.Machine{
		Accumulator:      0,
		Position:         0,
		VisitedPositions: make(map[int]bool),
	}

	for {

		if machine.Position >= len(commands) || machine.Position < 0 {
			fmt.Println("Invalid command jump to " + strconv.Itoa(machine.Position))
			return true, machine.Accumulator
		}

		if _, inMap := machine.VisitedPositions[machine.Position]; inMap {
			// fmt.Println("Command already executed")
			return false, machine.Accumulator
		}

		if machine.Position == len(commands) {
			fmt.Println("finished")
			return true, machine.Accumulator
		}

		// Register command is being executed
		machine.VisitedPositions[machine.Position] = true

		command := commands[machine.Position]

		machine.Execute(command)

	}

}

func getCommandsFromInputFile(filename string) []c.Command {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	instructions := strings.Split(string(content), "\n")

	var commands []c.Command

	for _, instruction := range instructions {
		commands = append(commands, c.ParseLineToCommand(instruction))
	}

	return commands

}
