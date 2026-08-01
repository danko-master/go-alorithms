// 121. leetcode. Макс-я прибыль от одной сделки купить-продать

// O(n) time, O(1) memory
// Отслеживать minPrice, вычислять maxProfit на каждом шаге
// Без вложенных циклов

// [7,1,5,3,6,4] -> прибыль 5, купить за 1, продать за 6
// [7,6,4,3,1] -> прибыль 0

package main

import "fmt"

func main() {
	// arr := []int{7, 1, 5, 3, 6, 4}
	arr := []int{7, 6, 4, 3, 1}
	maxProfit(arr)
}

func maxProfit(prices []int) int {
	fmt.Println(prices)
	minPrice := prices[0]
	maxProfit := 0

	for _, p := range prices[1:] {
		if p < minPrice {
			minPrice = p
			fmt.Println("minPrice", minPrice)

		} else if profit := p - minPrice; profit > maxProfit {
			fmt.Println("Price2", p)
			fmt.Println("minPrice2", minPrice)

			maxProfit = profit

			fmt.Println("maxProfit", maxProfit)

		}
	}

	fmt.Println("minPrice3", minPrice)
	fmt.Println("maxProfit3", maxProfit)

	fmt.Println(maxProfit)

	return maxProfit
}
