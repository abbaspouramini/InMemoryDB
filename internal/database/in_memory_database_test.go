package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInMemoryDatabase(t *testing.T) {
	t.Run("set city_tehran_lat 35.65 and get city_tehran_lat ", func(t *testing.T) {
		db := NewInMemoryDB("testDB")
		db.SetKeyValue("city_tehran_lat", 35.65)
		val, _ := db.GetValueByKey("city_tehran_lat")
		assert.Equal(t, 35.65, val)
	})

	t.Run("get test and should receive 'key not exist' error  ", func(t *testing.T) {
		db := NewInMemoryDB("testDB")
		val, err := db.GetValueByKey("test")
		assert.Equal(t, "key not exist", err.Error())
		assert.Equal(t, nil, val)
	})

	t.Run("set a key and delete then get key and should receive 'key not exist' error", func(t *testing.T) {
		db := NewInMemoryDB("testDB")
		db.SetKeyValue("test", 35.65)
		db.DeleteKey("test")
		val, err := db.GetValueByKey("test")
		assert.Equal(t, "key not exist", err.Error())
		assert.Equal(t, nil, val)
	})

}
