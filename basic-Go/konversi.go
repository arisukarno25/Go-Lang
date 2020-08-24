package main

import "fmt"

func main()  {
	var nilai32 int32 = 1000000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32) //ketika nilai32 dikonversi ke int8 maka error, karena maks nya 128

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	// konversi indeks byte ke string
	var name = "Ari Sukarno"
	var e byte = name[0] // byte = uint8
	var eString string = string(e) // konversi byte 0 -> string

	fmt.Println(name)
	fmt.Println(eString)

}