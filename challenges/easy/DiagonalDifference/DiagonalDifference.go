package DiagonalDifference

import "math"

func DiagonalDifference(arr [][]int) int {
	firstSum := 0
	secondSum := 0

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			if i == j {
				firstSum += arr[i][j]
			}
			if len(arr[0]) - 1 - i == j {
				secondSum += arr[i][j]
			}
		}
	}

	return int(math.Abs(float64(firstSum) - float64(secondSum)))
}
