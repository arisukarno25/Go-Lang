package main

import "fmt"

func main()  {
	var name string

	name = "Ari Sukarno"
	fmt.Println(name)

	name = "Ari"
	fmt.Println(name)

	var othername = "Budi" // di definisikan langsung, otomatis go-lang tau kalau itu string
	fmt.Println(othername)

	var age = 30 			//atau bisa var age int8 = 30, bisa dispesifikin  
	fmt.Println(age)

	country := "indonesia" // inisiasi agar setelah nya ga usah ribet pake := atau tinggal = saja
	fmt.Println(country)

	country = "Singapure" 	//tdk perlu pake :, krn udh dideklarasiin sebelumnya
	fmt.Println(country)

	//multiple variable
	var (
		FirstName = "Ari"
		LastName = "Sukarno"
	)
	fmt.Println(FirstName)
	fmt.Println(LastName)
}	