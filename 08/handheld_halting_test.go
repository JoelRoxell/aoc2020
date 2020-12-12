package a08

import (
	"strings"
	"testing"
)

func TestDemo(t *testing.T) {
	instructions := ReadInstructions("./demo.dat")
	count := Boot(instructions)

	if count != 5 {
		t.Errorf("Boot sequence result is invalid %d", count)
	}
}

func TestParseCommand(t *testing.T) {
	cmdText := "acc -1"

	cmd := ParseStringToCmd(cmdText)

	if strings.Compare(cmd.action, "acc") != 0 {
		t.Error("failed to parse command action")
	}

	if strings.Compare(cmd.id, "1") != 0 {
		t.Error("failed to parse comand action")
	}

	if cmd.param != -1 {
		t.Error("failed to parse comand param")
	}
}

func Test1(t *testing.T) {
	instructions := ReadInstructions("./input.dat")
	count := Boot(instructions)

	if count != 1801 {
		t.Errorf("Boot sequence result is invalid %d", count)
	}
}

func Test2(t *testing.T) {
	instructions := ReadInstructions("./input2.dat")
	count := Boot(instructions)

	if count != 2060 {
		t.Errorf("Boot sequence result is invalid %d", count)
	}
}