package domain

import "math"

func getSeqNumber(n int) int {
	return int(math.Pow(2, float64(n)))
}
