package main 

import "fmt"

func main()  {
	var nama1 = "Ari"
	var nama2 = "Ari"

	var banding bool = nama1 == nama2
	fmt.Println(banding)
	//atau 
	fmt.Println(nama1==nama2)

	var value1 = 100
	var value2 = 200
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}