package command

import "inmemorydb/pkg/database_manager"

type UseCommand struct {
	dbName string
}

func (u *UseCommand) Execute(manager database_manager.DatabaseManager) {
	manager.UseDatabase(u.dbName)
}

func NewUseCommand(dbName string) *UseCommand {
	return &UseCommand{
		dbName: dbName,
	}
}
