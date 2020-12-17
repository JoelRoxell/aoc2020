package a11

import (
	"fmt"
	"testing"
)


func TestLoadField(t *testing.T) {
	field := LoadField("./demo.dat")

	if len(field) != 10 {
		t.Error("field to parse field")
	}
}

func TestPassFieldDemo(t *testing.T) {
	field := LoadField("./demo.dat")
	terminationN := 3

	fieldSum := 0
	nextField := PassField(field, 4, false)

	for {
		nextField = PassField(nextField, 4, false)
		nextFieldSum := SumField(nextField)

		if nextFieldSum == fieldSum {
			terminationN--
		}

		if terminationN == 0 {
			break
		}

		fieldSum = nextFieldSum
	}

	if fieldSum != 37 {
		t.Error("failed to calculate occupy sum")
	}
}


func TestPassField(t *testing.T) {
	field := LoadField("./input.dat")
	terminationN := 3

	fieldSum := 0
	nextField := PassField(field, 4, false)

	for {
		nextField = PassField(nextField, 4, false)
		nextFieldSum := SumField(nextField)

		if nextFieldSum == fieldSum {
			terminationN--
		}

		if terminationN == 0 {
			break
		}

		fieldSum = nextFieldSum
	}

	if fieldSum != 2126 {
		t.Error("failed to calculate occupy sum")
	}
}


func TestPassFieldDemo2(t *testing.T) {
	nextField := LoadField("./demo.dat")
	fieldSum := 0
	threshold := 10

	for {
		nextField = PassField(nextField, 5, true)
		nextFieldSum := SumField(nextField)

		if nextFieldSum == fieldSum {
			threshold--
		}

		if threshold == 0 {
			break
		}

		fieldSum = nextFieldSum
	}

	if fieldSum != 26 {
		t.Error("failed to calculate occupy sum")
	}
}

func TestPassField2(t *testing.T) {
	nextField := LoadField("./input.dat")
	fieldSum := 0
	threshold := 3

	for {
		nextField = PassField(nextField, 5, true)
		nextFieldSum := SumField(nextField)

		if nextFieldSum == fieldSum {
			threshold--
		}

		if threshold == 0 {
			break
		}

		fieldSum = nextFieldSum
	}

	fmt.Println(fieldSum)

	if fieldSum != 1914 {
		t.Error("failed to calculate occupy sum")
	}
}