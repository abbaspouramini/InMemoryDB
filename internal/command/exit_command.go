package command

import (
	"inmemorydb/pkg/database_manager"
	"os"
)

type ExitCommand struct {
}

func (e *ExitCommand) Execute(manager database_manager.DatabaseManager) {
	os.Exit(0)
}

func NewExitCommand() *ExitCommand {
	return &ExitCommand{}
}
