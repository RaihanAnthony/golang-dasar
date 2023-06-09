package main

import "fmt"

func getFullName2() (firstName string, middleName string, lastName string) {
	firstName = "raihan malay"
	middleName = "faizah qurota ayun"
	lastName = "faiz maulana"
	return
}

func main() {
	_, _, c := getFullName2()
	fmt.Println(c)
	fmt.Println(getFullName2())
}