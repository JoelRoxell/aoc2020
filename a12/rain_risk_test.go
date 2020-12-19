package a12

import (
	"math"
	"testing"

	"github.com/joelroxell/aoc2020/utils"
)


func TestRainRisk(t * testing.T) {
	state := &State{
		location: [2]int{0, 0},
		facing: EAST,
	}
	s, f :=	utils.CreateScanner("./input.dat")	

	defer f.Close()

	for {
		lat, dir, err := ParseStep(s)		

		if err != nil {
			break
		}

		state = UpdateState(state, dir, lat)	
	}

	result := math.Abs(float64(state.location[0])) + math.Abs(float64(state.location[1]))

	if result != 582 {
		t.Error("Faild to find correct manhattan distance")
	}
}

func TestRainRisk2(t * testing.T) {
	state := &State{
		location: [2]int{0, 0},
		facing: EAST,
	}
	s, f :=	utils.CreateScanner("./input.dat")	

	defer f.Close()

	for {
		lat, dir, err := ParseStep(s)		

		if err != nil {
			break
		}

		state = UpdateState(state, dir, lat)	
	}

	result := math.Abs(float64(state.location[0])) + math.Abs(float64(state.location[1]))

	if result != 582 {
		t.Error("Faild to find correct manhattan distance")
	}
}

func TestNextLatitude(t *testing.T) {
	l := EAST
	
	w := NextLatitude(l, 180)

	if w != WEST {
		t.Error("faild to turn west")
	}

	w = NextLatitude(w, 360)

	if w != WEST {
		t.Error("faild to turn west")
	}

	n := NextLatitude(w, 90)

	if n != NORTH {
		t.Error("faild to turn north")
	}

	w = NextLatitude(n, -90)

	if w != WEST {
		t.Error("faild to turn north")
	}
}