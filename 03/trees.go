package o3

// TraverseTreesCounter ...
func TraverseTreesCounter(inputs []string, xStep int, yStep int) int {
	xPos := 0
	treeCount := 0
	mapWidth := len(inputs[0])
	
	for yPos := 0; yPos < len(inputs); yPos = yPos + yStep {
		character := inputs[yPos][xPos % mapWidth]

		if (string(character) == "#") {
			treeCount++
		}

		xPos += xStep
	}

	return treeCount
}
