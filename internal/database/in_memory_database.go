package database

import (
	"errors"
	"inmemorydb/pkg/value"
)

type InMemoryDB struct {
	keyValueStorage map[string]value.Value
	name            string
}

func (d *InMemoryDB) SetKeyValue(key string, value value.Value) error {
	d.keyValueStorage[key] = value
	return nil
}
func (d *InMemoryDB) GetValueByKey(key string) (value.Value, error) {
	val, keyExist := d.keyValueStorage[key]
	if keyExist {
		return val, nil
	}
	return nil, errors.New("key not exist")
}
func (d *InMemoryDB) DeleteKey(key string) error {
	_, keyExist := d.keyValueStorage[key]
	if keyExist {
		delete(d.keyValueStorage, key)
		return nil
	}
	return errors.New("key not exist")
}

func (d *InMemoryDB) GetCopyFromStorage() map[string]value.Value {
	return d.keyValueStorage
}

func NewInMemoryDB(name string) *InMemoryDB {
	return &InMemoryDB{
		keyValueStorage: make(map[string]value.Value),
		name:            name,
	}
}
