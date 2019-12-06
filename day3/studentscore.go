package day3

//Student ...
type Student struct {
	Name  []string
	Score []float64
}

//Summary ...
func (s Student) Summary() (min, max, average float64) {
	sum := 0.0
	min = s.Score[0]
	max = s.Score[0]
	for _, score := range s.Score {
		sum += score
		if score < min {
			min = score
		}
		if score > max {
			max = score
		}
	}

	average = sum / float64(len(s.Score))
	return
}
