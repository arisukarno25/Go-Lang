package main

import "fmt"

func main(){
	// dalam go-lang cma ada 1 perulangan yaitu for, seperti do while tdk ada	
	counter := 1
	for counter <= 10 {
		fmt.Println("Counter ke-", counter)
		counter++ // counter = counter + 1
	}

	// init dan post statement 
	for counter2 := 1; counter2 <= 10; counter2++ { // counter2 :=1 atau init hanya dieksekusi sekali saja, counter2++ akan di eksekusi terus saat kondisi berjalan
		fmt.Println("Counter2 ke-", counter2)
	}

	//for untuk iterasi data thdp data collection (array, slice dan map)
	slice := []string{"Ari","Sukarno","Andi","Kurniawan"}

	for i := 0; i < len(slice); i++ {
		//fmt.Println(i)
		fmt.Println(slice[i])
	}
	
	//for range -> bisa akses sekaligus
	//contoh di slice/array
	for i, value := range slice { // for key, value range nama-variable {}, for key, value karena slice adalah slice 
		fmt.Println("Index", i, "=", value) // i bisa diganti dengan "_" agar i bisa di lompati/tdk kedetect tdk dipake
	}
	//contoh map
	person := make(map[string]string)
	person["name"] = "Ari"
	person["address"] = "Klaten"
	for key, value2 := range person {
		fmt.Println(key, "=", value2)
	}
	// summary 
	// for range pada array/slice -> for indeks, value := range varible {}
	// for range pada map -> for key, value := range varible {}
}