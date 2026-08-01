// Когда применяется односвязный список:
// Стек (LIFO)
// Очередь с хвостом (FIFO)
// Разделение chaining в хэш-таблицах
// Объединение merge
// Простой список

// 206 leetcode
// Развернуть односвязный список
// "Инвертировать" последовательность
// O(n) time
// O(1) mem

// Алгоритм:
// три указателя - prev, curr, next
// Особенность: итеративный указатель на 4 строки (cleaner)

// Дано: 5-4-3-2-1
// Требуется: 1-2-3-4-5

package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

// Односвязный список
type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) Append(v int) {
	n := &Node{Val: v}
	if ll.Head == nil {
		ll.Head = n
		return
	}
	cur := ll.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = n
}

func (ll *LinkedList) Print() {
	cur := ll.Head
	for cur != nil {
		fmt.Printf("%d", cur.Val)
		if cur.Next != nil {
			fmt.Print(" -> ")
		}
		cur = cur.Next
	}
	fmt.Println()
}

// func main() {
// 	ll := &LinkedList{}
// 	ll.Append(1)
// 	ll.Append(2)
// 	ll.Append(3)
// 	ll.Append(4)
// 	ll.Append(5)

// 	ll.Print()
// }

func (ll *LinkedList) Reverse() {
	// три указателя - prev, curr, next

	// Будем хранить указатель на предыдущий элемент (копия указателя)
	var prev *Node
	// fmt.Println(prev)
	cur := ll.Head

	for cur != nil {
		// next - копия указателя, 8 байт
		// Сохраняем следующий
		next := cur.Next

		// Разворачиваем
		// Если первый проход, то nil, если не первый, то копия указателя
		cur.Next = prev

		// Двигаем prev
		prev = cur

		// Идем дальше, двигаем cur
		cur = next
	}

	// Новая голова
	ll.Head = prev
}

func main() {
	ll := &LinkedList{}
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.Append(4)
	ll.Append(5)

	ll.Print()

	ll.Reverse()

	ll.Print()
}
