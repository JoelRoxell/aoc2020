package a07

import (
	"strings"
	"testing"

	"github.com/joelroxell/aoc2020/utils"
)

func TestParseBag(t *testing.T) {
	bag1 := "light red bags contain 1 bright white bag, 2 muted yellow bags."
	// bag2 := "faded blue bags contain no other bags."

	bag := ParseBag(bag1)

	if strings.Compare(bag.name, "lightred") != 0 {
		t.Error("bag name is incorrect")
	}

	if (len(bag.contains) != 2) {
		t.Error("bag contain length is incorrect")
	}
}

func TestParseContainingBags(t *testing.T) {
	child := "1 bright white bag"

	name, count, _ := ParseContainingBag(child)

	if strings.Compare(name, "brightwhite") != 0 {
		t.Error("Child bag name is wrong")
	}

	if count != 1 {
		t.Error("The child count is off")
	}
}

func TestDemo(t *testing.T) {
	records := utils.ReadDat("./demo.dat")

	for _, line := range records {
		ParseBag(line)
	}

	Reverse("shinygold", nil)

	resultList := make([]string, 0)

	for _, end := range Endings {
		resultList = Concat(resultList, CollectPath(end))
	}

	resultList = Filter("shinygold", resultList)
	resultList = Unique(resultList)

	if len(resultList) != 4 {
		t.Error("result list is not valid")
	}
}

func Test1(t *testing.T) {
	records := utils.ReadDat("./input.dat")

	for _, line := range records {
		ParseBag(line)
	}

	Reverse("shinygold", nil)

	resultList := make([]string, 0)

	for _, end := range Endings {
		resultList = Concat(resultList, CollectPath(end))
	}

	resultList = Filter("shinygold", resultList)
	resultList = Unique(resultList)

	if len(resultList) != 241 {
		t.Error("result list is not valid")
	}
}
func Test2Demo(t *testing.T) {
	records := utils.ReadDat("./demo.dat")

	for _, line := range records {
		ParseBag(line)
	}

	count := BagValue("shinygold")

	if count != 32 {
		t.Error("Invalid forward count")
	}
}

func Test2(t *testing.T) {
	records := utils.ReadDat("./input.dat")

	for _, line := range records {
		ParseBag(line)
	}

	count := BagValue("shinygold")

	if count != 82930 {
		t.Error("Invalid forward count")
	}
}