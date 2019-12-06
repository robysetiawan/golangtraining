package main

import (
	"fmt"
	"reflect"
)

func main() {
	primes := []int{2, 3, 5}
	kata := "halo"

	for index := 0; index < len(primes); index++ {
		fmt.Print(primes[index])
	}
	fmt.Println(reflect.ValueOf(kata).Kind())
	for index, elem := range primes {
		fmt.Println(index, "=>", elem)
	}

}
