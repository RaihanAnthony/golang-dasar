package main

import "fmt"

func main() {
	name := "raihan"
	counter := 0

// untuk menghindari perubahan nilai di dulur scope dalam yaitu dengam mendeklarasikan kedalam 
// variabel kembali

	increment := func() {
		name = "malay"
		fmt.Println("increment")
		counter++
	}

	increment()
	increment()

	fmt.Println(counter, name)
 } 