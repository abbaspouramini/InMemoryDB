package database_manager

import "inmemorydb/pkg/value"

type DatabaseManager interface {
	UseDatabase(dbName string)
	DatabasesList() []string
	ExportDatabase(dbName string, path string) error
	ImportDatabase(path string, dbName string) error
	SetValueForKey(key string, value value.Value) error
	GetValueByKey(key string) (value.Value, error)
	DeleteKeyFromDB(key string) error
	SearchKeys(regex string) ([]string, error)
}
