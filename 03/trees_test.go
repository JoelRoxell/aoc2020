package o3

import (
	"testing"

	"github.com/joelroxell/aoc2020/utils"
)


func TestTreeTraverseCounterTestData(t *testing.T) {
    trees := TraverseTreesCounter(utils.ReadDat("./test-field.dat"), 3, 1)

    if (trees != 7) {
        t.Error("Didn't find the correct amount of trees")
    }
}

func TestTreeTraverseCounterInputData(t *testing.T) {
    trees := TraverseTreesCounter(utils.ReadDat("./input.dat"), 3, 1)

    if (trees != 169) {
        t.Error("Didn't find the correct amount of trees")
    }
}

func TestTreeTraverseCounter2InputData(t *testing.T) {
    field := utils.ReadDat("./input.dat")

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