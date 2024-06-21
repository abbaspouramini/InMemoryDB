package command

import (
	"inmemorydb/pkg/database_manager"
	"log"
)

type DumpCommand struct {
	dbName string
	path   string
}

func (d *DumpCommand) Execute(manager database_manager.DatabaseManager) {
	if err := manager.ExportDatabase(d.dbName, d.path); err != nil {
		log.Print(err)
	}
}

func NewDumpCommand(dbName string, path string) *DumpCommand {
	return &DumpCommand{
		dbName: dbName,
		path:   path,
	}
}
