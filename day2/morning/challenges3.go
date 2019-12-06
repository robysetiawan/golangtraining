package main

import (
	"fmt"
	"strings"
)

func main() {
	keys := "name,age"
	values := "Yovi,25"

	splitKeys := strings.Split(keys, ",")
	splitValues := strings.Split(values, ",")
	var result = map[string]string{}

	for index, key := range splitKeys {
		result[key] = splitValues[index]
	}
	fmt.Println(result)
}
