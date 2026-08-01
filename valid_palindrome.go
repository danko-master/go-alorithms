// Проверить является строка палиндромом - с учетом только букв и цифр

// O(n) time
// O(1) memory

// Алгоритм - два указателя с обоих концов
// Особенность - без аллокаций extra filtered string

// "A man, a plan, a canal: Panama" -> true
// "race a car" -> false

package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {

	l, r := 0, len(s)-1

	for l < r {
		ls := rune(s[l])
		rs := rune(s[r])

		// будем пропускать не буквы и не цифры, поэтому внутренние for двигаются только впереде или назад
		// каждый байт проверяется один раз

		// Пропускаем не буквы слева
		for !isAlnum(ls) && l < r {
			l++
			ls = rune(s[l])
		}

		// Пропускаем не буквы справа
		for !isAlnum(rs) && l < r {
			r--
			rs = rune(s[r])
		}

		if unicode.ToLower(ls) != unicode.ToLower(rs) {
			return false
		}

		l++
		r--
	}

	return true
}

// Проверка на букву и цифру
func isAlnum(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}
