package solution

func getChar(arr [][]string, rowIndex int, columnIndex int, rowStep int, columnStep int) string {
	rowCount := len(arr)
	columnCount := len(arr[0])

	finalRowIndex := rowIndex + rowStep
	finalColumnIndex := columnIndex + columnStep
	if finalRowIndex < 0 || finalRowIndex >= rowCount || finalColumnIndex < 0 || finalColumnIndex >= columnCount {
		return ""
	}

	return arr[finalRowIndex][finalColumnIndex]
}

func CheckVertical(arr [][]string, i int, j int) (result int) {
	if getChar(arr, i, j, 1, 0) == "M" &&
		getChar(arr, i, j, 2, 0) == "A" &&
		getChar(arr, i, j, 3, 0) == "S" {
		result++
	}

	if getChar(arr, i, j, -1, 0) == "M" &&
		getChar(arr, i, j, -2, 0) == "A" &&
		getChar(arr, i, j, -3, 0) == "S" {
		result++
	}

	return result
}

func CheckHorizontal(arr [][]string, i int, j int) (result int) {
	if getChar(arr, i, j, 0, 1) == "M" &&
		getChar(arr, i, j, 0, 2) == "A" &&
		getChar(arr, i, j, 0, 3) == "S" {
		result++
	}

	if getChar(arr, i, j, 0, -1) == "M" &&
		getChar(arr, i, j, 0, -2) == "A" &&
		getChar(arr, i, j, 0, -3) == "S" {
		result++
	}

	return result
}

func CheckDiagonal(arr [][]string, i int, j int) (result int) {
	if getChar(arr, i, j, 1, 1) == "M" &&
		getChar(arr, i, j, 2, 2) == "A" &&
		getChar(arr, i, j, 3, 3) == "S" {
		result++
	}

	if getChar(arr, i, j, -1, -1) == "M" &&
		getChar(arr, i, j, -2, -2) == "A" &&
		getChar(arr, i, j, -3, -3) == "S" {
		result++
	}

	if getChar(arr, i, j, 1, -1) == "M" &&
		getChar(arr, i, j, 2, -2) == "A" &&
		getChar(arr, i, j, 3, -3) == "S" {
		result++
	}

	if getChar(arr, i, j, -1, 1) == "M" &&
		getChar(arr, i, j, -2, 2) == "A" &&
		getChar(arr, i, j, -3, 3) == "S" {
		result++
	}

	return result
}

func CheckAllXmas(arr [][]string) (result int) {
	rowCount := len(arr)
	columnCount := len(arr[0])
	for i := 0; i < rowCount; i++ {
		for j := 0; j < columnCount; j++ {
			if arr[i][j] == "X" {
				result += CheckVertical(arr, i, j) + CheckHorizontal(arr, i, j) + CheckDiagonal(arr, i, j)
			}
		}
	}

	return result
}

func CheckAllCrossMas(arr [][]string) (result int) {
	rowCount := len(arr)
	columnCount := len(arr[0])
	for i := 0; i < rowCount; i++ {
		for j := 0; j < columnCount; j++ {
			if arr[i][j] == "A" {
				result += CheckCrossMas(arr, i, j)
			}
		}
	}

	return result
}

func CheckCrossMas(arr [][]string, i int, j int) (result int) {
	condition1 := (getChar(arr, i, j, 1, 1) == "M" && getChar(arr, i, j, -1, -1) == "S") ||
		(getChar(arr, i, j, 1, 1) == "S" && getChar(arr, i, j, -1, -1) == "M")

	condition2 := (getChar(arr, i, j, 1, -1) == "M" && getChar(arr, i, j, -1, 1) == "S") ||
		(getChar(arr, i, j, 1, -1) == "S" && getChar(arr, i, j, -1, 1) == "M")

	if condition1 && condition2 {
		result++
	}

	return result
}
