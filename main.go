package main

import (
	"bufio"
	"inmemorydb/internal/commandline_handler"
	"inmemorydb/internal/database_manager"
	"os"
)

func main() {

	defaultDB := database.NewInMemoryDB("default")

	dbManager := database_manager.NewInMemoryDatabaseManager(defaultDB)

	commandFactory := commandline_handler.NewCommandRecognizer(dbManager)

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		command := commandFactory.RecognizeCommand(scanner.Text())
		command.Execute()


	}

}
