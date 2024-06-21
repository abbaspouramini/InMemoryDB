package command

import (
	"inmemorydb/pkg/database_manager"
	"log"
)

type LoadCommand struct {
	dbName string
	path   string
}

func (l *LoadCommand) Execute(manager database_manager.DatabaseManager) {
	if err := manager.ImportDatabase(l.path, l.dbName); err != nil {
		log.Print(err)
	}

}

func NewLoadCommand(path string, dbName string) *LoadCommand {
	return &LoadCommand{
		dbName: dbName,
		path:   path,
	}
}
