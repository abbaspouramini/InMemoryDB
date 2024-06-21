package commandline_handler

import (
	"inmemorydb/internal/command"
	"inmemorydb/pkg/value"
	"regexp"
	"strconv"
	"strings"
)

type CommandRecognizer struct {
}



func (c *CommandRecognizer) RecognizeCommand(input string) command.Command {
	partCount := len(strings.Fields(input))
	if partCount == 0 {
		return command.NewNilCommand()
	}
	commandName := strings.Fields(input)[0]

	switch commandName {
	case "set":
		if partCount >= 3 {
			inputParams := input[len(commandName)+1:]
			keyName := strings.Fields(inputParams)[0]
			values := inputParams[len(keyName)+1:]
			val := recognizeAndCastTypeOfValue(values)
			return command.NewSetCommand(keyName, val)
		}

	case "get":
		if partCount == 2 {
			inputParams := input[len(commandName)+1:]
			return command.NewGetCommand(inputParams)
		}

	case "del":
		if partCount == 2 {
			inputParams := input[len(commandName)+1:]
			return command.NewDeleteCommand(inputParams)
		}
	case "keys":
		if partCount == 2 {
			inputParams := input[len(commandName)+1:]
			return command.NewKeysCommand(inputParams)
		}
	case "use":
		if partCount == 2 {
			inputParams := input[len(commandName)+1:]
			return command.NewUseCommand(inputParams)
		}
	case "load":

		if partCount == 3 {
			inputParams := input[len(commandName)+1:]
			input := strings.Fields(inputParams)
			return command.NewLoadCommand(input[0], input[1])
		}
	case "dump":
		if partCount == 3 {
			inputParams := input[len(commandName)+1:]
			input := strings.Fields(inputParams)
			return command.NewDumpCommand(input[0], input[1])
		}
	case "list":
		return command.NewListCommand()
	case "exit":
		return command.NewExitCommand()
	default:
		return command.NewNilCommand()
	}
	return command.NewNilCommand()
}

func recognizeAndCastTypeOfValue(val string) value.Value {
	if val[0] == '"' {
		return val
	} else if val[0] == '[' {
		val = val[1 : len(val)-1]
		values := strings.Split(val, ", ")
		list := make([]value.Value, 0)
		for _, v := range values {
			list = append(list, recognizeAndCastTypeOfValue(v))
		}
		return list
	} else {
		r, _ := regexp.Compile("\\..")
		if r.MatchString(val) {
			num, _ := strconv.ParseFloat(val, 64)
			return num
		} else {
			num, _ := strconv.ParseInt(val, 10, 64)
			return num
		}
	}
}

func NewCommandRecognizer() *CommandRecognizer {
	return &CommandRecognizer{}
}
