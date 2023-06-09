package main

import "fmt"

func main() {
	name := "malay"

	switch name {
	case "raihan":
		fmt.Println("hello raihan")
	case "ortus":
		fmt.Println("hallo ortus")
	default:
		fmt.Println("kenalkan nama saya malay")
	}

	// switch length := len(name); length > 5{
	// case true:
	// 	fmt.Println("nama terlalu panjang")
	// case false:
	// 	fmt.Println("nama terlalu pendek")
	// }

	length := len(name)
	switch {
	case length > 5:
		fmt.Println("nama terlalu panjang")
	case length < 5:
		fmt.Println("nama terlalu pendek")
	default: 
		fmt.Println("nama sudah benar")
	} 
}