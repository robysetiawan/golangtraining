package day2

//GetMinmax ...
func GetMinmax(numbers ...*int) (min, max int) {
	min = *numbers[0]
	max = *numbers[0]

	for _, value := range numbers {
		if *value > max {
			max = *value
		}
		if *value < min {
			min = *value
		}
	}
	return
}
