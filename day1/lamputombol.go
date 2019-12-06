package main

import "fmt"

func lamputombol(N int) string {
	count := 0
	for i := 1; i <= N; i++ {
		if N%i == 0 {
			count++
		}
	}
	if count%2 == 0 {
		return "lampu mati"
	} else {
		return "lampu menyala"
	}
}

func main() {
	for i := 1; i < 30; i++ {
		a := lamputombol(i)
		fmt.Println(i, a)
	}
}
