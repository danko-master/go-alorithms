package main

import (
	"fmt"
	"runtime"
	"sync"
)

// What will this code output?
func main() {
	// task1()
	// task2()
	// task3()
	// task4()
	// task5()

}

func task1() {
	fmt.Println("Task 1")

	m := map[string][]int{
		"a": {1, 2},
		"b": {3, 4},
	}

	// fmt.Println(m)
	// fmt.Println(m["a"]) // [1 2]
	// fmt.Println(m["b"]) // [3 4]

	for k, v := range m {
		// fmt.Println(v)
		v[0] = 9
		// fmt.Println(v) // [9 2] and [9 4]

		m[k] = append(v, 7)
	}

	fmt.Println(m["a"]) // [9 2 7]
	fmt.Println(m["b"]) // [9 4 7]
}

func task2() {
	fmt.Println("Task 2")

	a := []int{1, 2, 3}

	// fmt.Println(&a[0]) // 0x3aa820e02000
	// fmt.Println(&a[1]) // 0x3aa820e02008
	// fmt.Println(&a[2]) // 0x3aa820e02010

	p := &a[1] // 0x3aa820e02008

	// fmt.Println([]int(nil)) // [] - cap 0, len 0
	// fmt.Println(a[:2])      // [1, 2]
	// a[:2]... - unpacking - serial transmission to append 1 and 2
	b := append([]int(nil), a[:2]...) // new array, copying values into new array
	// b - new slice for new array

	a[0] = 9 // a -> [9 2 3]
	*p = 8   // a[1] = 8; a -> [9 8 3]

	fmt.Println(a)  // [9 8 3]
	fmt.Println(b)  // [1 2]
	fmt.Println(*p) // 8

}

func task3() {
	fmt.Println("Task 3")

	runtime.GOMAXPROCS(1)

	done := false

	go func() {
		done = true
	}()

	for !done {
	}

	// data race - print "finished" or hang forever
	fmt.Println("finished")
}

func task4() {
	fmt.Println("Task 4")

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		go func() {
			wg.Add(1) // This reason!
			defer wg.Done()

			fmt.Println(i) // Random print 0-1-2 or not print
		}()
	}

	wg.Wait()             // maybe panic, cause wg.Add() into goroutine, but wg.Wait() in this place
	fmt.Println("\ndone") // print done
}

// func task5() {
// 	fmt.Println("Task 5")

// }
