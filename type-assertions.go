package main

import "fmt"

func random() interface{} {
	return "jancok"
}

func main() {
	result := random()
	resultString := result.(string)
	fmt.Println(resultString)


// type assertions with switch
	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("unknown")
	}
}