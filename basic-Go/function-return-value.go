package main

import "fmt"

func getHello(name string) string { // string setelah parameter adalah tipe data untuk pengembaliannya
	// result := "Ari"
	// return "Hello" + name
	if name == "" {
		return "Hello bro !!" // return diikuti_value_data_berdasarkan_tipe_data_return
	} else {
		return "Hello" + name
	}
}

func main() {
	// result := getHello("Ari")
	// fmt.Println(result)
	
	result := getHello("Ari")
	fmt.Println(result)
	// atau bisa lngsng
	fmt.Println(getHello("Ari"))
	fmt.Println(getHello(""))
}