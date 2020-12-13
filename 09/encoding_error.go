package a09

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/joelroxell/aoc2020/utils"
)

func PreambleCheck(s *bufio.Scanner, size int) (res int, err error) {
	res = 0
	preamble := make([]int, size)
	i := 0

	for s.Scan() {
		data := s.Text()
		num, convertErr := strconv.Atoi(data)
		
		if convertErr != nil {
			err = convertErr

			return
		}

		if (i < size) {
			// Skip validation until premable is filled with initial values
			preamble = Shift(preamble, num)	
			i++

			continue
		}

		// validate next num
		_, _, complimentErr := FindCompliment(preamble, num)

		if complimentErr != nil {
			err = complimentErr

			return
		}

		preamble = Shift(preamble, num)	
		i++
	}

	return 
}

func FindWeakness(s *bufio.Scanner, n int) (smallest int, largest int, err error) {
	seen := make([]int, 0)

	for s.Scan() {
		data := s.Text()
		num, convertErr := strconv.Atoi(data)
		
		if convertErr != nil {
			err = convertErr

			return
		}

		seen = append(seen, num)

		if (len(seen) < 2) {
			// Skip validation until premable is filled with initial values
			continue
		}

		if utils.Sum(seen) < n {
			continue
		}

		stack := seen

		for i := 0; i < len(stack); i++ {
			nextStack := stack[i:]
			nextSum := utils.Sum(nextStack)	

			if nextSum == n {
				largest = utils.Max(nextStack)
				smallest = utils.Min(nextStack)

				return
			} else if (nextSum < n) {
				break
			}
		}
	}

	return 
}

func FindCompliment(arr []int, num int) (n1 int, n2 int, err error) {
	seen := make(map[int]int)
	terminate := false

	for _, v := range arr {
		if terminate {
			break
		}

		if seen[v] != 0 {
			n1 = seen[v]
			n2 = v
			terminate = true
		} else {
			res := num - v
			seen[res] = v
		}
	}

	if !terminate {
		err = fmt.Errorf("Found no sum match for %d", num)
	}

	return
}

func Shift(arr []int, value int) []int {
	return append([]int{value}, arr[:len(arr) - 1]...)
}
