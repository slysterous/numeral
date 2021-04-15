package numeral_test

import (
	"testing"

	"github.com/slysterous/numeral"
)

func benchmarkNumeralIncrement(initialValue string , values []rune, b *testing.B){
	num,err:=numeral.NewNumeral(values, initialValue)
	if err != nil {
		b.Fatal(err)
	}
	for n := 0; n < b.N; n++ {
		_=num.Increment()
	}
}

func BenchmarkNumeralIncrementDecimalSmall(b *testing.B) {
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	initValue:="0"
	benchmarkNumeralIncrement(initValue,testValues,b)
}
func BenchmarkNumeralIncrementDecimalLarge(b *testing.B) {
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	initValue:="1111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"
	benchmarkNumeralIncrement(initValue,testValues,b)
}

func BenchmarkNumeralIncrementBinarySmall(b *testing.B) {
	testValues := []rune{'0', '1'}
	initValue:="0"
	benchmarkNumeralIncrement(initValue,testValues,b)
}
func BenchmarkNumeralIncrementBinaryLarge(b *testing.B) {
	testValues := []rune{'0', '1'}
	initValue:="1111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"
	benchmarkNumeralIncrement(initValue,testValues,b)
}

func BenchmarkNumeralIncrementHexSmall(b *testing.B) {
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9','a', 'b', 'c', 'd', 'e', 'f'}
	initValue:="0"
	benchmarkNumeralIncrement(initValue,testValues,b)
}
func BenchmarkNumeralIncrementHexLarge(b *testing.B) {
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9','a', 'b', 'c', 'd', 'e', 'f'}
	initValue:="1111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"
	benchmarkNumeralIncrement(initValue,testValues,b)
}

func BenchmarkNumeral_Decrement(b *testing.B) {
	num,err:=numeral.NewNumeral(testValues,"zzzzzzzzzzzzzzzzzzzzz")
	if err != nil {
		//whatever
	}
	for n := 0; n < b.N; n++ {
		_=num.Decrement()
		//if err !=nil{
		//	b.Errorf("expected nil, got err:%v",err)
		//}
	}
}


