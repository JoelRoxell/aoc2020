package utils

import "math"

func Max(arr []int) (largest int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] > largest {
			largest = arr[i]
		}
	}

	return
}

func Min(arr []int) (smallest int) {
	smallest = math.MaxInt64

	for i := 0; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
		}
	}

	return
}

func Sum(arr []int) int {
	total := 0

	for _, v := range arr {
		total += v
	}

	return total
}
