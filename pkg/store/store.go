package store

import (
	"sync"

	"github.com/edanko/gen"
)

func New() *Store {
	return &Store{
		m: make(map[string]*gen.PartData),
	}
}

type Store struct {
	mx sync.RWMutex
	m  map[string]*gen.PartData
}

func (c *Store) Keys() []string {
	c.mx.RLock()
	defer c.mx.RUnlock()

	var res []string

	for k := range c.m {
		res = append(res, k)
	}

	return res
}

func (c *Store) Load(key string) (*gen.PartData, bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	val, ok := c.m[key]
	return val, ok
}

func (c *Store) Store(key string, value *gen.PartData) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key] = value
}

func (c *Store) Inc(key string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key].Quantity++
}
