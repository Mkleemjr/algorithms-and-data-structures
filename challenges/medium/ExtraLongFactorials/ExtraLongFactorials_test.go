package ExtraLongFactorials

import (
	"math/big"
	"fmt"
	"testing"
)

func BigIntToStr(bigInt *big.Int) string {
	return fmt.Sprintf("%v", bigInt)
}

func PrintError(t *testing.T, n int64, expect_y *big.Int, actual_y *big.Int) {
	t.Errorf("ExtraLongFactorials(%v) returns %v; expects %v", n, BigIntToStr(actual_y), BigIntToStr(expect_y))
}

func TestExtraLongFactorials1(t *testing.T) {
	var n int64 = 30;
	var expect, _ = new(big.Int).SetString("265252859812191058636308480000000", 10)
	var actual *big.Int

	actual = ExtraLongFactorials(n)

	if actual.Cmp(expect) != 0 {
		PrintError(t, n, expect, actual)
	}
}

func TestExtraLongFactorials2(t *testing.T) {
	var n int64 = 51;
	var expect, _ = new(big.Int).SetString("1551118753287382280224243016469303211063259720016986112000000000000", 10)
	var actual *big.Int

	actual = ExtraLongFactorials(n)

	if actual.Cmp(expect) != 0 {
		PrintError(t, n, expect, actual)
	}
}

func TestExtraLongFactorials3(t *testing.T) {
	var n int64 = 3;
	var expect, _ = new(big.Int).SetString("6", 10)
	var actual *big.Int

	actual = ExtraLongFactorials(n)

	if actual.Cmp(expect) != 0 {
		PrintError(t, n, expect, actual)
	}
}

func TestExtraLongFactorials4(t *testing.T) {
	var n int64 = 6;
	var expect, _ = new(big.Int).SetString("720", 10)
	var actual *big.Int

	actual = ExtraLongFactorials(n)

	if actual.Cmp(expect) != 0 {
		PrintError(t, n, expect, actual)
	}
}
