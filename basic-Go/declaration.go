package main 

import "fmt"

func main()  {
	// membuat alias dari tipe data yang sudah ada
	type NoKTP string
	type Married bool

	var myKTP NoKTP = "12343566"
	var Status Married = true
	fmt.Println(myKTP)
	fmt.Println(Status)
}