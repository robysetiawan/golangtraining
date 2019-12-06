package main

import "fmt"

func main() {
	kata := "halo"
	for _, char := range kata {
		fmt.Println(string(char))
	}
}
