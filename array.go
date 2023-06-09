 package main

 import "fmt"
 

func main() {
	values := [...] int {
		20,
		10,
		20,
	}
	
	firstvalue := len(values)
	values[1] = 20

	fmt.Println(firstvalue)
	fmt.Println(values[2])
	fmt.Println(values)

}
 
 