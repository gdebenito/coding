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

	fmt.Println(commands)

	machine := &m.Machine{
		Accumulator:      0,
		Position:         0,
		VisitedPositions: make(map[int]bool),
	}

	for {

		if machine.Position >= len(commands) || machine.Position < 0 {
			fmt.Println("Invalid command jump to " + strconv.Itoa(machine.Position))
			break
		}

		if _, inMap := machine.VisitedPositions[machine.Position]; inMap {
			fmt.Println("Command already executed")
			break
		}

		// Register command is being executed
		machine.VisitedPositions[machine.Position] = true

		command := commands[machine.Position]

		machine.Execute(command)
	}

	fmt.Printf("Value of accumulator: %d\nLast Position: %d\n", machine.Accumulator, machine.Position)

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
