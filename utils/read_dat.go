package utils

import (
	"bufio"
	"os"
)

// ReadDat reads a \n separated file and returns it's contents as and array
func ReadDat(file string) []string {
	var field []string
	var scanner *bufio.Scanner

	f, err := os.Open(file)

	defer f.Close()

	if err != nil {
			panic(err)
	}

	field = []string{}
	scanner = bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		field = append(field, scanner.Text())
	}

	return field
}


func Sum(arr []int) int {
	total := 0

	for _, v := range arr {
		total += v
	}

	return total
}