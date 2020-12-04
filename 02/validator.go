package a02

import (
	"strconv"
	"strings"
)

// ValidatePassword check that a password is following policies
func ValidatePassword(policyAndPassword string) bool {
	count := 0
	polPwdSlice := strings.Split(policyAndPassword, ":")
	atMin, atMax, char := readPolicy(&polPwdSlice[0])
	password := readPassword(&polPwdSlice[1])

	for i := 0; i < len(password); i++ {
		if (char == string(password[i])) {
			count++
		}
	}

	return atMin <= count && count <= atMax
}

// ValidatePasswordStep ...
func ValidatePasswordStep(policyAndPassword string) bool {
	polPwdSlice := strings.Split(policyAndPassword, ":")
	firstPos, secondPos, char := readPolicy(&polPwdSlice[0])
	password := readPassword(&polPwdSlice[1])
	first := string(password[firstPos - 1]) == char
	second := string(password[secondPos - 1]) == char
	isValid := (first || second) && (first != second)

	return isValid  
}

func readPassword(rawPassword *string) string {
	pwd := strings.TrimSpace(*rawPassword)

	return pwd
}

func readPolicy(policy *string) (n1 int, n2 int, char string) {
	policyRes := strings.Split(*policy, " ")
	minMaxRes := strings.Split(policyRes[0], "-")
	char = policyRes[1]
	n1, err := strconv.Atoi(minMaxRes[0])

	if (err != nil) { panic(err) }

	n2, err = strconv.Atoi(minMaxRes[1])

	if (err != nil) { panic(err) }

	return
}