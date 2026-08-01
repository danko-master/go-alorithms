// 146 LRU Cache
// Least Recently Used - алгоритм управления паматью и структура данных,
// которая хранит только самые последние использованные данные

// Кэш с вытеснением

// O(1) - get\put
// Двусвязный список, Get или Put перемещают узел в голову (most recent)

// Особенность:
// sentinel - узлы (или фиктивные/dummy-узлы) head/tail (dummy), чтобы не проверять на nil
// Специальные вспомогательные элементы двусвязного списка: фиктивная голова (head) и фиктивный хвост (tail).
// Они не хранят реальных данных кэша, а служат фиксированными границами списка, упрощая логику работы.

// Пример:

// c := New(2)
// // cache: {1 = 1}
// c.Put(1, 1)
// // cache: {2 = 2, 1 = 1}
// c.Put(2, 2)
// // get=1, cache: {1=1, 2=2}
// c.Get(1)
// // вытесним 2, cache: {3 = 3, 1 = 1}
// c.Put(3, 3)
// // get=-1
// c.Get(2)

package main

// Работаем с указателями
// Двусвязный список не массив, а цепочка объектов в памяти

// ссылочная структура prev-next хранит ссылки на узлы.
// m[key] -> *Node

// Если бы хранили map[int]Node (значения), то при pushFront(),
// то пришлось бы копировать всю структуру и обновлять map,
// а нам требуется уложиться в O(1)

type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

type LRUCache struct {
	cap int
	m   map[int]*Node
	h   *Node // sentinel head
	t   *Node // sentinel tail
}

func NewCache(cap int) LRUCache {
	c := LRUCache{
		cap: cap,
		m:   make(map[int]*Node, cap),
		h:   &Node{},
		t:   &Node{},
	}

	c.h.next = c.t
	c.t.prev = c.h

	return c
}

func (c *LRUCache) Get(key int) int {
	n, ok := c.m[key]

	if !ok {
		return -1
	}

	c.remove(n)
	c.pushFront(n)

	// Особенность Go - компилятор сам разыменует его
	// Равносильно - (*n).val
	// Go видит, что n это *Node и подставляет (*n)
	return n.val
}

func (c *LRUCache) Put(key, val int) {
	if n, ok := c.m[key]; ok {
		n.val = val
		c.remove(n)
		c.pushFront(n)
		return
	}

	if len(c.m) == c.cap {
		last := c.t.prev
		c.remove(last)
		delete(c.m, last.key)
	}

	// Берём адрес свежесозданной структуры
	n := &Node{key: key, val: val}
	c.m[key] = n
	c.pushFront(n)
}

func (c *LRUCache) remove(n *Node) {
	// n.next - Go сам разыменует его
	// в данном случае n.next = (*n).next
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (c *LRUCache) pushFront(n *Node) {
	n.next = c.h.next
	n.prev = c.h
	c.h.next.prev = n
	c.h.next = n
}

func main() {
	c := NewCache(2)
	// cache: {1 = 1}
	c.Put(1, 1)
	// cache: {2 = 2, 1 = 1}
	c.Put(2, 2)
	// get=1, cache: {1=1, 2=2}
	c.Get(1)
	// вытесним 2, cache: {3 = 3, 1 = 1}
	c.Put(3, 3)
	// get=-1
	c.Get(2)
}
