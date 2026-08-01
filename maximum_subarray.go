// 53 leetcode

// Найти подмассив с максимальной суммой

// O(n) time
// O(1) memory

// Алгоритм kadane - при каждом i решаем брать элемент в текущий подмассив или начать новый
// Классический DP за O(1)

package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, -4}
	best := maxSubArraySum(nums)
	fmt.Println(best)
}

func maxSubArraySum(nums []int) int {
	// Лучшая сумма
	best := nums[0]
	// Текущая сумма
	cur := nums[0]

	for i := 1; i < len(nums); i++ {
		// Если cur < 0 - начинаем с nums[i]
		// Иначе продолжаем
		// max возвращает большее из двух чисел
		cur = max(nums[i], cur+nums[i])
		best = max(best, cur)
	}

	return best
}
