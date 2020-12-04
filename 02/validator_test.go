package a02

import (
	"bufio"
	"os"
	"testing"
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

func TestMain(t *testing.T) {
	actual := 705
	f, err := os.Open("./input-2.dat")

	if (err != nil) { panic(err) }

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	validPwdCount := 0

	for s.Scan() {
		if (ValidatePasswordStep(s.Text())) {
			validPwdCount++
		}
	}

	if (validPwdCount != actual) {
		t.Errorf("Valid passwords should be %d", actual)
	}
}