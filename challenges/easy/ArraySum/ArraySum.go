package ArraySum

func ArraySum(arr []int) int {
	sum := 0

	for i := range arr {
		sum += arr[i]
	}

	return sum
}
