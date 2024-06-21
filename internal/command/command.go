package command

import "inmemorydb/pkg/database_manager"

type Command interface {
	Execute(manager database_manager.DatabaseManager)
}
