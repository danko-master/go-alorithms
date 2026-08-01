// 1 Two sum
// Найти два числа, дающих в сумме target

// O(n) time
// O(n) memory

// Однослойный рпоход, без предзаполнения map

// [2, 7, 11, 15], target = 9 -> [0, 1]
// [3, 2, 4], target = 6 -> [1, 2]

package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 26
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	res := []int{}

	// Будем хранить значение -> индекс
	seen := make(map[int]int, len(nums))

	for i, v := range nums {
		// Остаток до target
		compl := target - v

		// j - по значению получаем индекс элемента
		// exists - true ключ найден, false ключ не найден
		if j, exists := seen[compl]; exists {
			return []int{j, i}
		}

		// Не нашли, запоминаем текущее число для будущих проверок
		seen[v] = i
	}

	return res
}
