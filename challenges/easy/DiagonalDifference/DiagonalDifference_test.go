package DiagonalDifference

import "testing"

func TestDiagonalDifference(t *testing.T) {
	arr := [][]int{{1,2,3},{4,5,6},{9,8,9}}

	diff := DiagonalDifference(arr)
	if diff != 2 {
		t.Errorf("DiagonalDifference([][]int{{1,2,3},{4,5,6},{9,8,9}}) expects 2; actual %v", diff)
	}
}
