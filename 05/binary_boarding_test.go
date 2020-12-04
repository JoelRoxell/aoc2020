package a05

import (
	"fmt"
	"sort"
	"testing"

	"github.com/joelroxell/aoc2020/utils"
)

func TestBinaryBoardingDemo(t *testing.T) {
	actual := map[string]int {
		"BFFFBBFRRR": 567,
		"FFFBBBFRRR": 119,
		"BBFFBBFRLL": 820,
	}
	boardingPasses := []string { "BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL"	}
	highestID := 0

	for _, b := range boardingPasses {
		seatID := GetSeatID(GetRow(b), GetColumn(b))

		if (seatID > highestID) {
			highestID = seatID
		}

		if (actual[b] != seatID) {
			t.Errorf("%d is not %d", actual[b], seatID)
		}
	}

	if (highestID != 820) {
		t.Error("Did not find highest ID")
	}
}

func TestRowId(t *testing.T) {
	actual := GetSeatID(44, 5)
	expected := 357

	if (actual != expected) {
		t.Errorf("expected %d, not %d", expected, actual)	
	}
}

func TestGetRow(t *testing.T) {
	boarding := "FBFBBFFRLR"
	expected := 44
	actual := GetRow(boarding)
	
	if (actual != expected) {
		t.Errorf("expected %d, not %d", expected, actual)	
	}
}

func TestGetColumn(t *testing.T) {
	boarding := "FBFBBFFRLR"
	expected := 5
	actual := GetColumn(boarding)
	
	if (actual != expected) {
		t.Errorf("expected %d, not %d", expected, actual)	
	}
}

// func TestBinaryBoarding1(t *testing.T) {
// 	records := utils.ReadDat("./input.dat")
// 	highestID := 0

// 	for _, b := range records {
// 		if (len(b) != 10)  {
// 			continue
// 		}

// 		seatID := GetSeatID(GetRow(b), GetColumn(b))

// 		if (seatID > highestID) {
// 			highestID = seatID
// 		}
// 	}

// 	fmt.Println(highestID)
// }

func TestBinaryBoarding2(t *testing.T) {
	records := utils.ReadDat("./input.dat")
	idList := []int {}

	for _, b := range records {
		if (len(b) != 10)  {
			continue
		}

		seatID := GetSeatID(GetRow(b), GetColumn(b))

		idList = append(idList, seatID)
	}

	sort.Ints(idList)
	fmt.Println(len(idList))
	start := 51
	current := start

	for i := 0; i < len(idList); i++ {
		for (true) {
			if (idList[i] != current) {
				fmt.Printf("missing seat %d\n", current)
				current++
			} else {
				break
			}
		}

		current++
	}
}