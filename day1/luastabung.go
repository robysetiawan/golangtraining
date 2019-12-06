package main

import (
	"fmt"
)

func main() {
	var r, t, luas float32
	const phi = 3.14

	fmt.Print("Masukkan nilai tinggi tabung : ")
	fmt.Scan(&t)
	fmt.Print("Masukkan nilai jari-jari tabung : ")
	fmt.Scan(&r)

	luas = 2 * phi * r * (r + t)
	fmt.Printf("Luas tabung dengan jari-jari %v dan tinggi %v adalah %v", r, t, luas)
}
