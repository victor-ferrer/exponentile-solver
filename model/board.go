package model

import (
	"fmt"
	"math"
	"math/rand"
)

type Board [][]int
type Tile struct {
	X, Y int
}

type ExponenTile int

func GetSeqNumber(n int) int {
	return int(math.Pow(2, float64(n)))
}

const width = 4

func Initialize() Board {

	b := make([][]int, width)

	for rows := 0; rows < width; rows++ {
		row := make([]int, width)
		for cols := 0; cols < width; cols++ {
			row[cols] = GetSeqNumber(rand.Intn(5) + 1)
		}
		b[rows] = row
	}
	return b
}

func (b Board) Render() {
	for row := 0; row < width; row++ {
		for column := 0; column < width; column++ {
			fmt.Printf("%02d ", b[row][column])
		}
		fmt.Println()
	}
}
