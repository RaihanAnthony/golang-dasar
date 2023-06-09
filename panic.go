package main

import "fmt"

func endApp() {
	
	massage := recover()
	if massage != nil {
		fmt.Println("ERROR dengan massage : ", massage)
	}
	fmt.Println("aplikasi selesai")
}

func runApp(err bool) {
	defer endApp()
	if err {
		panic("aplikasi ERR") 
	}
	fmt.Println("aplikasi berjalan")
}

func main() {
	runApp(true)
	fmt.Println("memey")
}