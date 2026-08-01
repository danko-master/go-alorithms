// 242 Valid anagram
// Является ли t анаграмой s

// O(n) time
// O(k) memory (k - кол-во уникальных run-ов)

// Особенности от латиницы:
// 1. len(s) != len(t) - т.к. длина в байтах, не в символах
// 2. for range для расповки rune
// 3. ранний return при отрицательном счетчике

// Пример:
// s="кот", t="ток" -> true
// s="лес", t="село" -> false (разная длина)
// s="карета", t="ракета" -> true

package main

import "fmt"

func main() {
	// s := "карета"
	// t := "ракета"
	s := "лесм"
	t := "село"
	res := isAnagramSimple(s, t)
	fmt.Println(res)
}

func isAnagramSimple(s, t string) bool {
	rs := []rune(s)
	rt := []rune(t)

	fmt.Println(rs)
	fmt.Println(rt)

	if len(rs) != len(rt) {
		return false
	}

	// Счетчик для всей кириллицы + латиницы
	cnt := make(map[rune]int, len(rs))

	for i := 0; i < len(rs); i++ {
		cnt[rs[i]]++
		cnt[rt[i]]--
	}

	fmt.Println(cnt)

	for _, v := range cnt {
		if v != 0 {
			return false
		}
	}

	return true
}
