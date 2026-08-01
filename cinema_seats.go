// Выбор мест в кинотеатре

// Дан ряд мест [1, 0, 0, 0, 1, 0, 1]
// 1 - занято
// 0 - свободно

// Нужно сесть так, чтобы максимально отдалиться от ближайшего зрителя

// Алгоритм - скользящее окно, считать макс-ю дистанцию

// [1,0,0,1] - макс.дист 2, сесть между двумя
// [1,0,1] - макс.дист 1, сесть посередине
// [1,0,0,1] - макс.дист 2, сесть слева

package main

import "fmt"

func main() {
	seats := []int{0, 0, 0, 0, 1, 0, 1, 0, 1, 0}
	maxDist, seatIdx := maxDistance(seats)

	fmt.Printf("Места: %v\n", seats)
	fmt.Printf("Макс-е расстояние: %d, место %d\n", maxDist, seatIdx)

	seats2 := []int{0, 0, 1, 0, 0, 0, 1, 1, 1}
	maxDist2, seatIdx2 := maxDistance(seats2)

	fmt.Printf("Места: %v\n", seats2)
	fmt.Printf("Макс-е расстояние: %d, место %d\n", maxDist2, seatIdx2)

}

func maxDistance(seats []int) (maxDist int, seatIdx int) {
	maxDist = 0
	seatIdx = -1

	// Флаг левой позиции
	// left := -1
	// Следующая после left
	// next := -1

	n := len(seats)

	// Счетчик прохода слева-направо
	cntr := 0

	// Пишем в map
	// Ключ "длина свободных", значение "индекс конца окна"
	freem := make(map[int]int)

	for i := 0; i < n; i++ {
		// Если свободно
		if seats[i] == 0 {
			cntr++

			// Пограничный случай - справа последнее место
			if i == n-1 {
				freem[cntr] = i
			}
		} else {
			freem[cntr] = i
			cntr = 0
		}

		// Следующее
		// next := i + 1

		// fmt.Println(i)
		// fmt.Println(s)

		// Крайнее слева
		// Если свободно, берем его, если нет, то пропускаем
		// if i == 0 && seats[i] == 0 {
		// 	seatIdx = i
		// }

		// Центральные - между крайними (лево->право)
		// if i > 0 && i < (len(seats)-1) {
		// 	if seats[i] == 0 {
		// 		if seatIdx == -1 {
		// 			seatIdx = i
		// 		} else {
		// 			next += 1
		// 		}
		// 	}

		// }

		// Крайнее справа

	}

	maxSum := 0
	maxIndx := -1
	for k, v := range freem {
		if k > maxSum {
			maxSum = k
			maxIndx = v
		}
	}

	fmt.Println(maxSum)
	fmt.Println(maxIndx)
	fmt.Println(n)

	maxDist = maxSum
	// Крайний справа
	if maxIndx == n-1 {
		fmt.Println(maxIndx == n-1)
		seatIdx = n
	} else {
		if maxIndx-maxSum == 0 {
			seatIdx = 1
		} else {
			seatIdx = maxIndx - (maxSum / 2)
		}
	}

	// fmt.Println(freem)

	return maxDist, seatIdx
}
