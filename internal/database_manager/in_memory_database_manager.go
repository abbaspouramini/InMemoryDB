package database_manager

import (
	"encoding/json"
	"errors"
	"inmemorydb/internal/database"
	"inmemorydb/pkg/repository"
	"inmemorydb/pkg/value"
	"log"
	"os"
	"regexp"
)

type InMemoryDatabaseManager struct {
	databases        map[string]repository.DatabaseRepository
	selectedDatabase string
}

func (i *InMemoryDatabaseManager) UseDatabase(dbName string) {
	_, exist := i.databases[dbName]
	i.selectedDatabase = dbName
	if !exist {
		db := database.NewInMemoryDB(dbName)
		i.databases[dbName] = db
	}

}

func (i *InMemoryDatabaseManager) DatabasesList() []string {
	keys := make([]string, 0, len(i.databases))

	for key := range i.databases {
		keys = append(keys, key)
	}
	return keys
}

func (i *InMemoryDatabaseManager) ExportDatabase(dbName string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		log.Print(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	db, exist := i.databases[dbName]
	if !exist {
		return errors.New("database not exist")
	}
	storage := db.GetCopyFromStorage()
	if err := encoder.Encode(storage); err != nil {
		return err
	}

	return nil
}

func (i *InMemoryDatabaseManager) ImportDatabase(path string, dbName string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	importedDB := make(map[string]value.Value)

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&importedDB); err != nil {
		return err
	}

	i.UseDatabase(dbName)
	for key, val := range importedDB {
		i.databases[i.selectedDatabase].SetKeyValue(key, val)
	}

	return nil
}

func (i *InMemoryDatabaseManager) SetValueForKey(key string, value value.Value) error {
	err := i.databases[i.selectedDatabase].SetKeyValue(key, value)
	if err != nil {
		return err
	}
	return nil
}

func (i *InMemoryDatabaseManager) GetValueByKey(key string) (value.Value, error) {
	val, err := i.databases[i.selectedDatabase].GetValueByKey(key)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (i *InMemoryDatabaseManager) DeleteKeyFromDB(key string) error {
	err := i.databases[i.selectedDatabase].DeleteKey(key)
	if err != nil {
		return err
	}
	return nil
}

func (i *InMemoryDatabaseManager) SearchKeys(regex string) ([]string, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return nil, err
	}
	matchingItems := make([]string, 0)
	storage := i.databases[i.selectedDatabase].GetCopyFromStorage()
	for key := range storage {
		if r.MatchString(key) {
			matchingItems = append(matchingItems, key)
		}
	}
	return matchingItems, nil
}

func NewInMemoryDatabaseManager(defaultDB repository.DatabaseRepository) *InMemoryDatabaseManager {
	
	dbManager := InMemoryDatabaseManager{
		databases:        make(map[string]repository.DatabaseRepository),
		selectedDatabase: "default",
	}
	dbManager.databases["default"] = defaultDB
	return &dbManager
}

