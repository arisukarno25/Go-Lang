package main

import "fmt"

func main()  {
	var months = [...]string{ // ... artinya jika kapasitas tdk tau, kalau tau bisa lngsng di define
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}
	// program 1
	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1)) // melihat panjang data slice
	fmt.Println(cap(slice1)) // kapasitas dihitung dari batas awal sampai akhir, ini dari 4 - 12

	months[5] = "juni-diubah" // mengubah slice sama aja mengubah array
	fmt.Println(slice1) // indeks array 5 telah diubah dari slice

	slice1[0] = "Mei-diubah" // indeks 0 pada slice 
	fmt.Println(months) // data array month berubah

	// program 2
	// append slice atau nambah slice pada array
	//var slice2 = months[10:]
	//fmt.Println(slice2)
	
	//var slice3 = append(slice2, "Ari") // masih bisa ditambah di slice3 nya dgn cara buat aray baru, tp bukan di array nya
	//fmt.Println(slice3) 
	
	//slice3[1] = "desember-diubah"
	//fmt.Println(slice3)

	//fmt.Println(slice2) 
	//fmt.Println(months) // ari ga bisa nambah di array krn udh penuh dan tdk bisa buat array baru

	// membuat slice sendiri, dengan ini menghemat penulisan dan array akan disembunyikan
	//newSlice := make([]string, 2, 5)
	//newSlice[0] = "Ari"
	//newSlice[1] = "Sukarno"

	//fmt.Println(newSlice)
	//fmt.Println(len(newSlice))
	//fmt.Println(cap(newSlice))

	// copy slice
	//copySlice := make([]string, len(newSlice), cap(newSlice)) // pstikan len dan cap sama agar tdk kepotong
	//copy(copySlice, newSlice) // copy(dstination, source)
	//fmt.Println(copySlice)

	// array vs slice
	//iniArray := [5]int{1,2,3,4,5} // [indeks] atau [...]
	//iniSlice := []int{1,2,3,4,5} 
	//fmt.Println(iniArray)
	//fmt.Println(iniSlice)
}