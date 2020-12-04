package utils

import (
	"testing"
)

func TestReadDat(t *testing.T) {
	datFile := "./test.dat"
	records := ReadDat(datFile)
	lineCount := 3

	if (len(records) != lineCount) {
		t.Errorf("line count should be %d", lineCount)
	}
}