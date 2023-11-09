// пример решения домашнего задания модуль 22.4
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 10

func main() {
	rand.NewSource(time.Now().UnixNano())

	var numbers [size]int
	for i := 0; i < size; i++ {
		numbers[i] = rand.Intn(10 * size)
	}

	fmt.Printf("%+v\n", numbers)

	n := 0
	fmt.Scan(&n)

	count := 0 //количество чисел после найденного числа в массиве
	found := false

	for i := 0; i < size; i++ {
		if found {
			count++
		} else if numbers[i] == n {
			found = true
		}
	}
	fmt.Println("Numbers after yours:", count)
}
