// 242 Valid anagram
// Является ли t анаграмой s

// O(n) time
// O(1) memory (алфавит 26 букв)

// Алгоритм:
// счетчик частот на [26]int

// Особенность:
// можно без map - go гарантирует [26]int{} == zeroed
// [26]int{} - пустой массив, все элементы 0

// [26]int{} - массив на стеке

// make([]int, 26) - массив в куче (аллокация + GC)
// т.к. размер неизвестен на этапе компилляции
// вызывается runtime.makeslice -> выделяет память на куче
// type slice struct {
// 	data *int // указатель на underlaying array
// 	len int // длина
// 	cap int // емкость
// }

// Когда может остаться в стеке:
// escape analysis

// small slice может уйти на стек
// func small() {
// 	s := make([]int, 4)
// 	s[0] = 42
// 	return s[0]
// }
// Если слайс не убегает из функции, может попасть в стек.
// Но большие размеры уйдут в кучу

package main

import "fmt"

func main() {
	s := "break"
	t := "baker"

	fmt.Println(isAnagram(s, t))
}

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// [0,0,0,0...0] - все 26 нулей
	var cnt [26]int

	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
		cnt[t[i]-'a']--
	}

	return cnt == [26]int{}
}
