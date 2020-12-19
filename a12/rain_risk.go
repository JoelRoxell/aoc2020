package a12

import (
	"bufio"
	"fmt"
	"math"
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
type State2 struct {
	location [2]int
	waypoint [2]int
}
type vector [2]int


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

func UpdateState2(s *State2, dir *DirectionAction, lat *LatitudeAction) *State2 {
	if dir != nil {
		switch dir.action {
		case LEFT:
			s.waypoint = Rotate(s.waypoint, -dir.value)	
		case RIGHT:
			s.waypoint = Rotate(s.waypoint, dir.value)	
		case FORWARD:
			nMove := [2]int{
				s.waypoint[0] * dir.value,
				s.waypoint[1] * dir.value,
			}

			s.location[0] += nMove[0]
			s.location[1] += nMove[1]
		}
	} else {
		s.waypoint = CalibrateWaypoint(s.waypoint, lat.action, lat.value)
	}


	return s
}

func Move(ship vector, waypoint vector, sectors int) vector {
	xPos := ship[0] + waypoint[0] * sectors
	yPos := ship[1] + waypoint[1] * sectors
	
	return vector{xPos, yPos}
}

func CalibrateWaypoint(wp vector, l Latitude, value int) vector {
	v := GetSymbol(value, l)
	nWp := vector{
		wp[0],
		wp[1],
	}

	if l == NORTH || l == SOUTH {
		nWp[1] += v
	} else {
		nWp[0] += v
	}

	return nWp
}

func GetLatitudeIndex(l Latitude) int {
	if l == NORTH || l == SOUTH {
		return 1
	} 

	return 0
}

func GetSymbol(val int, l Latitude) int {
	symbol := 1

	if SOUTH == l || WEST == l {
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

func Rotate(wp [2]int, deg int) [2]int {
	rad := float64(deg) * (math.Pi / 180);
	w := [][]int{
		{wp[0]},
		{wp[1]},
	}
	T := [][]int{
		{int(math.Cos(rad)), int(math.Sin(rad))},
		{int(-math.Sin(rad)), int(math.Cos(rad))},
	}
	m := multiply(T, w)

	// for _, v := range T {
	// 	fmt.Println(v)
	// }

	return [2]int{m[0][0],m[1][0]}
}

func multiply(x, y [][]int) [][]int {
	if len(x[0]) != len(y) {
		panic("can't do matrix op")
	}

	out := make([][]int, len(x))

	for i := 0; i < len(x); i++ {
		out[i] = make([]int, len(y[0]))

		for j := 0; j < len(y[0]); j++ {
			for k := 0; k < len(y); k++ {
				out[i][j] += x[i][k] * y[k][j]
			}
		}
	}
	
	return out
}
