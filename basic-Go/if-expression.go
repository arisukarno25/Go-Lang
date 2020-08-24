package main

import "fmt"

func main() {
	name := "Ari Sukarno"
	
	if name == "Ari Sukarno"{
		fmt.Println("Benar")
	} else if name == "Ari"{
		fmt.Println("Benar, tapi tidak lengkap")
	} else if name == "Sukarno"{
		fmt.Println("Benar, tapi tidak lengkap")
	} else {
		fmt.Println("Salah")
	
	//short statement versi panjang
	//var length = len(name) // len() bisa juga cek panjang string
	//versi short statement
	if length := len(name); length > 5 {
		fmt.Println("Karakter terlalu panjang")
	} else {
		fmt.Println("Karakter sudah benar")
	}
}