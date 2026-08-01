package main

import (
	"fmt"
	"unsafe"
)

func main() {

	sl := make([]int, 5, 7)

	data := []int{0, 1, 2, 3, 4}
	copy(sl, data)

	slPtr := unsafe.SliceData(sl)

	fmt.Println(slPtr)

	f1(sl)

	// Поменялся массив, но в структуре len и cap данного инстанса остались прежние
	fmt.Println(len(sl))
	fmt.Println(cap(sl))
	fmt.Println(sl)

	// sl2 := sl[1:2]
	// slPtr2 := unsafe.SliceData(sl2)

	// fmt.Println(slPtr2)

	// fmt.Println(cap(sl2))
	// fmt.Println(len(sl2))

	// f1(sl2)

	// fmt.Println(sl2)

	// fmt.Println(sl)

}

func f1(sl []int) {
	slPtr := unsafe.SliceData(sl)

	fmt.Println(slPtr)

	sl = append(sl, 9)
	fmt.Println(len(sl))
	fmt.Println(cap(sl))
	fmt.Println(sl)
}
