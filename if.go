package main

import "fmt"

func main() {
	name := "raihan"

	if name == "malay" {
		fmt.Println("hello guys")
	} else if name == "zeze" {
		fmt.Println("kenalin raihan")
	} else {
		fmt.Println("kenalkan nama saya raihan")
	}

	// short statement
	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah sesuai")
	}
}