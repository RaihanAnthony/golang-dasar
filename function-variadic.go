package main

import "fmt"

func sumAll(value ...int) int {
	total := 0
	for _, value := range value {
		total += value
	}
	return total
}

func main() {
	amount := sumAll(10, 20, 20, 50)
	fmt.Println(amount)

	// menjadikan slice sebagai variadic 
	slice := []int{20, 96, 87, 90}
	amount = sumAll(slice...)
	fmt.Println(amount)
}