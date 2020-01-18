package DiagonalDifference

import "testing"

func TestDiagonalDifference(t *testing.T) {
	arr := [][]int{{1,2,3},{4,5,6},{9,8,9}}
	diff := DiagonalDifference(arr)
	if diff != 2 {
		t.Errorf("DiagonalDifference([][]int{{1,2,3},{4,5,6},{9,8,9}}) expects 2; actual %v", diff)
	}

	arr = [][]int{{1,1,1},{2,2,2},{3,3,3}}
	diff = DiagonalDifference(arr)
	if diff != 0 {
		t.Errorf("DiagonalDifference([][]int{{1,1,1},{2,2,2},{3,3,3}}) expects 0; actual %v", diff)
	}

	arr = [][]int{{33,15,4},{13,21,2},{5,99,40}}
	diff = DiagonalDifference(arr)
	if diff != 64 {
		t.Errorf("DiagonalDifference([][]int{{33,15,4},{13,21,2},{5,99,40}}) expects 64; actual %v", diff)
	}
}
