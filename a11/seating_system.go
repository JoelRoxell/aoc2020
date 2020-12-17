package a11

import (
	"math"

	"github.com/joelroxell/aoc2020/utils"
)

const (
	FLOOR = iota
	FREE
	OCCUPIED
)

func LoadField(name string) [][]int {
	row := 0
	field := make([][]int, 0)
	s, f := utils.CreateScanner(name)

	defer f.Close()

	for s.Scan() {
		record := s.Text()	
		size := len(record)
		field = append(field, make([]int, size))

		for i, c  := range record {
			switch c {
			case 'L': 
				field[row][i] = FREE
			case '.' :
				field[row][i] = FLOOR
			}
		}

		row++
	}

	return field
}

func GetAdjesentCount(pos [2]int, field [][]int, los bool) int {
	x := pos[0]
	y := pos[1]
	count := 0
	fieldHeight := len(field) - 1
	fieldWidth := len(field[0]) - 1
	MAX := int(math.Max(float64(len(field[0])), float64(len(field))))
	directions := [][]int{
		{-1, -1}, 
		{0, -1}, 
		{1, -1}, 
		{-1, 0}, 
		{1, 0}, 
		{-1, 1}, 
		{0, 1}, 
		{1, 1},
	} 

	for _, neighbour := range directions {
		var nX int
		var nY int

		for i := 1; i <= MAX; i++ {
			if los {
				nX = x + neighbour[0] * i
				nY = y + neighbour[1] * i
			} else {
				nX = x + neighbour[0]
				nY = y + neighbour[1]
			}

			outOfBounds := nX < 0 || nX > fieldWidth || nY < 0 || nY > fieldHeight 

			if (outOfBounds || field[nY][nX] == FREE) {
				break
			}
	
			if field[nY][nX] == OCCUPIED {
				count++

				break
			}
		}
	}

	return count
}

func PassField(field [][]int, maxAdjecent int, los bool) [][]int {
	width := len(field[0])
	height := len(field)
	nextField := make([][]int, height)
	
	for y, row := range field {
		nextField[y] = make([]int, width)

		for x, positionValue := range row {
			if positionValue == FLOOR {
				nextField[y][x] = field[y][x]

				continue
			}

			count := GetAdjesentCount([2]int{x, y}, field, los)

			if count >= maxAdjecent {
				nextField[y][x] = FREE
			} else if count == 0 && positionValue == FREE {
				nextField[y][x] = OCCUPIED
			} else {
				nextField[y][x] = field[y][x]
			}
		}
	}

	return nextField
}

func SumField(field [][]int) int {
	total := 0

	for _, row := range field {
		for _, v := range row {
			if v == OCCUPIED {
				total++
			}
		}
	}

	return total
}
// return occupied seats 