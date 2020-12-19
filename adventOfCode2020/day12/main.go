package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	commands := parseInput("./input.txt")

	manhattanDistance := solve1(commands)
	manhattanDistance2 := solve2(commands)

	fmt.Println(manhattanDistance)
	fmt.Println(manhattanDistance2)

}

type CommandType byte

const (
	North          CommandType = 'N'
	South          CommandType = 'S'
	East           CommandType = 'E'
	West           CommandType = 'W'
	Left           CommandType = 'L'
	Rigth          CommandType = 'R'
	Forward        CommandType = 'F'
	DirectionNorth             = 0
	DirectionEast              = 1
	DirectionSouth             = 2
	DirectionWest              = 3
)

type Command struct {
	Type  CommandType
	Value int
}

func solve1(commands []Command) int {

	direction := DirectionEast
	x := 0
	y := 0
	for _, command := range commands {

		switch command.Type {
		case North:
			y += command.Value
		case East:
			x += command.Value
		case South:
			y -= command.Value
		case West:
			x -= command.Value
		case Left:
			direction = calculateDirection(true, command.Value, direction)
		case Rigth:
			direction = calculateDirection(false, command.Value, direction)
		case Forward:

			switch direction {
			case DirectionNorth:
				y += command.Value
			case DirectionEast:
				x += command.Value
			case DirectionSouth:
				y -= command.Value
			case DirectionWest:
				x -= command.Value
			}

		}

		// fmt.Println(x, y, direction)

	}

	return abs(x) + abs(y)

}

func calculateDirection(isLeft bool, degree int, currentDirection int) int {

	if isLeft {
		degree = -degree
	}

	newDirection := (4 + currentDirection + degree/90) % 4

	return newDirection

}

func solve2(commands []Command) int {

	x := 0
	y := 0
	wx := 10
	wy := 1
	for _, command := range commands {

		switch command.Type {
		case North:
			wy += command.Value
		case East:
			wx += command.Value
		case South:
			wy -= command.Value
		case West:
			wx -= command.Value
		case Left:
			wx, wy = calculateDirection2(true, command.Value, wx, wy)
		case Rigth:
			wx, wy = calculateDirection2(false, command.Value, wx, wy)
		case Forward:
			x = x + command.Value*wx
			y = y + command.Value*wy
		}

		// fmt.Println(string(command.Type), command.Value)
		// fmt.Println(x, y, wx, wy)
	}

	return abs(x) + abs(y)

}

func calculateDirection2(isLeft bool, degree int, waypointX int, waypointY int) (int, int) {

	degree = degree / 90

	if isLeft {
		degree = (4 - degree) % 4
	}

	switch degree {
	case 1:
		return waypointY, -waypointX
	case 2:
		return -waypointX, -waypointY
	case 3:
		return -waypointY, waypointX
	default:
		panic("no specified in input")
	}

}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func parseInput(filename string) []Command {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	commandsString := strings.Split(string(content), "\n")

	commands := make([]Command, len(commandsString))

	for i, command := range commandsString {

		c := command[0]
		numb, err := strconv.Atoi(command[1:])
		if err != nil {
			panic(err)
		}

		commands[i] = Command{
			Type:  CommandType(c),
			Value: numb,
		}
	}

	return commands

}
