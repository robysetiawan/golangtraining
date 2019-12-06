package main

import (
	"fmt"
	"reflect"
)

func main() {
	even := [5]int{1: 2, 2: 4}

	fmt.Println(reflect.ValueOf(even).Kind())
	fmt.Println(len(even))
}
