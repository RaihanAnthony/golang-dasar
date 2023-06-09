package main

import "fmt"

type Male struct {
	Name string
}


// cara yang direkomendasikan menggunakan pointer untuk struct method 
func (male *Male) Married() {
	 male.Name = "mr. " + male.Name 
}

func main() {
	raihan := Male{"yanto"}
	raihan.Married()

	fmt.Println(raihan.Name)
}