package main

import "fmt"

func sayHello(firstName string, lastName string){
	fmt.Println("Hello", firstName, lastName)
}

// function parameter -> untuk menambahkan data dari luar fucntion

func main() {
	sayHello("Ari", "Sukarno")
	firstName := "Ari" // bisa di define dulu
	sayHello(firstName, "Sukarno")
}