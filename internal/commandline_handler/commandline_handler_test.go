package commandline_handler

import (
	"github.com/stretchr/testify/assert"
	"inmemorydb/pkg/value"
	"reflect"
	"testing"
)

func TestRecognizeAndCastFunction(t *testing.T) {
	t.Run("type of output must be int64", func(t *testing.T) {
		val := recognizeAndCastTypeOfValue("56")
		assert.Equal(t, reflect.TypeOf(int64(56)), reflect.TypeOf(val))
	})

	t.Run("type of output must be float64", func(t *testing.T) {
		val := recognizeAndCastTypeOfValue("56.65")
		assert.Equal(t, reflect.TypeOf(56.64), reflect.TypeOf(val))
	})

	t.Run("type of output must be string", func(t *testing.T) {
		val := recognizeAndCastTypeOfValue("\"text\"")
		assert.Equal(t, reflect.TypeOf("text"), reflect.TypeOf(val))
	})

	t.Run("type of output must be slice", func(t *testing.T) {
		val := recognizeAndCastTypeOfValue("[56, \"text\",56.659]")
		assert.Equal(t, reflect.TypeOf([]value.Value{}), reflect.TypeOf(val))
	})
}
