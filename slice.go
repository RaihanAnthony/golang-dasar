package main

import "fmt"

func main() {
	days := [...] string {
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"minggu",
	}

	slice := days[3:5]
	fmt.Println(slice[1])
	fmt.Println(len(slice))
	fmt.Println(cap(slice))


	daysSlice := days[5:]
	daysSlice[0] = "sabtu baru"
	daysSlice[1] = "minggu baru"
	fmt.Println(daysSlice)
	
	daysSlice1 := append(daysSlice, "okeyyy")
	daysSlice1[0] = "libur oyy"
	fmt.Println(daysSlice1)
	fmt.Println(days)

	// membuat slice from zero
	mounts := make([]string, 2, 5) 
	mounts[0] = "januari"
	mounts[1] = "februari"

	fmt.Println(mounts[1])

	// mengcopy
	copySlice := make([]string, len(mounts), cap(mounts))
	copy(copySlice, mounts)

	fmt.Println(copySlice)

	// different
	iniArray := [...] int {1, 2, 3, 4, 5}
	iniSlice := [] int {1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}