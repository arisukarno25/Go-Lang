package main

import "fmt"

func main()  {
	name := "Ari Sukarno"

	switch name {
	case "Ari Sukarno":
		fmt.Println("Nama Sudah Benar")
	case "Sukarno":
		fmt.Println("Nama Belum Lengkap")
	default:
		fmt.Println("Nama Salah")
		fmt.Println("Masukan nama dengan benar")
	}

	//short statement
	switch lenght := len(name); lenght > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama sudah bagus")
	}

	// memberi kondisit di setiap statement
	lenght2 := len(name)
	switch{ // tdk ada konisi, format umum switch kondisi {}
	case lenght2 > 5:
		fmt.Println("Terlalu panjang")
	case lenght2 > 3 :
		fmt.Println("Cukup")
	default:
		fmt.Println("Salah")
	}
}