package main

import "fmt"

func main()  {
	const FirstName string = "Ari" // bisa dikasi tipe data 
	const LastName = "Sukarno" // bisa tdk dikasi, krn go-lang udh otomatis tau
	const age = 20
	 // beda dengan variable adlh, jika varible ga dipake maka error, tp const misal ga dipake ttp aman
	 // const tidak bisa diubah lagi
	fmt.Println(FirstName)
	fmt.Println(LastName)
	fmt.Println(age)
	const (
		Firstname	= "Andi"
		Lastname	= "Setiawan"
		Age			= 21
	)
	fmt.Println(Firstname)
	fmt.Println(Lastname)
	fmt.Println(Age)
 }