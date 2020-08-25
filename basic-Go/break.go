package main

import "fmt"

func main() {
	// setiap kali perulangan menemui kata break maka perulangan akan dihentikan saat itu juga 
	for i := 1; i < 10; i++ {
		//fmt.Println("Perulangan ke-", i) // jika ditempatkan disini maka "Perulanagn ke-5 akan di tampilkan karena sblm break"
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke-", i) // gas di print karena sudah di break sebelumnya 
	}
}