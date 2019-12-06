package main

import "fmt"

func main() {
	var a, t, luas float32

	fmt.Print("Masukkan nilai alas : ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai tinggi : ")
	fmt.Scan(&t)

	luas = (a * t) / 2
	fmt.Printf("Luas segitiga dengan alas %v dan tinggi %v adalah %v", a, t, luas)
}
