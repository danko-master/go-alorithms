// 49.leetcode

// Сгруппировать анаграммы (слова из одних и тех же букв)
// Сложность: O(n*k) time, O(n*k) memory
// n = len(strs), k = max(len(s) в рунах)

// Алгоритм: частотный map[rune]int- каноническая строка

// Отличие от латиницы:
// кириллица занимает 2 байта на символ - нельзя сделать `s[i] - 'a'`
// используем for range для распаковки rune
// ключ - строка из отсортированных run-ов (эквивалент fingerprint)

// Пример:
// Запрос - ["кот", "ток", "вода", "дом", "кто", "мод"]
// Ответ - [["кот", "ток",  "кто"], ["вода"], ["дом", "мод"]]

package main

import "sort"

func main() {

}

func groupAnagramsRU(strs []string) [][]string {
	m := make(map[string][]string)

	// s - КОПИЯ строки из исходного массива
	for _, s := range strs {
		// конвертация строки в слайс юникод символов - "п" -> 1087, "р" -> 1088, "и" -> 1080, "в" -> 1074, "е" -> 1077, "т" -> 109
		runes := []rune(s)
		// Например
		// "ток" [1090, 1086, 1082]

		// sort.Slice(x any, less func(i, j, int) bool)
		// Принимает x any - пустой интерфейс
		// Через reflect получает reflect.ValueOf(x)
		// Проверяет Kind() == reflect.Slice
		// Получает Len()
		// и доступ через Index() / Swap()
		// sort.Slice синтаксический сахар над sort.Interface
		// type Interface interface {
		// 	Len() Interface
		// 	Less(i, j int) bool
		// 	Swap(i, j int)
		// }

		// рановсильно
		// можно так:
		// type RuneSlice []rune
		// func (rs RuneSlice) Len() int { return len(rs) }
		// func (rs RuneSlice) Less(i, j int) bool { return rs[i] < rs[j] }
		// func (rs RuneSlice) Swap(i, j int) { rs[i], rs[j] = rs[j], rs[i] }
		// sort.Slice(RuneSlice(runes))

		// но пишем так:
		// sort.Slice(runes, func(i, j int) bool) {
		// 	return runes[i] < runes[j]
		// })

		// сортируем слайс на месте
		// было "ток" [1090, 1086, 1082], стало "кот" [1082, 1086, 1090]
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		// Собираем строку обратно
		// key = "кот"
		key := string(runes)

		// m["кот"] = append(m["кот"], "ток")
		// "ток" попадет в группу с "кот"
		m[key] = append(m[key], s)
	}

	// Создаем слайс слайсов, длина 0 и емкость len(m)
	r := make([][]string, 0, len(m))
	// nil, nil - cap = 2, len = 0
	// Без емкости будет переаллокация массива и перекопировании элементов
	// С емкостью выделен массив на максимальное количество групп, длина = 0 , append использует остаток без копирования

	for _, v := range m {
		r = append(r, v)
	}

	// Эквивалент
	// r := make([][]string, len(m))
	// i := 0
	// for _, v := range m {
	// 	r[i] = v
	// 	i++
	// }

	return r
}
