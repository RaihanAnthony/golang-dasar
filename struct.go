package main

import "fmt"

type Custemer struct {
	Name, Addres string
	Age int
}


// struct method or function struct
func (custemer Custemer) sayhay(Name string) {
	fmt.Println("hallo", Name, "my name is", custemer.Name)
}

func main() {
	// fill data in struct many way
	
	// one way
	var raihan Custemer
	raihan.Name = "Raihan malay"
	raihan.Addres = "Indonesia"
	raihan.Age = 21

	raihan.sayhay("toni")

	// fmt.Println(raihan.Name)
	// fmt.Println(raihan.Addres)
	// fmt.Println(raihan.Age)

	// // second way
	//  faizah := Custemer {
	// 	Name: "Qurota ayun",
	// 	Addres: "Batam",
	// 	Age: 27,
	//  }

	//  fmt.Println(faizah)
	
	// // theared way but not recomendation
	// faiz := Custemer{"maulana", "banten", 25}

	// fmt.Println(faiz)
}