package main

import "fmt"

type Addres struct{
	city, province, country string
}

func changeCountryToIndonesia(Addres *Addres) {
	Addres.country = "indoensia"
}

func main() {
	// pass by value
	var adders1  Addres = Addres {
		city : "serang",
		province : "banten",
		country : "indonesia",
	}

	adders1.city = "cilegon"

	// pointer
	// pass by rafrance untuk menjadikan addres2 sebagain memory utama dengan 
	// oprator &
	adders2 := &adders1

	// operator * untuk menjadikan main memory
	*adders2 =  Addres{"tanah abang", "jakarta", "indonesia",}
	
	fmt.Println(adders1)
	fmt.Println(adders2) 

	
	// membuat pointer dengan function new
	alamat1 := new(Addres)
	alamat2 := alamat1

	alamat2.country = "singapura"

	
	fmt.Println(alamat1)
	fmt.Println(alamat2)


	// pointer di function
	var Alamat = Addres{
		city: "bogor",
		province: "jawa barat",
		country: "",
	}

	var AlamatPointer = &Alamat
	changeCountryToIndonesia(AlamatPointer)
	fmt.Println(Alamat)
}
	