package main

import "fmt"

func ups() interface{} {
	// return 1
	// return true
	return "ups"
}

func uups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2{
		return true
	} else {
		return "ups"
	}
}

func main() {
	kosong := ups()
	fmt.Println(kosong)

	number := uups(1)
	fmt.Println(number)
}