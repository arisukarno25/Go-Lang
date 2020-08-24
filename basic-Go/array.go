package main

import "fmt"

func main()  {
	var names [3]string

	names[0] = "Ari"
	names[1] = "Sukarno"
	names[2] = "SKA"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		80,
		70,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(names)) // panjang array/data nya bkn isi datanya
	fmt.Println(len(values))

	var test [10]int
	fmt.Println(len(test)) // kedetek 10 dari sebelumna yg udh dikash panjang array nya bkn datanya
}