package util

import "testing"

func TestEmpty(t *testing.T) {
	var c = NewLruCache(0)
	var _, ok = c.Get("xxx")
	if ok {
		t.Error("Found value in empty cache")
	}
}

func TestFind(t *testing.T) {
	var c = NewLruCache(1)
	c.Put("xxx", "yyy")
	var val, ok = c.Get("xxx")
	if !ok {
		t.Log("Value not found")
		t.FailNow()
	}
	if val != "yyy" {
		t.Errorf("Expected 'yyy', got '%s'", val)
	}
}
