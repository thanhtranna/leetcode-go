package template

import "container/list"

// LFUCache define
type LFUCache struct {
	nodes    map[int]*list.Element
	lists    map[int]*list.List
	capacity int
	min      int
}

type node struct {
	key       int
	value     int
	frequency int
}

// Constructor define
func Constructor(capacity int) LFUCache {
	return LFUCache{nodes: make(map[int]*list.Element),
		lists:    make(map[int]*list.List),
		capacity: capacity,
		min:      0,
	}
}

// Get define
func (lfuCache *LFUCache) Get(key int) int {
	value, ok := lfuCache.nodes[key]
	if !ok {
		return -1
	}
	currentNode := value.Value.(*node)
	lfuCache.lists[currentNode.frequency].Remove(value)
	currentNode.frequency++
	if _, ok := lfuCache.lists[currentNode.frequency]; !ok {
		lfuCache.lists[currentNode.frequency] = list.New()
	}
	newList := lfuCache.lists[currentNode.frequency]
	newNode := newList.PushFront(currentNode)
	lfuCache.nodes[key] = newNode
	if currentNode.frequency-1 == lfuCache.min && lfuCache.lists[currentNode.frequency-1].Len() == 0 {
		lfuCache.min++
	}
	return currentNode.value
}

// Put define
func (lfuCache *LFUCache) Put(key int, value int) {
	if lfuCache.capacity == 0 {
		return
	}
	if currentValue, ok := lfuCache.nodes[key]; ok {
		currentNode := currentValue.Value.(*node)
		currentNode.value = value
		lfuCache.Get(key)
		return
	}
	if lfuCache.capacity == len(lfuCache.nodes) {
		currentList := lfuCache.lists[lfuCache.min]
		backNode := currentList.Back()
		delete(lfuCache.nodes, backNode.Value.(*node).key)
		currentList.Remove(backNode)
	}
	lfuCache.min = 1
	currentNode := &node{
		key:       key,
		value:     value,
		frequency: 1,
	}
	if _, ok := lfuCache.lists[1]; !ok {
		lfuCache.lists[1] = list.New()
	}
	newList := lfuCache.lists[1]
	newNode := newList.PushFront(currentNode)
	lfuCache.nodes[key] = newNode
}
