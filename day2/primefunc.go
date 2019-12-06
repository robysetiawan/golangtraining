package day2

import "math"

//GetPrime ...
func GetPrime(x int) int {
	isPrime := func(num int) bool {
		if num < 2 {
			return false
		}
		for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}

	j := 0
	var i int = 1
	for j < x {
		if isPrime(i) {
			j++
		}
		if j == x {
			return i
		}

		i++
	}
	return i
}
