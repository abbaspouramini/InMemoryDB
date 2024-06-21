package command

import (
	"inmemorydb/pkg/database_manager"
	"inmemorydb/pkg/value"
	"log"
)

type SetCommand struct {
	key   string
	value value.Value
}

func (s *SetCommand) Execute(manager database_manager.DatabaseManager) {
	if err := manager.SetValueForKey(s.key, s.value); err != nil {
		log.Print(err)
	}
}

func NewSetCommand(key string, value value.Value) *SetCommand {
	return &SetCommand{
		key:   key,
		value: value,
	}
}
