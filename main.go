package main

import (
	"bufio"
	"inmemorydb/internal/commandline_handler"
	"inmemorydb/internal/database_manager"
	"inmemorydb/internal/database"
	"os"
)

func main() {

	defaultDB := database.NewInMemoryDB("default")

	dbManager := database_manager.NewInMemoryDatabaseManager(defaultDB)

	commandFactory := commandline_handler.NewCommandRecognizer()

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		command := commandFactory.RecognizeCommand(scanner.Text())
		command.Execute(dbManager)


	}

}
