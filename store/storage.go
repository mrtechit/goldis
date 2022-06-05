package store

import "sync"

// StorageMap maps a string key to a string value
type StorageMap struct {
	sync.RWMutex
	internal map[string]string
}

// NewStorageMap creates a new string map
func NewStorageMap() *StorageMap {
	gm := StorageMap{
		internal: make(map[string]string),
	}
	return &gm
}

// Dump return a new value from the map or nil, if it does not exist
func (sm *StorageMap) Dump(key string) (value string, ok bool) {
	sm.RLock()
	defer sm.RUnlock()
	result, ok := sm.internal[key]
	return result, ok
}

// Set a given string value at given key
func (sm *StorageMap) Set(key string, value string) bool {
	sm.Lock()
	defer sm.Unlock()
	sm.internal[key] = value
	return true
}

// Rename a given key with new key
func (sm *StorageMap) Rename(key string, newKey string) bool {

	if key == newKey {
		return true
	}
	sm.Lock()
	defer sm.Unlock()
	result, ok := sm.internal[key]
	if ok {
		sm.internal[newKey] = result
		delete(sm.internal, key)
		return true
	}
	return false
}
