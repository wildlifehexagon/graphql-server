package registry

import "github.com/emirpasic/gods/maps/hashmap"

type Collection struct {
	ClientMap *hashmap.Map
}

type Registry interface {
	AddAPIClient(key string, client interface{})
	GetAPIClient(key string) (interface{}, bool)
}

// NewRegistry generates a new (empty) hashmap
func NewRegistry() (Registry, error) {
	return &Collection{ClientMap: hashmap.New()}, nil
}

// AddAPIClient adds a new entry to the hashmap
func (c *Collection) AddAPIClient(key string, client interface{}) {
	c.ClientMap.Put(key, client)
}

// GetAPIClient looks up a client in the hashmap
func (c *Collection) GetAPIClient(key string) (interface{}, bool) {
	return c.ClientMap.Get(key)
}
