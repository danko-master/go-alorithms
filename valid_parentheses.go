// 20 проверка корректности скобочной последовательности

// O(n) time
// O(n) memory

// Алгоритм:
// стек - открывающие складываем, закрывающие проверяем
// map закрывающая -> открывающая для чистоты кода

// () - true
// ()[]{} - true
// (] - false
// ([)] - false
// {[]} - true

package main

import "fmt"

func main() {
	s := "{[]}"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	// map закрывающей - открывающей скобки
	// pairs[')'] = '('
	// pairs[']'] = '['
	// pairs['}'] = '{'

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	// Пустой слайс байт, с емкостью равной длине строки
	// длина 0
	stack := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			// Открывающая - пушим в стек
			stack = append(stack, s[i])
		case ')', ']', '}':
			fmt.Println(stack)
			// Закрывающая - проверяем, что последняя открывающая - пара
			// stack[len(stack)-1] - смотрим на последний элемент
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			// pop - удаление последнего элемента
			// stack[:len(stack)-1] - срез без последнего элемента
			// O(1) - быстрее и дешевле
			stack = stack[:len(stack)-1]

			// Если сделать stack = stack[1:] сдвиг всего ("дорого")

		}
	}

	return len(stack) == 0
}
