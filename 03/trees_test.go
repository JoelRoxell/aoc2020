package o3

import (
	"bufio"
	"os"
	"testing"
)


func readFieldFromDat(file string) []string {
    f, err := os.Open(file)

    if err != nil {
        panic(err)
    }

    s := bufio.NewScanner(f)
    s.Split(bufio.ScanLines)

    field := []string{}

    for s.Scan() {
        field = append(field, s.Text())
    }

    return field
}

func TestTreeTraverseCounterTestData(t *testing.T) {
    trees := TraverseTreesCounter(readFieldFromDat("./test-field.dat"), 3, 1)

    if (trees != 7) {
        t.Error("Didn't find the correct amount of trees")
    }
}

func TestTreeTraverseCounterInputData(t *testing.T) {
    trees := TraverseTreesCounter(readFieldFromDat("./input.dat"), 3, 1)

    if (trees != 169) {
        t.Error("Didn't find the correct amount of trees")
    }
}

func TestTreeTraverseCounter2InputData(t *testing.T) {
    field := readFieldFromDat("./input.dat")

    runs := [5][2]int {
        {1,1},
        {3,1},
        {5,1},
        {7,1},
        {1,2},
    }

    productOfTrees := 1

    for _, run := range (runs) {
       count := TraverseTreesCounter(field, run[0], run[1]) 

       productOfTrees = productOfTrees * count
    }

    if (productOfTrees != 7560370818) {
        t.Error("Didn't find the correct product of trees")
    }
}