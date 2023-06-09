package main

import "fmt"

func main() {
	person := map[string]string {
		"nama": "raihan",
		"kepanjangan": "malay",
		"umur": "17",
	}

	person["kepanjangan"] = "player"
	length := len(person)

	fmt.Println(length)
	fmt.Println(person)
	fmt.Println(person["nama"])
	fmt.Println(person["kepanjangan"])

	delete(person, "umur")
	fmt.Println(person)
}