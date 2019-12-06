package main

import (
	"fmt"
	"sync"
)

func main() {
	sentence := "lorem ipsum dolor sit amet, consectetur adipiscing elit. sed dolor neque, feugiat at ornare vel, eleifend vel lacus. curabitur mollis cursus dapibus. suspendisse vel enim eget nisi mattis aliquet"
	frekuensi := make(map[string]int)
	var wg sync.WaitGroup
	for _, letter := range sentence {
		if _, exists := frekuensi[string(letter)]; !exists && letter >= 97 && letter <= 122 {
			frekuensi[string(letter)] = 0
		}
	}
	for huruf, _ := range frekuensi {
		wg.Add(len(frekuensi))
		go func() {
			for _, letter := range sentence {
				if string(letter) == huruf {
					fmt.Println("ngitung huruf", huruf)
					frekuensi[huruf]++
					wg.Done()
				}
			}
		}()
		wg.Wait()
	}

	for huruf, banyak := range frekuensi {
		fmt.Printf("%v : %v\n", huruf, banyak)
	}
}
