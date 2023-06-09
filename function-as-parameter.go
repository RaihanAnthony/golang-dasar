package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
	nameFilterd := filter(name)
	fmt.Println("hello", nameFilterd)
}

func spamFilter(name string) string{
	if name == "anjing" {
		return "...."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("raihan", spamFilter)
	sayHelloWithFilter("anjing", spamFilter)
}