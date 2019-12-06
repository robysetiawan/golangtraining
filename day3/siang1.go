package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	var sentence string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Input Text : ")
	if scanner.Scan() {
		sentence = scanner.Text()
	}
	// sentence := "lorem ipsum dolor sit amet, consectetur adipiscing elit. sed dolor neque, feugiat at ornare vel, eleifend vel lacus. curabitur mollis cursus dapibus. suspendisse vel enim eget nisi mattis aliquet"
	// sentence := "aaaz"
	frekuensi := make(map[string]int)
	var wg sync.WaitGroup
	for _, letter := range sentence {
		if _, exists := frekuensi[string(letter)]; !exists && letter >= 97 && letter <= 122 {
			frekuensi[string(letter)] = 0
		}
		if _, exists := frekuensi[string(letter)]; exists {
			wg.Add(1)
			go func() {
				// fmt.Println("ngitung huruf", string(letter))
				if strings.Contains(sentence, string(letter)) {
					frekuensi[string(letter)]++
					wg.Done()
				}
			}()
			wg.Wait()
		}
	}

	for huruf, banyak := range frekuensi {
		fmt.Printf("%v : %v\n", huruf, banyak)
	}
}
