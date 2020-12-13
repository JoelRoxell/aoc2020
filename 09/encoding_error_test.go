package a09

import (
	"fmt"
	"testing"

	"github.com/joelroxell/aoc2020/utils"
)

func TestPreambleCheckDemo(t *testing.T) {
	s, f := utils.CreateScanner("./demo.dat")

	defer f.Close()

	_, err := PreambleCheck(s, 5)

	fmt.Println(err)
}

func TestPreambleCheck1(t *testing.T) {
	s, f := utils.CreateScanner("./input.dat")

	defer f.Close()

	_, err := PreambleCheck(s, 25)

	fmt.Println(err)
}

func TestFindWeaknessDemo(t *testing.T) {
	s, f := utils.CreateScanner("./demo.dat")

	defer f.Close()

	sm, lg, _ := FindWeakness(s, 127)

	if sm + lg != 62 {
		t.Error("failed ot calucalte weakness sum")
	}
}


func TestFindWeakness2(t *testing.T) {
	s, f := utils.CreateScanner("./input.dat")

	defer f.Close()

	sm, lg, _ := FindWeakness(s, 1639024365)

	if sm + lg != 219202240 {
		t.Error("failed ot calucalte weakness sum")
	}
}