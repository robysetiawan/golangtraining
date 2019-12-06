package main

import "fmt"

func prima(x int) string {
	if x < 2 {
		return "bukan bilangan prima"
	}
	for i := 2; i <= x/2; i++ {
		if x%i == 0 {
			return "bukan bilangan prima"
		}
	}
	return "bilangan prima"
}

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(prima(a))
}
