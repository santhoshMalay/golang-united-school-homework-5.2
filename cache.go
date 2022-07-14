package cache

import (
	"time"
)

type elements struct {
	value   string
	expires time.Time
}

type Cache struct {
	mapCache map[string]elements
}

func NewCache() Cache {
	return Cache{mapCache: make(map[string]elements)}
}

func (obj Cache) Get(key string) (string, bool) {
	element, ok := obj.mapCache[key]
	if !ok {
		return "", false
	}
	return element.value, true
}

func (obj Cache) Put(key, value string) {
	obj.mapCache[key] = elements{
		expires: time.Time{},
		value:   value,
	}
}

func (obj Cache) Keys() []string {
	var keys []string
	for key := range obj.mapCache {
		keys = append(keys, key)
	}
	return keys
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.mapCache[key] = elements{
		expires: deadline,
		value:   value,
	}
}
