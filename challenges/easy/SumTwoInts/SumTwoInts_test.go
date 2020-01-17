package SumTwoInts

import "testing"

func PrintError(t *testing.T, a, b, expect_x, actual_x uint32) {
  t.Errorf("SumTwoInts(%d, %d) returns %d; expects %d", a, b, actual_x, expect_x)
}

func TestSumTwoInts(t *testing.T) {
	var sum uint32

	sum = SumTwoInts(1, 1)
	if sum != 2 { PrintError(t, 1, 1, 2, sum) }
	
	sum = SumTwoInts(1, 2)
	if sum != 3 { PrintError(t, 1, 2, 3, sum) }
	
	sum = SumTwoInts(2, 2)
	if sum != 4 { PrintError(t, 2, 2, 4, sum) }
	
	sum = SumTwoInts(5, 7)
	if sum != 12 { PrintError(t, 5, 7, 12, sum) }
	
	sum = SumTwoInts(25, 100)
	if sum != 125 { PrintError(t, 25, 100, 125, sum) }
}
