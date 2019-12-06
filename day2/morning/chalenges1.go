package main

import "fmt"

func main() {
	arrayA := []string{"kazuya", "jin", "lee"}
	arrayB := []string{"kazuya", "feng", "kazuya", "jin"}
	arrayC := append(arrayA, arrayB...)

	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range arrayC {
		if _, exists := keys[entry]; !exists {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	fmt.Println(list)
}
