package command

import (
	"fmt"
	"inmemorydb/pkg/database_manager"
)

type ListCommand struct {
}

func (l *ListCommand) Execute(manager database_manager.DatabaseManager) {
	fmt.Printf("%+v\n", manager.DatabasesList())
}

func NewListCommand() *ListCommand {
	return &ListCommand{}
}
