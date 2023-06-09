package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke", counter)
		counter++
	}


	// kode program
	for counter1 := 1; counter1 <= 5; counter1++ {
		fmt.Println("perulangan ke", counter1)
	}

	// loop data like slice, array, and map
	slice := []string{"raihan", "faizah", "faiz"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// for range
	for i, value := range slice {
		fmt.Println("index", i, "=", value)
	}

	person := make(map[string] string) 
	person["nama"] = "malay"
	person["cewe"] = "faizah"
	person["laki"] = "faiz"

	for key, value:= range person {
		fmt.Println(key, "=", value)
	}
}