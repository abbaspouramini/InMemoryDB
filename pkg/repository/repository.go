package repository

import "inmemorydb/pkg/value"

type DatabaseRepository interface {
	SetKeyValue(key string, value value.Value) error
	GetValueByKey(key string) (value.Value, error)
	DeleteKey(key string) error
	GetCopyFromStorage() map[string]value.Value
}
