package main

import "fmt"

func main()  {
	var jumlah = 10 + 10
	fmt.Println(jumlah)
	
	var a = 10
	var b = 20
	var c = a + b
	fmt.Println(c)

	var i = 10
	i += 10 // i = i sebelummya + 10 atau i = i + 10
	// a += 5 // a harus di deklarasikan dl, krn kalau gini artinya a = a + 10
	fmt.Println(i)

	i++ // artinya i = i + 1
	fmt.Println(i)
	i--
	fmt.Println(i)
}