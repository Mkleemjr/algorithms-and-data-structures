package ArraySum

import "testing"

func TestArraySum(t *testing.T) {
	arr := []int{1,2,3}
	sum := ArraySum(arr)
	if sum != 6 {
		t.Errorf("ArraySum([1,2,4]) expects 6; actual %v", sum)
	}
	
	arr = []int{10000,1000,100,10,1}
	sum = ArraySum(arr)
	if sum != 11111 {
		t.Errorf("ArraySum([10000,1000,100,10,1]) expects 11111; actual %v", sum)
	}
}
