// 3. leetcode

// Найти длину самой длинной подстроки без повторяющихсмя символов
// O(n) time
// O(min(n, alphabet)) memory

// Алгоритм: скользящее окно - map[byte]int (символ -> индекс + 1)
// Особенность - при сдвиге left берем max(left, lastSeen+i)

// Пример:
// "abcabcbb" - длина 3 ("abc")
// "bbbbb" - длина 1 ("b")
// "pwwkew" - длина 3 ("wke")

package main

import "fmt"

func main() {

	s := "abcabcdbbabcd"
	fmt.Println([]byte(s))

	lens := longestSubstr(s)
	fmt.Println(lens)
}

func longestSubstr(s string) int {
	// храним следующий индекс после последнего вхождения
	// Строка в Go - массив байт
	last := make(map[byte]int)
	// left - сдвиг за повтор
	// best - обновляем максимум

	// Строка - слайс байт
	// b := []byte("abc")
	// [97, 98, 99]

	// Строка - слайс рун
	// b := []rune("привет")
	// [1087 1088 1080 1074 1077 1090]

	left, best := 0, 0

	for i := 0; i < len(s); i++ {
		c := s[i]
		fmt.Println(c)

		// pos - значение по ключу "с", иначе zaro value (в данном случае 0)
		// ok - true если ключ существует, false если ключа нет

		// если pos >= left - повтор внутри окна c
		// сужаем окно
		pos, ok := last[c]
		fmt.Println(pos)
		fmt.Println(ok)
		fmt.Println(pos >= left)
		fmt.Println(left)
		fmt.Println("===========")

		if ok && pos >= left {
			left = pos + 1
		}

		if i-left+1 > best {
			best = i - left + 1
		}

		last[c] = i
	}

	return best
}
