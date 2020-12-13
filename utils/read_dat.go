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

func CreateScanner(file string) (*bufio.Scanner, *os.File){
	var scanner *bufio.Scanner

	f, err := os.Open(file)

	if err != nil {
			panic(err)
	}

	scanner = bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	return scanner, f
}