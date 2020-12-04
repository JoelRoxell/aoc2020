package a04

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/joelroxell/aoc2020/utils"
)

type Object = map[string]interface{}

type Rule struct {
	regex string
	count int
	min int
	max int
	validator func(string) bool
}

func ParsePassports(file string) []Object {
	records := utils.ReadDat(file)
	passports := []Object{}
	currentPassport := Object{}

	for i := 0; i < len(records); i++ {
		record := records[i]

		if (len(record) == 0) {
			passports = append(passports, currentPassport)
			currentPassport = Object{}

			continue	
		}

		properties := strings.Split(record, " ")

		for _, prop := range properties {
			keyValuePair := strings.Split(prop, ":")

			currentPassport[keyValuePair[0]] = keyValuePair[1]
		}
	}

	if len(records[len(records) - 1]) != 0 {
		passports = append(passports, currentPassport)
	}

	return passports
}


var requiredFields = [...]string {"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func IsValidPassport(passport Object) bool {
	for _, key := range requiredFields {
		if passport[key] == nil {
			return false
		}
	}

	return true
}

var rules = map[string]Rule {
	"byr":  {
		count: 4,
		min: 1920,
		max: 2002,
	},
	"iyr": {
		count: 4,
		min: 2010,
		max: 2020,
	},
	"eyr": {
		count: 4,
		min: 2020,
		max: 2030,
	},
	"hgt": {
		validator: func(s string) bool {
			if (strings.Contains(s, "cm")) {
				value, _:= strconv.Atoi(strings.Split(s, "cm")[0])

				return value >= 150 && value <= 193
			} else if strings.Contains(s, "in") {
				value, _:= strconv.Atoi(strings.Split(s, "in")[0])

				return value >= 59 && value <= 76
			}

			return false
		},
	},
	"hcl": {
		regex: "^#[0-9a-f]{6}$",	
	},
	"ecl": {
		validator: func(s string) bool {
			guards := []string {"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

			return strings.Contains(strings.Join(guards, ","), s)
		},
	},
	"pid": {
		regex: "^[0-9]{0,9}$",
	},
}

func IsValidPassport2(passport Object) bool {
	isValid := true

	for _, key := range requiredFields {
		if passport[key] == nil {
			return false
		}

		value := passport[key].(string)

		if (rules[key].count > 0) {
			if (len(value) != rules[key].count) {
				return false
			}
		}
		
		if (rules[key].max > 0) {
			intValue, _ := strconv.Atoi(value) 

			if (intValue > rules[key].max) {
				return false
			}
		}
		
		if (rules[key].min > 0) {
			intValue, _ := strconv.Atoi(value) 

			if (intValue < rules[key].min) {
				return false 
			}
		}

		if (len(rules[key].regex) > 0) {
			reg := rules[key].regex

			matched, _ := regexp.MatchString(reg, value)

			if (!matched) {
				return false
			}
		}

		if (rules[key].validator != nil) {
			if (!rules[key].validator(value)) {
				return false
			}
		}
	}

	return isValid
}

