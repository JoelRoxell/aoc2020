package a12

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Direction int
type Latitude int
type State struct {
	location [2]int
	facing Latitude
}

const (
	NORTH Latitude = iota
	EAST
	SOUTH
	WEST
)
const (
	LEFT Direction = iota - 1
	FORWARD
	RIGHT
)

type LatitudeAction struct {
	action Latitude
	value int
}
type DirectionAction struct {
	action Direction
	value int
} 

var nReg *regexp.Regexp = regexp.MustCompile(`\d+`)
var actionReg  *regexp.Regexp = regexp.MustCompile(`[A-Z]`)

func ParseStep(s *bufio.Scanner) (lat *LatitudeAction, dir *DirectionAction, err error){
	if !s.Scan() {
		err = fmt.Errorf("no more actions to read")

		return
	}

	data := s.Text()
	action := actionReg.FindString(data)
	n, err := strconv.Atoi(nReg.FindString(data))

	if err != nil {
		panic(err)
	}

	if strings.ContainsAny(action, "NESW") {
		lat = &LatitudeAction {
			action: LatitudeStringToIota(action),
			value: n,
		}
	} else {
		dir = &DirectionAction {
			action: DirectionStringToIota(action),
			value: n,
		}
	}

	return
}

func LatitudeStringToIota(lat string) Latitude {
	latRune := lat[0]

	switch latRune {
	case 'N':
		return NORTH
	case 'E':
		return EAST
	case 'S':
		return SOUTH
	case 'W':
		return WEST
	}

	panic("couldn't parse latitude action")
}

func DirectionStringToIota(dir string) Direction {
	dirRune := dir[0]

	switch dirRune {
	case 'L':
		return LEFT
	case 'F':
		return FORWARD
	case 'R':
		return RIGHT
	}

	panic("couldn't parse direction action")
}



func UpdateState(state *State, dir * DirectionAction, lat *LatitudeAction) *State {
	if dir != nil {
		if dir.action == FORWARD {
			state.location[GetLatitudeIndex(state.facing)] += GetSymbol(dir.value, state.facing)
		} else {
			newLatitude := NextLatitude(state.facing, dir.value * int(dir.action))
			state.facing = newLatitude
		}
	} else {
		state.location[GetLatitudeIndex(lat.action)] += GetSymbol(lat.value, lat.action)
	}

	return state
}

func GetLatitudeIndex(l Latitude) int {
	if l == NORTH || l == SOUTH {
		return 1
	} 

	return 0
}

func GetSymbol(val int, l Latitude) int {
	symbol := 1

	if NORTH == l || WEST == l {
		symbol = -1
	}

	return symbol * val
}

func NextLatitude(l Latitude, degrees int) Latitude {
	latDeg := []int{0, 90, 180, 270}
	n := (360 + latDeg[l] + degrees) % 360
	
	return Latitude(IndexOf(n, latDeg))
}

func IndexOf(n int, arr []int) int {
	for i, v := range arr {
		if v == n {
			return i
		}
	}

	return  -1
}