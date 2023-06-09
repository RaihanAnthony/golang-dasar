package main

import "fmt"

func factorialLoop(value int) int {
	// result := 1
	// for i := value; i > 0; i--{
	// 	result *= i // result = result * i
	// }
	// return result

	// recursive function
	if value == 1{
		return 1
	} else {
		return value * factorialLoop(value-1)
	}
} 

func main() {
	fmt.Println(factorialLoop(6))
	fmt.Println(6 * 5 * 4 * 3 * 2 * 1)
}