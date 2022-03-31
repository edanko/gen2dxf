package store

import (
	"testing"

	"github.com/edanko/gen"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	store := New()
	assert.NotNil(t, store)
	assert.NotNil(t, store.m)

	t.Run("keys test", func(t *testing.T) {
		key := "test"
		value := &gen.PartData{Name: "name"}

		store.Store(key, value)

		got := store.Keys()

		assert.Equal(t, []string{"test"}, got)
	})

	t.Run("store test", func(t *testing.T) {
		key := "test"
		value := &gen.PartData{Name: "name"}

		store.Store(key, value)

		assert.Equal(t, value.Name, store.m[key].Name)
	})

	t.Run("load test", func(t *testing.T) {
		key := "test"
		value := &gen.PartData{Name: "name"}

		store.Store(key, value)

		got, ok := store.Load(key)

		assert.Equal(t, true, ok)
		assert.Equal(t, value, got)
	})

	t.Run("inc test", func(t *testing.T) {
		key := "test"
		value := &gen.PartData{Name: "name"}

		store.Store(key, value)
		store.Inc(key)
		store.Inc(key)

		got, ok := store.Load(key)

		assert.Equal(t, true, ok)
		assert.Equal(t, 2, got.Quantity)
	})

}
