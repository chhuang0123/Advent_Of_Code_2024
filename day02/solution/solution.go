package solution

const MaxDiff = 3

func IsAllIncreasing(arr []int) bool {
	maxIndex := len(arr) - 1
	for i := range arr {
		if i == maxIndex {
			break
		}

		if (arr[i+1] <= arr[i]) || (arr[i+1]-arr[i]) > MaxDiff {
			return false
		}
	}

	return true
}

func IsAllDecreasing(arr []int) bool {
	maxIndex := len(arr) - 1
	for i := range arr {
		if i == maxIndex {
			break
		}

		if (arr[i+1] >= arr[i]) || ((arr[i] - arr[i+1]) > MaxDiff) {
			return false
		}
	}

	return true
}

func IsAllIncreasingOrDecreasing(arr []int) bool {
	return IsAllIncreasing(arr) || IsAllDecreasing(arr)
}

func makeNewRecord(arr []int, removeIndex int) []int {
	left := arr[:removeIndex]
	right := arr[min(removeIndex+1, len(arr)):]
	return append(left, right...)
}

func IsAllIncreasingOrDecreasingWithErrorTolerance(arr []int) bool {
	for i := range arr {
		var temp = make([]int, len(arr))
		copy(temp, arr)
		newRecord := makeNewRecord(temp, i)

		if IsAllIncreasingOrDecreasing(newRecord) {
			return true
		}
	}

	return false
}
