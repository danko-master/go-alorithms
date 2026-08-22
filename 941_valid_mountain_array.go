// Given an array of integers arr, return true if and only if it is a valid mountain array.

// Recall that arr is a mountain array if and only if:

//     arr.length >= 3
//     There exists some i with 0 < i < arr.length - 1 such that:
//         arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
//         arr[i] > arr[i + 1] > ... > arr[arr.length - 1]

// 		Example 1:

// 		Input: arr = [2,1]
// 		Output: false

// 		Example 2:

// 		Input: arr = [3,5,5]
// 		Output: false

// 		Example 3:

// 		Input: arr = [0,3,2,1]
// 		Output: true

// 		Constraints:

// 			1 <= arr.length <= 10^4
// 			0 <= arr[i] <= 10^4

package main

import "fmt"

func main() {
	arr1 := []int{2, 1}
	// false
	fmt.Println("arr1 - res ", validMountainArray(arr1))

	arr2 := []int{3, 5, 5}
	// false
	fmt.Println("arr2 - res ", validMountainArray(arr2))

	arr3 := []int{0, 3, 2, 1}
	// true
	fmt.Println("arr3 - res ", validMountainArray(arr3))

	arr4 := []int{1, 3, 2}
	// true
	fmt.Println("arr4 - res ", validMountainArray(arr4))

	arr5 := []int{6, 7, 7, 8, 6}
	// false
	fmt.Println("arr5 - res ", validMountainArray(arr5))
}

func validMountainArray(arr []int) bool {
	// Если в массиве менее 3 элементов, то невозможно найти пик
	if len(arr) < 3 {
		return false
	}

	// Флаг достижения пика
	// Изначально не были на пике
	wasntPeak := true
	for i, j := 0, 1; i <= len(arr)-2; i++ {
		vali := arr[i]
		valj := arr[j]

		// fmt.Println(vali)
		// fmt.Println(valj)
		// fmt.Println(wasntPeak)

		// Если равные элементы
		if vali == valj {
			return false
		}

		// Вершину еще не встретили
		if wasntPeak {
			if i == 0 && vali > valj {
				return false
			}

			// Встретили вершину
			if vali > valj {
				wasntPeak = false
			}
		} else {
			// Были на вершине, но вдруг снова начался подъем
			if valj >= vali {
				return false
			}
		}
		// fmt.Println(i)
		// fmt.Println(j)

		j++

	}

	// fmt.Println(wasntPeak)
	// Если так и не достигли вершины
	if wasntPeak {
		return false
	}

	return true
}
