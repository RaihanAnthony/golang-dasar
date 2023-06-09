package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("permabagian dengan NOL")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	hasil, err  := pembagian(10, 10)
	if err == nil{
		fmt.Println("hasil", hasil)
	} else {
		fmt.Println("error", err.Error())
	}
}