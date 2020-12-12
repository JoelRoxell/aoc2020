package a08

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/joelroxell/aoc2020/utils"
)


type Command struct {
	id string
	action string
	param int
}

func Boot(sequence []Command) int {
	acc := 0
	duplicateGuard := make(map[string]bool)
	var previousInstruction Command
		
	for i := 0; i < len(sequence); i++ {
		next := sequence[i]

		if duplicateGuard[next.id] {
			fmt.Printf("cmd %s with id %s is causing a loop using param %d", previousInstruction.action, previousInstruction.id, previousInstruction.param)
			break		
		}

		duplicateGuard[next.id] = true

		switch next.action {
		case "acc":
			acc += next.param	
		case "jmp":
			i += next.param - 1
		}	

		previousInstruction = next
	}

	return acc
}

var idCounter int = 0

func ParseStringToCmd(s string) Command {
	idCounter++
	parts := strings.Split(s, " ")
	action := parts[0]
	param := parts[1]

	positiveReg := regexp.MustCompile(`\+`)
	numberReg := regexp.MustCompile(`\d+`)
	isPositive :=  positiveReg.MatchString(param)
	stringNumber := numberReg.FindString(param)
	number, err := strconv.Atoi(stringNumber)

	if err != nil {
		panic(err)
	}

	if (!isPositive) {
		number = -number
	}

 return Command{id: strconv.Itoa(idCounter), action: action, param: number }	
}

func ReadInstructions(file string) []Command {
	records := utils.ReadDat(file)
	cmds := make([]Command, 0)

	for _, s := range records {
		cmds = append(cmds, ParseStringToCmd(s))
	}

	return cmds
}