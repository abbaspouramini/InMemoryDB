package command

import "inmemorydb/pkg/database_manager"

type DeleteCommand struct {
	key string
}

func (d *DeleteCommand) Execute(manager database_manager.DatabaseManager) {
	manager.DeleteKeyFromDB(d.key)
}

func NewDeleteCommand(key string) *DeleteCommand {
	return &DeleteCommand{
		key: key,
	}
}
