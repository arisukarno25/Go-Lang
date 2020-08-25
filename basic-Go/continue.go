package main

import "fmt"

func main() {
	// continue akan memberhentikan pada saat perulangan tertentu saja / saat kondisi tertentu
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
	fmt.Println("Perulangan ke-", i)
	}
}