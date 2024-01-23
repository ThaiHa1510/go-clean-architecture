package cache

import (
	"container/list"
	"fmt"
)
type LRU struct {
	capacity int
	cache map([int] *list.Element)
	list list.List
}

type entry struct {
	key   int
	value struct{}
}

func NewLRU(capacity int) *LRU {
	return &LRU{
		capacity: capacity,
		cache: make(map[int] *list.Element),
		list: list.New()

	}
}

func(l *LRU) Get(key int) struct{} {
	if value, found := l.cache[key]; found {
		l.list.MoveToFront(value)
		return value.Value.(*entry).value
	}
	return struct{}
}

func(l *LRU) Put(key int, value struct{}) {
	if ele, found := l.cache[key]; found {
		oldVal = ele.Value.(*entry).value
		if oldVal == value {
			l.list.MoveToFront(value)
		} else {
			// Delete old cache
			l.list.Remove(ele)
			newVal = l.list.PushFront(&entry{key: key, value: value})
			l.cache[key] = newVal
		} 
		return
	}
	// delete lastest item
	if l.capacity <= len(l.cache) {
		l.list.Remove(l.list.Back())
	}
	newVal = l.list.Put(value)
	l.cache[key] = &newVal
	return 
}

func(l *LRU) Remove(key int) {
	if ele, found := l.cache[key]; found {
		l.list.Remove(ele)
		delete(l.cache, key)
	}
}