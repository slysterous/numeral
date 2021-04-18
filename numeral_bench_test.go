package numeral_test

import (
	"testing"

	"github.com/slysterous/numeral"
)

func benchmarkNumeralIncrement(initialValue string, values []rune, b *testing.B) {
	num, err := numeral.NewNumeral(values, initialValue)
	if err != nil {
		b.Fatal(err)
	}
	for n := 0; n < b.N; n++ {
		_ = num.Increment()
	}
}

func benchmarkNumeralDecimal(num numeral.Numeral, b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = num.Decimal()
	}
}

func benchmarkNumeralAdd(num numeral.Numeral, num2 numeral.Numeral, b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = num.Add(num2)
	}
}

func benchmarkNumeralSum(num numeral.Numeral, num2 numeral.Numeral, b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = numeral.Sum(testValues, num, num2)
	}
}

func BenchmarkNumeralIncrementDecimalSmall(b *testing.B) {
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	initValue := "0"
	benchmarkNumeralIncrement(initValue, testValues, b)
}

func BenchmarkNumeralIncrementDecimalLarge(b *testing.B) {
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	initValue := "1111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"
	benchmarkNumeralIncrement(initValue, testValues, b)
}

func BenchmarkNumeralIncrementBinarySmall(b *testing.B) {
	testValues := []rune{'0', '1'}
	initValue := "0"
	benchmarkNumeralIncrement(initValue, testValues, b)
}

func BenchmarkNumeralIncrementBinaryLarge(b *testing.B) {
	testValues := []rune{'0', '1'}
	initValue := "1111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"
	benchmarkNumeralIncrement(initValue, testValues, b)
}

func BenchmarkNumeralIncrementHexSmall(b *testing.B) {
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	initValue := "0"
	benchmarkNumeralIncrement(initValue, testValues, b)
}

func BenchmarkNumeralIncrementHexLarge(b *testing.B) {
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	initValue := "1111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"
	benchmarkNumeralIncrement(initValue, testValues, b)
}

func BenchmarkNumeralDecimalFromHex0(b *testing.B) {
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	initValue := "0"
	num, _ := numeral.NewNumeral(testValues, initValue)
	benchmarkNumeralDecimal(*num, b)
}

func BenchmarkNumeralDecimalFromHexfffff(b *testing.B) {
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	initValue := "ffffffffff"
	num, _ := numeral.NewNumeral(testValues, initValue)
	benchmarkNumeralDecimal(*num, b)
}

func BenchmarkNumeralDecimalFromHexffffffffff(b *testing.B) {
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	initValue := "ffffffffff"
	num, _ := numeral.NewNumeral(testValues, initValue)
	benchmarkNumeralDecimal(*num, b)
}

func BenchmarkNumeralDecimalFromHexffffffffffffffffffff(b *testing.B) {
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	initValue := "ffffffffffffffffffff"
	num, _ := numeral.NewNumeral(testValues, initValue)
	benchmarkNumeralDecimal(*num, b)
}

func BenchmarkNumeralDecimalFromBinary0(b *testing.B) {
	testValues := []rune{'0', '1'}
	initValue := "0"
	num, _ := numeral.NewNumeral(testValues, initValue)
	benchmarkNumeralDecimal(*num, b)
}

func BenchmarkNumeralDecimalFromBinary10000(b *testing.B) {
	testValues := []rune{'0', '1'}
	initValue := "10000"
	num, _ := numeral.NewNumeral(testValues, initValue)
	benchmarkNumeralDecimal(*num, b)
}

func BenchmarkNumeralDecimalFromBinary1000000000(b *testing.B) {
	testValues := []rune{'0', '1'}
	initValue := "1000000000"
	num, _ := numeral.NewNumeral(testValues, initValue)
	benchmarkNumeralDecimal(*num, b)
}

func BenchmarkNumeralDecimalFromBinary1000000000000000000(b *testing.B) {
	testValues := []rune{'0', '1'}
	initValue := "1000000000000000000"
	num, _ := numeral.NewNumeral(testValues, initValue)
	benchmarkNumeralDecimal(*num, b)
}

func BenchmarkNumeralAddBinaryOnHex(b *testing.B) {
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	testValues2 := []rune{'0', '1'}
	initValue := "999"
	num, _ := numeral.NewNumeral(testValues, initValue)
	num2, _ := numeral.NewFromDecimal(testValues2, 999000)
	benchmarkNumeralAdd(*num, *num2, b)
}

func BenchmarkNumeralAddHexOnHex(b *testing.B) {
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	testValues2 := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	initValue := "999"
	num, _ := numeral.NewNumeral(testValues, initValue)
	num2, _ := numeral.NewFromDecimal(testValues2, 999000)
	benchmarkNumeralAdd(*num, *num2, b)
}

func BenchmarkNumeralAddDecOnDec(b *testing.B) {
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	testValues2 := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	initValue := "999"
	num, _ := numeral.NewNumeral(testValues, initValue)
	num2, _ := numeral.NewFromDecimal(testValues2, 999000)
	benchmarkNumeralAdd(*num, *num2, b)
}

func BenchmarkNumeralAddSumBinaryOnHex(b *testing.B) {
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	testValues2 := []rune{'0', '1'}
	initValue := "999"
	num, _ := numeral.NewNumeral(testValues, initValue)
	num2, _ := numeral.NewFromDecimal(testValues2, 999000)
	benchmarkNumeralSum(*num, *num2, b)
}

func BenchmarkNumeralAddSumHexOnHex(b *testing.B) {
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	testValues2 := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	initValue := "999"
	num, _ := numeral.NewNumeral(testValues, initValue)
	num2, _ := numeral.NewFromDecimal(testValues2, 999000)
	benchmarkNumeralSum(*num, *num2, b)
}

func BenchmarkNumeralAddSumDecOnDec(b *testing.B) {
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	testValues2 := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	initValue := "999"
	num, _ := numeral.NewNumeral(testValues, initValue)
	num2, _ := numeral.NewFromDecimal(testValues2, 999000)
	benchmarkNumeralSum(*num, *num2, b)
}
