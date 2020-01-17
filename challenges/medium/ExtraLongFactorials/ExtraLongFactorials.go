package ExtraLongFactorials

import "math/big"

func ExtraLongFactorials(n int64) *big.Int {
	var result = new(big.Int)
	return result.MulRange(1, n)
}
