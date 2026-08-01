// 49 Group Anagrams

// Сгруппировать анаграммы (слова из одних и тех же букв)

// O(n*k) time
// O(n*k) memory
// n = len(strs)
// k = max(len(s))

// Алгоритм - частотный массив [26]byte - ключ
// Особенность - [26]byte - хэшируемый тип в Go, можно использовать как ключ map

// Пример:
// ["eat","tea","tan","ate","nat","bat"]
// -> [["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]

package main

import "fmt"

func main() {
	strs := []string{
		"eat", "tea", "tan", "ate", "nat", "bat",
	}
	// var key [26]byte
	// fmt.Println(key)
	fmt.Println(groupAnagram(strs))
}

func groupAnagram0(strs []string) [][]string {
	m := make(map[[26]byte][]string)

	for _, s := range strs {
		// key - [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
		var key [26]byte
		for i := 0; i < len(s); i++ {
			key[s[i]-'a']++
		}
		m[key] = append(m[key], s)
	}

	// тип слайс слайсов строк, длина 0б ёмкоть равна количеству групп map
	r := make([][]string, 0, len(m))

	for _, v := range m {
		r = append(r, v)
	}

	return r
}

func groupAnagram(strs []string) [][]string {
	// res := [][]string{}

	m := make(map[int][]string)

	for _, v := range strs {
		// fmt.Println(i)
		// fmt.Println(v)

		ksum := 0
		for _, r := range v {
			// fmt.Println(r - 'a')

			ksum += int(r - 'a')
		}

		// fmt.Println(ksum)

		// Слайс ссылочный тип - data, len, cap
		m[ksum] = append(m[ksum], v)

		// fmt.Println(m)

	}

	res := make([][]string, 0, len(m))

	for _, v := range m {
		res = append(res, v)
	}

	return res
}
