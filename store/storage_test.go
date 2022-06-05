package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorageMapSingleClientSetAndDump(t *testing.T) {
	m := NewStorageMap()
	m.Set("foo", "bar")
	m.Set("foo2", "2")
	val, ok := m.Dump("foo")
	assert.Equal(t, ok, true)
	assert.Equal(t, val, "bar")
	val, ok = m.Dump("foo2")
	assert.Equal(t, val, "2")
	_, ok = m.Dump("foo3")
	assert.Equal(t, ok, false)
}

func TestStorageMapSingleClientSetAndRename(t *testing.T) {
	m := NewStorageMap()
	m.Set("foo", "bar")
	val := m.Rename("foo", "newfoo")
	assert.Equal(t, val, true)
	value, ok := m.Dump("newfoo")
	assert.Equal(t, ok, true)
	assert.Equal(t, value, "bar")
}
