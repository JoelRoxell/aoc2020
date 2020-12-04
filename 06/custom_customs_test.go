package a06

import (
	"testing"

	"github.com/joelroxell/aoc2020/utils"
)

func TestDemo(t *testing.T) {
	groups := ReadGroups("./demo.dat")
	resultList := make([]int, 0)

	for _, group := range groups {
		answers := ""

		for _, record := range group {
			answers += record
		}

		uniqueAnswers := Unique(answers)
		resultList = append(resultList, len(uniqueAnswers))
		answers = ""
	}

	if utils.Sum(resultList) != 11 {
		t.Error("sum is not 11")
	}
}

func TestDemo2(t *testing.T) {
	groups := ReadGroups("./input.dat")
	counts := make([]int, 0)

	for _, group := range groups {
		count := Includes(group)

		counts = append(counts, count)
	}

	if utils.Sum(counts) != 3288 {
		t.Error("Didn't find the correct sum")
	}
}

func TestUnique(t *testing.T) {
	test1 := "aaaa"
	test2 := "abcd"
	test3 := "abbc"

	a := Unique(test1)
	b := Unique(test2)
	c := Unique(test3)

	if len(a) != 1 {
		t.Error("length err")
	}

	if len(b) != 4 {
		t.Error("length err")
	}

	if len(c) != 3 {
		t.Error("length err")
	}
}
