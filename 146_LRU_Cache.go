// Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

// Implement the LRUCache class:

//     LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
//     int get(int key) Return the value of the key if the key exists, otherwise return -1.
//     void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.

// The functions get and put must each run in O(1) average time complexity.

// Example 1:

// Input
// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
// Output
// [null, null, null, 1, null, -1, null, -1, 3, 4]

// Explanation
// LRUCache lRUCache = new LRUCache(2);
// lRUCache.put(1, 1); // cache is {1=1}
// lRUCache.put(2, 2); // cache is {1=1, 2=2}
// lRUCache.get(1);    // return 1
// lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
// lRUCache.get(2);    // returns -1 (not found)
// lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
// lRUCache.get(1);    // return -1 (not found)
// lRUCache.get(3);    // return 3
// lRUCache.get(4);    // return 4

// Constraints:

//     1 <= capacity <= 3000
//     0 <= key <= 104
//     0 <= value <= 105
//     At most 2 * 105 calls will be made to get and put.

// Кратко:
// Кэш с алгоритмом «наименее недавно использовавшийся» (LRU).
// Алгоритм и структура данных для временного хранения информации.
// Когда память заполняется, он удаляет те элементы, к которым дольше всего не было обращений.

// Усложним задачу - ключи будут строковые, вместо int

package main

import (
	"container/list"
	"errors"
	"iter"
)

var ErrKeyNotFound = errors.New("key not found")

// Интерфейс
type LRUCache[K comparable, V any] interface {
	// Put O(1)
	Put(key K, value V)

	// Get O(1)
	Get(key K) (V, error)

	// Size O(1)
	Size() int

	// type Seq2[K, V any] func(yield func(K, V) bool) - означает функцию-итератор для последовательности пар значений
	// K и V — это дженерик-параметры (типы ключа и значения)
	// yield — это вспомогательная функция обратного вызова (callback), которую итератор вызывает для каждого элемента
	// Если yield возвращает false, итератор понимает, что цикл нужно прервать (например, если сработал break)

	// iter.Seq[V] — перебирает по одному значению за раз (как for v := range slice)
	// iter.Seq2[K, V] — перебирает по два значения за раз (как for k, v := range map)

	// All O(capacity)
	All() iter.Seq2[K, V]
}

var _ LRUCache[any, any] = (*cacheImpl[any, any])(nil)

// Реализация через LinkedList - связанный список
// head - более свежий
// будем удалять элементы ближе к tail

// Сортировка элементов по времени последнего использования
// head k1 k4 k5 k3 k2 tail
// head 1s 2s 3s 5s 10s tail

// Используем LinkedList из go

// Вспомогательная структура для хранения ключа и значения для доступа к map (keyToElement) и к list (linkedList)
type node[K comparable, V any] struct {
	key   K
	value V
}

type cacheImpl[K comparable, V any] struct {
	linkedList   *list.List
	keyToElement map[K]V
	capacity     int
}

func New[K comparable, V any](capacity int) *cacheImpl[K, V] {
	return &cacheImpl[K, V]{
		linkedList:   list.New(),
		keyToElement: make(map[K]*list.Element, capacity),
		capacity:     capacity,
	}
}

// Вспомогательная функция выборки последнего элемента
func (c *cacheImpl[K, V]) extractLatest() {
	del := c.linkedList.Back()
	c.linkedList.Remove(del)

	delete(c.keyToElement, c.getNodeFromElement(del).key)
}

func (c *cacheImpl[K, V]) Put(key K, value V) {
	// Общая идея
	// if c.Size() == c.capacity {
	// 	// delete
	// }

	// Если значение есть
	if link, ok := c.keyToElement[key]; ok {
		n := c.getNodeFromElement(link)
		// Обновляем значение
		n.value = value
		// Двигаем ноду
		c.linkedList.MoveToFront(link)
		return
	}

	// Если новый элемент
	if c.Size() == c.capacity {
		// delete наиболее старого элемента
		c.extractLatest()
	}

	// insert
	n := &node[K, V]{
		key:   key,
		value: value,
	}

	c.keyToElement[key] = c.linkedList.PushFront(n)
}

// Функция получения значения из list
// func (c *cacheImpl[K, V]) getValueFromNode(element *list.Element) V {
// 	switch v := element.Value.(type) {
// 	case V:
// 		return v
// 	default:
// 		panic("unreachable")
// 	}
// }

func (c *cacheImpl[K, V]) getNodeFromElement(element *list.Element) *node[K, V] {
	switch v := element.Value.(type) {
	case *node[K, V]:
		return v
	default:
		panic("unreachable")
	}
}

func (c *cacheImpl[K, V]) Get(key K) (V, error) {
	// Проверяем наличие ноды
	if link, ok := c.keyToElement[key]; ok {
		// Двигаем ноду в начало списка
		c.linkedList.MoveToFront(link)
		n := c.getNodeFromElement(link)
		return n.value, nil
	}

	// Создаем zero-value
	var zero V
	// Вернем ошибку, если нет ноды
	return zero, ErrKeyNotFound

}

func (c *cacheImpl[K, V]) Size() int {
	return len(c.keyToElement)
}

func (c *cacheImpl[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		cur := c.linkedList.Front()

		for range c.Size() {
			n := c.getNodeFromElement(cur)

			if !yield(n.key, n.value) {
				return
			}

			cur = cur.Next()
		}
	}
}
