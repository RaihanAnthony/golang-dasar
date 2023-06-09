package main

import "fmt"

func getFullName() (string, string, string) {
	return "raihan", "malay", "cilegon"
}

func main() {
	firstName, middleName, _ := getFullName()
	fmt.Println(firstName)
	fmt.Println(middleName)
}