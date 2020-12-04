package a06

import (
	"bufio"
	"os"
	"strings"
)

func Unique(s string) string {
	u := make(map[string]bool)
	res := ""

	for _, v := range s {
		if u[string(v)] {
			continue
		} else {
			u[string(v)] = true
			res += string(v)
		}
	}

	return res
}

func ReadGroups(file string) [][]string {
	f, err := os.Open(file)

	if err != nil {
		panic(err)
	}

	defer f.Close()	

	r := bufio.NewReader(f)

	group := make([]string, 0)
	groups := make([][]string, 0)
	eof := false

	for {
		s, err := r.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				eof = true
			} else {
				panic(err)
			}
		}

		record := strings.TrimSpace(s)

		if len(record) == 0 {
			groups = append(groups, group)
			group	= make([]string, 0) 
		} else {
			group = append(group, record)
		}

		if eof {
			groups = append(groups, group)
			break
		}
	}

	return groups
}

func Includes(answers []string) int {
	uniqueAnswers := Unique(Join(answers))
	memberCount := len(answers)
	count := 0

	for _, c := range uniqueAnswers {
		charCount := 0
		
		for i := 0; i < len(answers); i++ {
			if strings.Contains(answers[i], string(c)) {
				charCount++
			}
		}

		if (charCount == memberCount) {
			count++
		}
	}

	return count
}

func Join(answers []string) string {
	res := ""

	for _, s := range answers {
		res += s
	}

	return res
}