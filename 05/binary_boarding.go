package a05

import (
	"math"
)

func GetSeatID(row int, column int) int {
	return row * 8 + column
}

func GetRow(boarding string) int {
	return SplitNGet(boarding[:7], 0, 127, "B", "F")
}

func GetColumn(boarding string) int {
	seatString := boarding[7:]

	return SplitNGet(seatString, 0, 7, "R", "L")
}

func SplitNGet(selection string, start int, end int, upper string, lower string) int {
	currentSelection := []int{start, end}
	
	for _, c := range selection {
		balance := mid(sum(currentSelection))

		if (string(c) == lower) {
			upper := int(math.Floor(balance))
			currentSelection = []int{currentSelection[0], upper}
		} else if (string(c) == upper) {
			lower := int(math.Ceil(balance))
			currentSelection = []int{lower, currentSelection[1]}
		} else {
			panic("Invalid boarding pass")
		}
	}

	if (currentSelection[0] != currentSelection[1]) {
		panic("The passport is not in balance")
	}

	return currentSelection[0]
}

func mid(num int) float64 {
	return float64(num) / 2
}

func sum(arr []int) int {
	total := 0

	for _, v := range arr {
		total += v
	}

	return total
}