package main

import (
	"errors"
	"fmt"
)

// Простой кеш

type Cache[K comparable, V any] interface {
	Get(key K) (V, error)
	Put(key K, value V)
	Size() int
}

var ErrKeyNotFound = errors.New("key not found")

// Compile check
// Проверка на этапе компиляции - удовлетворяет ли тип данному интерфейсу
var _ Cache[any, any] = (*cacheImpl[any, any])(nil)

type cacheImpl[K comparable, V any] struct {
	keyToElement map[K]V
}

func New[K comparable, V any](capacity int) *cacheImpl[K, V] {
	return &cacheImpl[K, V]{
		keyToElement: make(map[K]V, capacity),
	}
}

func (c *cacheImpl[K, V]) Get(key K) (V, error) {
	if v, ok := c.keyToElement[key]; ok {
		return v, nil
	} else {
		return v, ErrKeyNotFound
	}
}

func (c *cacheImpl[K, V]) Put(key K, value V) {
	c.keyToElement[key] = value
}

func (c *cacheImpl[K, V]) Size() int {
	return len(c.keyToElement)
}

func main() {
	// Создаем кеш: строка -> строка
	var cache Cache[string, string] = New[string, string](10)

	// Пишем в кеш
	cache.Put("name", "Go")

	// Читаем данные
	val, err := cache.Get("name")
	if err == nil {
		fmt.Println(val)
	}

	// Проверяем размер
	fmt.Println(cache.Size())

	// Пишем в кеш
	cache.Put("name", "Go lang")

	// Читаем данные
	val, err = cache.Get("name")
	if err == nil {
		fmt.Println(val)
	}

	// Проверяем размер
	fmt.Println(cache.Size())
}
