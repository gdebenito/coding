package command

import (
	"strconv"
	"strings"
)

var Operations = map[string]bool{
	"nop": true,
	"jmp": true,
	"acc": true,
}

type Command struct {
	Operation string
	Value     int
}

func ParseLineToCommand(s string) Command {
	operationAndValue := strings.Split(s, " ")

	operation := operationAndValue[0]
	value, err := strconv.Atoi(operationAndValue[1])

	if err != nil {
		panic(err)
	}

	if !Operations[operation] {
		panic("Operation: " + operation + " does not exist")
	}

	return Command{
		Operation: operation,
		Value:     value,
	}

}
