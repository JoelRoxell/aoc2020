package a10

import (
	"strconv"
	"testing"

	"github.com/joelroxell/aoc2020/utils"
)

func TestAdapterArray1(t *testing.T) {
	records := utils.ReadDat("./input.dat")
	adapters := make([]int, 0)

	for _, v := range records {
		n, err := strconv.Atoi(v)

		if err != nil {
			panic(err)
		}

		adapters = append(adapters, n)	
	}

	ones, threes, err := FindJoltStack(adapters)

	if err != nil {
		panic(err)
	}

	if ones != 75 || threes != 40 {
		t.Error("failed to count joltages")
	}
}

func TestFindCombinations(t *testing.T) {
	records := utils.ReadDat("./input.dat")
	adapters := make([]int, 0)

	for _, v := range records {
		n, err := strconv.Atoi(v)

		if err != nil {
			panic(err)
		}

		adapters = append(adapters, n)	
	}

	if FindCombinations(adapters) != 193434623148032 {
		t.Error("Failed to find combination count")
	}
}