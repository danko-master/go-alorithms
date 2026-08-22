package main

import (
	"fmt"
	"math"
)

// Дано:
// Отсортированный массив по неубыванию
// [-7,-3,2,3,11]

// Задача:
// выдать отсортированный массив квадратов данных числе по неубыванию
// [4, 9, 9, 49, 121]

// Ограничение:
// O(n)

func main() {
	nums := []int{-10, -7, -3, 2, 3, 11, 20}

	// Последний элемент
	j := len(nums) - 1

	// Слайс результатов
	res := make([]int, len(nums))

	// Исходное состояние
	fmt.Println(nums)
	fmt.Println(res)

	fmt.Println("=== for ===")
	// Первый элемент для фиксации прохода по квадратам
	i := 0
	for k := len(res) - 1; k >= 0; k-- {
		vi := int(math.Abs(float64(nums[i])))
		vj := int(math.Abs(float64(nums[j])))

		// fmt.Println(vi)
		// fmt.Println(vj)
		// fmt.Println("--")

		if vi > vj {
			res[k] = vi * vi
			i++
		} else {
			res[k] = vj * vj
			j--
		}

	}

	// Результат [4 9 9 49 100 121 400]
	fmt.Println("=== res ===")
	fmt.Println(res)

}
