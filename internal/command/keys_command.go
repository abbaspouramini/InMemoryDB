package command

import (
	"fmt"
	"inmemorydb/pkg/database_manager"
	"log"
)

type KeysCommand struct {
	regex string
}

func (k *KeysCommand) Execute(manager database_manager.DatabaseManager) {
	res, err := manager.SearchKeys(k.regex)
	if err != nil {
		log.Print(err)
	}
	fmt.Printf("%+v\n", res)
}

func NewKeysCommand(regex string) *KeysCommand {
	return &KeysCommand{
		regex: regex,
	}
}
