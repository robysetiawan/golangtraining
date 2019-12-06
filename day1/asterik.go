package main

import "fmt"

func main() {
	var tinggi int
	fmt.Print("Input tinggi piramid : ")
	fmt.Scan(&tinggi)

	for i := 0; i < tinggi; i++ {
		for j := tinggi - 1; j > i; j-- {
			fmt.Print(" ")
		}

		for k := 0; k <= i; k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
