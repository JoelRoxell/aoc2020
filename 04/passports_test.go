package a04

import (
	"testing"
)


func TestPassportInputDemo(t *testing.T) {
	records := ParsePassports("./test.dat")
	actual := 4

	if (len(records) != actual) {
		t.Errorf("There should only be %d objects not %d", actual, len(records))
	}
}

func TestPassportValidator(t *testing.T) {
	validObj := Object{
		"ecl": "gry",
		"pid": 860033327,
		"eyr": 2020, 
		"hcl": "#fffffd",
		"byr": 1937,
		"iyr": 2017,
		"hgt": "183cm",
	}
	
	isValid := IsValidPassport(validObj)

	if (!isValid) {
		t.Error("should be valid")
	}
}

func TestPassport1(t *testing.T) {
	records := ParsePassports("./input.dat")
	validPassports := 0

	for _, record := range records {
		if IsValidPassport(record) {
			validPassports++
		}
	}

	if (validPassports != 264) {
		t.Error("should be 264")
	}
}

func TestPassport2(t *testing.T) {
	records := ParsePassports("./input.dat")
	validPassports := 0

	for _, record := range records {
		isValid := IsValidPassport2(record) 

		if (isValid) {
			validPassports++
		}
	}

	if validPassports != 224 {
		t.Errorf("should be 224 not %d", validPassports)
	}
}
