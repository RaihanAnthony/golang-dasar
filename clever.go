package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil function")
}

func runAppLication() {
	defer logging()
	fmt.Println("run application")
}

func main() {
	runAppLication()
}