package main

import "fmt"

func main() {
	var N int

	fmt.Print("Input N : ")
	fmt.Scan(&N)

	for i := 1; i <= N/2; i++ {
		if N%i == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println(N)
}
