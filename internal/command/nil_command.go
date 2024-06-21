package command

import "inmemorydb/pkg/database_manager"

type NilCommand struct {
}

func (n *NilCommand) Execute(manager database_manager.DatabaseManager) {
}

func NewNilCommand() *NilCommand {
	return &NilCommand{}
}
