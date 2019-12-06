package main

import (
	"fmt"
)

func main() {
	sampleInput := [6]int{2, 1, 3, 1, 1, 2}
	// sampleInput := [4]int{1, 1, 2, 2}

	for i := 0; i < len(sampleInput); i++ {
		if sampleInput[i] == 1 {
			for j := i + 1; j < len(sampleInput); j++ {
				if sampleInput[j] == 2 {
					sampleInput[i+1], sampleInput[j] = sampleInput[j], sampleInput[i+1]
					i = j + 1
					break
				}
			}
		} else if sampleInput[i] == 2 {
			for j := i + 1; j < len(sampleInput); j++ {
				if sampleInput[j] == 1 {
					sampleInput[j+1], sampleInput[i] = sampleInput[i], sampleInput[j+1]
					i = j + 1
					break
				}
			}
		}
	}

	fmt.Println(sampleInput)
}
