// 823. Binary Trees With Factors
// Medium
// Topics
// premium lock icon
// Companies

// Given an array of unique integers, arr, where each integer arr[i] is strictly greater than 1.
// 		Дан массив уникальных целых чисел arr, где каждое число arr[i] строго больше 1.

// We make a binary tree using these integers, and each number may be used for any number of times.
// Each non-leaf node's value should be equal to the product of the values of its children.
// 		Из этих чисел мы составляем бинарное дерево, при этом каждое число может использоваться любое количество раз.
// 		Значение каждого нелистового узла должно быть равно произведению значений его дочерних узлов.
// 		*Лист бинарного дерева — это узел, у которого нет никаких потомков (детей),
// 		 то есть его указатели на левого и правого ребенка равны null или пусты.

// Return the number of binary trees we can make.
// The answer may be too large so return the answer modulo 10^9 + 7.
// 		Верните количество возможных бинарных деревьев.
// 		Ответ может быть слишком большим, поэтому верните его по модулю 10^9 + 7.

// Example 1:

// Input: arr = [2,4]
// Output: 3
// Explanation: We can make these trees: [2], [4], [4, 2, 2]
// Example 2:

// Input: arr = [2,4,5,10]
// Output: 7
// Explanation: We can make these trees: [2], [4], [5], [10], [4, 2, 2], [10, 2, 5], [10, 5, 2].

// Одноэлементные массивы (например, [2], [5]) — это объявления существующих узлов или листьев.
// Трехэлементные массивы (например, [4, 2, 2], [10, 2, 5]) — это правила ветвления в формате
// 							[родитель, левый_потомок, правый_потомок].

// [2]	   [4] 		[5] 			[10]
//  x	    |		 x				|  |
// 		[4, 2, 2] 			[10, 2, 5] [10, 5, 2]

// Constraints:
// 1 <= arr.length <= 1000
// 2 <= arr[i] <= 10^9
// All the values of arr are unique.

package main

import (
	"fmt"
	"sort"
)

func main() {
	// arr1 := []int{2, 4}
	// fmt.Println(numFactoredBinaryTrees(arr1))

	// arr2 := []int{2, 4, 5, 10}
	// fmt.Println(numFactoredBinaryTrees(arr2))

	arr3 := []int{18, 3, 6, 2}
	fmt.Println(numFactoredBinaryTrees(arr3))
}

func numFactoredBinaryTrees(arr []int) int {
	fmt.Println(arr)

	const mod = 1_000_000_007
	sort.Ints(arr)

	// dp maps a number to the total count of valid binary trees rooted at that number
	dp := make(map[int]int)
	totalTrees := 0

	for i, parent := range arr {
		// Base case: A single node tree with just the parent itself
		dp[parent] = 1

		// Find pairs (left, right) such that left * right == parent
		for j := 0; j < i; j++ {
			left := arr[j]

			// Optimization: If left * left > parent, no more pairs can exist
			if left*left > parent {
				break
			}

			// Check if left is a factor and the complementary right factor exists
			if parent%left == 0 {
				right := parent / left
				if rightCount, exists := dp[right]; exists {
					// Ways to form trees with this specific pair
					ways := (dp[left] * rightCount) % mod

					// If factors are distinct (e.g., 2 and 5 for 10), we can swap left and right
					if left != right {
						ways = (ways * 2) % mod
					}

					dp[parent] = (dp[parent] + ways) % mod
				}
			}
		}
		totalTrees = (totalTrees + dp[parent]) % mod
	}

	return totalTrees
}
