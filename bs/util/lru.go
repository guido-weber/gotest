package util

import "container/list"

// LruCache is a Cache
type LruCache struct {
	/* l: list.Elements in LRU order
	   m: key -> *list.Element
	*/
	max int
	l   *list.List
	m   map[string]*list.Element
}

// NewLruCache returns a new LruCache with maximum capacity max
func NewLruCache(max int) *LruCache {
	return &LruCache{max: max, l: list.New(), m: make(map[string]*list.Element, max)}
}

// Get returns the value for key, true or nil, false if not found. If found,
// marks the key as least recently used.
func (c *LruCache) Get(key string) (string, bool) {
	if elem, ok := c.m[key]; ok {
		c.l.MoveToFront(elem)
		return elem.Value.(string), true
	}
	return "", false
}

// Put inserts a value, possibly dropping another element
func (c *LruCache) Put(key, value string) {
	if elem, ok := c.m[key]; ok {
		elem.Value = value
		c.l.MoveToFront(elem)
	} else {
		if c.l.Len() >= c.max {
			last := c.l.Back()
			c.l.Remove(last)
			// TODO: remove from map
		}
		newElem := c.l.PushFront(value)
		c.m[key] = newElem
	}
}
