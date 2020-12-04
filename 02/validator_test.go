package a02

import (
	"testing"

	"github.com/joelroxell/aoc2020/utils"
)

func TestReadPolicy(t *testing.T) {
	policy := "1-3 a"
	atMin, atMax, char := readPolicy(&policy)
	
	if (atMin != 1 && atMax != 3 && char != "a") {
		t.Error("Policy generation is faulty")	
	}
}

func TestValidatePassword(t * testing.T) {
	validPasword := "1-3 a: abcde"
	invalidPassword := "1-3 b: cdefg"

	isValid := ValidatePassword(validPasword)

	if (!isValid) {
		t.Error("The valid password is not valid")
	}

	isValid = ValidatePassword(invalidPassword)

	if (isValid) {
		t.Error("The invalid password is valid")
	}
}

func TestValidatePasswordStep(t *testing.T) {
	valid := "1-3 a: abcde"
	invalid1 := "1-3 b: cdefg"
	invalid2 := "2-9 c: ccccccccc"

	if !ValidatePasswordStep(valid) {
		t.Error("should be valid")
	}

	if ValidatePasswordStep(invalid1) {
		t.Error("should be invalid")
	}

	if ValidatePasswordStep(invalid2) {
		t.Error("should be invalid")	
	}
}

func Test1(t *testing.T) {
	actual := 628
	records := utils.ReadDat("./input.dat")

	validPwdCount := 0

	for _, record := range(records) {
		if (ValidatePassword(record)) {
			validPwdCount++
		}
	}

	if (validPwdCount != actual) {
		t.Errorf("Valid passwords should be %d", actual)
	}
}

func Test2(t *testing.T) {
	actual := 705
	records := utils.ReadDat("./input.dat")

	validPwdCount := 0

	for _, record := range(records) {
		if (ValidatePasswordStep(record)) {
			validPwdCount++
		}
	}

	if (validPwdCount != actual) {
		t.Errorf("Valid passwords should be %d", actual)
	}
}