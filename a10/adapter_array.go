package a10

import (
	"fmt"
	"sort"
)

func FindJoltStack(adapterRatings []int) (ones int, threes int, err error) {
	withInitialSocket := append(adapterRatings, 0)
	sort.Ints(withInitialSocket)
	threes = 1

	for i := 1; i < len(withInitialSocket); i++ {
		previous := withInitialSocket[i - 1]
		current := withInitialSocket[i]

		diff := current - previous

		if diff == 1 {
			ones++
		} else if diff == 3 {
			threes++
		} else if diff > 3 {
			err = fmt.Errorf("previous adapter (%d) is not compatilbe with the current one (%d)", previous, current)

		} else if (diff == 0) {
			fmt.Println("0000")
		}
	}

	return	
}

var Cache map[int]int = make(map[int]int)

func FindCombinations(adapters []int) int {
	adapters = append(adapters, 0)
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters) - 1] + 3)
	ends := FindEndings(0, adapters)

	return ends
}

func FindEndings(n int, adapters []int) int {
	if Cache[n] != 0 {
		return Cache[n]
	}

	children := FindChildren(n, adapters)
	endings := 0

	for _, child := range children {
		endings += FindEndings(child, adapters)
	}

	if len(children) == 0 {
		endings++
	}

	Cache[n] = endings

	return endings
}

func FindChildren(n int, adapters []int) []int {
	childrens := make([]int, 0)

	for _, adapter := range adapters {
		diff := adapter - n		

		if diff > 0 && diff < 4 {
			childrens = append(childrens, adapter)
		}
	}

	return childrens
}
