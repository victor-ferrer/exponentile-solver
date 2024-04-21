package solver

import "math"

func GetSeqNumber(n int) int {
	return int(math.Pow(2, float64(n)))
}
