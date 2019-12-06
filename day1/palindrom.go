package main

import (
	"bufio"
	"fmt"
	"os"
)

func ispalindrom(kata string) string {
	panjang := len(kata)
	for i := 0; i < panjang/2; i++ {
		if kata[i] != kata[panjang-1-i] {
			return "Bukan palindrom"
		}
	}
	return "Palindrom"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		kata := scanner.Text()
		fmt.Println(ispalindrom(kata))
	}
}
