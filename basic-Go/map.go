package main

import "fmt"

func main() {
	person := map[string]string{ // deklarasi panjangnya var person map[string]string = map[string]string, [string] kalau di array [0] dst
		"name" : "Ari", // "name" karena map[string]string, format map[tipedata-key]tipedata-value makanya kunci nya berupa string
		"address" : "Klaten",
	}
	//tambah key:value
	person ["pekerjaan"] = "Student"

	fmt.Println(person) // show key:value
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// buat map
	var book map[string]string = make(map[string]string)
	book["judul"] = "Tips Sukses Muda"
	book["author"] = "Ari Sukarno"
	book["hobby"] = "Makan"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "hobby") // delete map , delete(map, key)
	fmt.Println(book)
	fmt.Println(len(book))
}