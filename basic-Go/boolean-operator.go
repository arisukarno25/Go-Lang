package main

import "fmt"

func main()  {
	var nilaiAkhir = 80
	var nilaiAbsensi = 75

	var lulusUjian = nilaiAkhir >= 80
	var lulusAbsensi = nilaiAbsensi > 80
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var hasilAkhir = lulusUjian && lulusAbsensi
	fmt.Println(hasilAkhir)

	// best practice seperti ini
	fmt.Println(nilaiAkhir >= 80 && nilaiAbsensi > 80)
	fmt.Println("Anda lulus?", nilaiAkhir >= 80 && nilaiAbsensi > 80)
}