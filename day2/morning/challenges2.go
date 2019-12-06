package main

import "fmt"

func main() {
	arrayInput := [6]float64{3, 5, 7, 5, 3, 4}
	panjangArray := len(arrayInput)
	summary := 0.0
	for _, value := range arrayInput {
		summary += value
	}
	mean := summary / float64(panjangArray)

	var median float64
	if len(arrayInput)%2 == 0 {
		median = (arrayInput[panjangArray/2] + arrayInput[(panjangArray/2)-1]) / 2
	} else {
		median = arrayInput[(panjangArray / 2)]
	}
	fmt.Printf("mean : %v median %v", mean, median)
}
