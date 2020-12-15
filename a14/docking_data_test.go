package a14

import (
	"strings"
	"testing"

	"github.com/joelroxell/aoc2020/utils"
)

func TestDockingData1(t *testing.T) {
	s, f := utils.CreateScanner("./input.dat")

	defer f.Close()

	result := DockingData(s)

	if result != 8332632930672 {
		t.Error("failed ot calculate eding values in memory")
	}
}

func TestDockingData2(t *testing.T) {
	s, f := utils.CreateScanner("./input.dat")

	defer f.Close()

	result := DockingData2(s)

	if result != 4753238784664 {
		t.Error("failed to calculate proper memory sum")
	}
}

func TestParseRecord(t *testing.T) {
	op, mask := ParseRecord("mem[7] = 101")

	if op.pos != 7 || op.value != 101 {
		t.Error("the operation struct should pos = 7 and val = 101")
	}

	op, mask = ParseRecord("mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")

	if strings.Compare(mask, "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X") != 0 {
		t.Error("Failed to parse mask from string")
	}
}

func TestParseMask(t *testing.T) {	
	mask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"
	andBits, clearBits := ParseMask(mask)
	
	if andBits != 64 {
		t.Errorf("invalid and value")
	}

	if clearBits != 2 {
		t.Errorf("invalid clear value")
	}
}


func TestVariants(t *testing.T) {
	variantString := "0001X0XX"
	expected := []uint{16,17,18,19,24,25,26,27}
	stringCombos := generateCombinationsFrom(variantString, "%08b")
	createVariants(variantString, stringCombos)

	for i, s := range stringCombos {
		v := stringBinaryToDec(s)

		if v != expected[i] {
			t.Errorf("fild to determine %d -> %d", v, expected[i])
		}
	}
}

func TestInsertAt(t *testing.T) {
	x := "00000111"

	insertAt("0", 5, &x)

	if strings.Compare(x, "00001011") != 0 {
		t.Errorf("faild to shift part of bits left")
	}
}
