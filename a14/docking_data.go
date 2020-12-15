package a14

import (
	"bufio"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Operation struct {
	value uint
	pos uint
}

func DockingData(s *bufio.Scanner) int {
	var andBits uint
	var clearBits uint
	memory := make(map[uint]uint)

	for s.Scan() {
		data := s.Text()
		op, newMask := ParseRecord(data)

		if newMask != "" {
			andBits, clearBits = ParseMask(newMask)		

			continue
		}

		memory[op.pos] = applyMask(op.value, andBits, clearBits)
	}	

	total := 0

	for _, v := range memory {
		total += int(v)
	}

	return total
}

func DockingData2(s *bufio.Scanner) int {
	memory := make(map[uint]uint)
	mask := ""

	for s.Scan() {
		data := s.Text()
		op, newMask := ParseRecord(data)

		if newMask != "" {
			mask = newMask

			continue
		}
		format := "%036b"

		binaryValueAsString := fmt.Sprintf(format, op.pos)
		resultMask := applyMask2(binaryValueAsString, mask)
		combinations := generateCombinationsFrom(resultMask, format)
		createVariants(resultMask, combinations)
	
		for _, s := range combinations {
			address := stringBinaryToDec(s)
		
			memory[address] = op.value
		}
	}	

	total := 0

	for _, v := range memory {
		total += int(v)
	}


	return total
} 

func applyMask(value uint, andBits uint, clearBits uint) uint {
	value |= andBits	
	value &^= clearBits

	return value
}

func applyMask2(value string, mask string) string {
	result := ""

	for i := 0; i < len(mask); i++ {
		maskBit := mask[i]
		valueBit := value[i]
	
		if maskBit == '1' {
			result += "1"
		} else if (maskBit == 'X') {
			result += string(maskBit)
		} else if (maskBit == '0') {
			result += string(valueBit)
		}
	}

	return result
}

func parseStringToUint(s string) uint {
	n, err := strconv.ParseUint(s, 10, 32)

	if err != nil {
		panic(err)
	}	

	return uint(n)
}

func ParseRecord(s string) (op Operation, mask string) {
	maskReg := regexp.MustCompile(`[X|1|0]{36}`)
	numReg := regexp.MustCompile(`\d+`)

	mask = maskReg.FindString(s)

	if mask != "" {
		return
	}
	
	memArr := strings.Split(s, "=")
	memPos := numReg.FindString(memArr[0])
	op.pos = parseStringToUint(memPos)
	op.value = parseStringToUint(numReg.FindString(memArr[1]))

	return
}

func ParseMask(mask string) (andBits uint, clearBits uint) {
	AndSymbol := '1'
	ClearSymbol := '0'

	for i := 0; i < len(mask); i++ {
		pos := float64(35 - i)
		decValue := math.Pow(2, pos)
		value := uint(decValue)

		switch rune(mask[i]) {
		case AndSymbol:
			andBits += value
		case ClearSymbol:
			clearBits += value
		}	
	}

	return
}

func generateCombinationsFrom(floatingString string, format string) []string {
	floatingCount := 0

	for _, c := range floatingString {
		if c == 'X' {
			floatingCount++	
		}
	}

	combinations := make([]uint, int(math.Pow(2, float64(floatingCount))))
	stringCombos := make([]string, 0)

	for i := 0; i < len(combinations); i++ {
		bytesAsString := fmt.Sprintf(format, i)
		stringCombos = append(stringCombos, bytesAsString)
	}

	return stringCombos
}

func createVariants(maskResult string, combinations []string) {
	for i := len(maskResult) - 1; i > -1; i-- {
		c := maskResult[i]

		if c == '1' {
			for j := 0; j < len(combinations); j++ {
				combo := &combinations[j]

				insertAt("1", i, combo)
			}
		}
		if c == '0' {
			for j := 0; j < len(combinations); j++ {
				combo := &combinations[j]

			 	insertAt("0", i, combo)
			}
		}
	}
}

func insertAt(n string, at int, set *string) {
	start := (*set)[:at + 1] 
	end := (*set)[at + 1:]
	adjusted := start + n + end
	keep := len(adjusted) - len(*set)

	*set = adjusted[keep:]
}

func stringBinaryToDec(s string) uint {
	total := uint(0)
	sigBitPos := len(s) - 1

	for i, v := range s {
		if v == '0' {
			continue
		}

		pos := math.Abs(float64(i - sigBitPos))
		decValue := math.Pow(2, pos)
		total += uint(decValue)
	}

	return total
}