package cache

import "time"

type Cache struct {
	Key, Value string
	KVPair     map[string]string
}

func NewCache(k string, v string) Cache {
	kvp := make(map[string]string)
	return Cache{Key: k, Value: v, KVPair: kvp}
}

func (obj Cache) Get(key string) (string, bool) {
	val, ok := obj.KVPair[key]
	if ok == true {
		return val, true
	} else {
		return "", false
	}
}

func (obj Cache) Put(key, value string) {
	obj.KVPair[key] = value
}

func (obj Cache) Keys() []string {
	cacheKeys := make([]string, len(obj.KVPair))
	for k := range obj.KVPair {
		cacheKeys = append(cacheKeys, k)
	}
	return cacheKeys
}

func (obj Cache) PutTill(key, value string, deadline time.Time) {
}
