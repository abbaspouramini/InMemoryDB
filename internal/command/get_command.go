package command

import (
	"fmt"
	"inmemorydb/pkg/database_manager"
	"log"
)

type GetCommand struct {
	key string
}

func (g *GetCommand) Execute(manager database_manager.DatabaseManager) {
	res, err := manager.GetValueByKey(g.key)
	if err != nil {
		log.Print(err)
	}
	fmt.Printf("%+v\n", res)
}

func NewGetCommand(key string) *GetCommand {
	return &GetCommand{
		key: key,
	}
}
