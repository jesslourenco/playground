package main

import (
	"fmt"
	"sort"
)

// Interval scheduling problem: minimize number of weddings to attend
func main() {
	weddingsC := [][]int{
		{1, 2}, {3, 8}, {1, 5}, {7, 8}, {3, 9},
	}

	sort.Slice(weddingsC, func(i, j int) bool {
		return weddingsC[i][0] > weddingsC[j][0]
	})

	S := []int{}
	lastWedding := []int{0, 0} // assuming any time will be > 0

	for _, wedding := range weddingsC {
		if len(S) == 0 || wedding[1] < lastWedding[0] {
			S = append(S, wedding[0])
			lastWedding = wedding
		}
	}

	for _, d := range S {
		fmt.Printf("%d ", d)
	}
}
