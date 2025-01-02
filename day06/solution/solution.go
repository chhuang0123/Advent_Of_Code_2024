package solution

func getStartPosition(guardMap [][]string) (string, int, int) {
	rowCount := len(guardMap)
	columnCount := len(guardMap[0])

	for i := 0; i < rowCount; i++ {
		for j := 0; j < columnCount; j++ {
			if guardMap[i][j] == "^" || guardMap[i][j] == "v" || guardMap[i][j] == "<" || guardMap[i][j] == ">" {
				return guardMap[i][j], i, j
			}

		}
	}

	return "", -1, -1
}

func getChar(arr [][]string, rowIndex int, columnIndex int, rowStep int, columnStep int) (string, int, int) {
	rowCount := len(arr)
	columnCount := len(arr[0])

	finalRowIndex := rowIndex + rowStep
	finalColumnIndex := columnIndex + columnStep
	if finalRowIndex < 0 || finalRowIndex >= rowCount || finalColumnIndex < 0 || finalColumnIndex >= columnCount {
		return "", -1, -1
	}

	return arr[finalRowIndex][finalColumnIndex], finalRowIndex, finalColumnIndex
}

func changeDirection(direction string) string {
	switch direction {
	case "^":
		return ">"
	case "v":
		return "<"
	case "<":
		return "^"
	case ">":
		return "v"
	}
	return ""
}

func checkNextStepMap(guardMap [][]string) ([][]string, bool) {
	direction, rowIndex, columnIndex := getStartPosition(guardMap)
	if direction == "" {
		return nil, false
	}

	rowStep, columnStep := 0, 0
	switch direction {
	case "^":
		rowStep, columnStep = -1, 0
	case "v":
		rowStep, columnStep = 1, 0
	case "<":
		rowStep, columnStep = 0, -1
	case ">":
		rowStep, columnStep = 0, 1
	}

	nextChar, nextRow, NextColumn := getChar(guardMap, rowIndex, columnIndex, rowStep, columnStep)
	switch nextChar {
	case ".", "X":
		guardMap[nextRow][NextColumn] = guardMap[rowIndex][columnIndex]
		guardMap[rowIndex][columnIndex] = "X"
	case "#":
		guardMap[rowIndex][columnIndex] = changeDirection(direction)
	case "":
		guardMap[rowIndex][columnIndex] = "X"
		return guardMap, true
	}

	return guardMap, false
}

func GetFinalMap(guardMap [][]string) [][]string {
	doneCondition := false
	for {
		guardMap, doneCondition = checkNextStepMap(guardMap)
		if doneCondition {
			return guardMap
		}
	}
}
