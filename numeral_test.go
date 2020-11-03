package numeral_test

import (
	"testing"

	"github.com/slysterous/numeral"
)

var testValues = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

func TestNumeralString(t *testing.T) {
	want := "150000"
	number, _ := numeral.NewNumeral(testValues, want)
	if got, want := number.String(), want; got != want {
		t.Errorf("String of numeral, want: %s got: %s", want, got)
	}
}

func TestIncrement(t *testing.T) {
	incrementTests := []struct {
		number string
		want   string
	}{
		{"150000", "150001"}, // Regular Increment
		{"zz", "100"},        // Increment that adds a digit
		{"0", "1"},           // Increment that adds a digit

	}
	for _, tt := range incrementTests {
		t.Run(tt.number, func(t *testing.T) {
			want := tt.want
			number, _ := numeral.NewNumeral(testValues, tt.number)
			number.Increment()
			if got, want := number.String(), want; got != want {
				t.Errorf("got: %s want: %s", got, want)
			}
		})
	}
}

func TestDecrement(t *testing.T) {
	decrementTests := []struct {
		number string
		want   string
	}{
		{"150001", "150000"}, // Regular decrement
		{"15000", "14zzz"},   // Decrement that borrows carry
		{"1000", "0zzz"},     // Decrement creates leading carry
	}
	for _, tt := range decrementTests {
		t.Run(tt.number, func(t *testing.T) {
			want := tt.want
			number, _ := numeral.NewNumeral(testValues, tt.number)
			err := number.Decrement()
			if err != nil {
				t.Errorf("expected nil,got err:%v", err)
			}
			if got, want := number.String(), want; got != want {
				t.Errorf("got: %s want: %s", got, want)
			}
		})
	}
}

func TestDecimal(t *testing.T) {
	decimalTests := []struct {
		number string
		want   int
	}{
		{"1", 1},
		{"100", 1296},
		{"abc", 13368},
	}
	for _, tt := range decimalTests {
		t.Run(tt.number, func(t *testing.T) {
			want := tt.want
			number, _ := numeral.NewNumeral(testValues, tt.number)
			dNumeral := number.Decimal()
			if dNumeral != want {
				t.Errorf("got: %d, want: %d", dNumeral, want)
			}
		})
	}
}

func TestFromDecimal(t *testing.T) {
	fromDecimalTests := []struct {
		number int
		want   string
	}{
		{1, "1"},
		{1296, "100"},
		{13368, "abc"},
	}

	for _, tt := range fromDecimalTests {
		t.Run(tt.want, func(t *testing.T) {
			want := tt.want
			dNumeral, err := numeral.NewFromDecimal(testValues, tt.number)
			if err != nil {
				t.Errorf("expected nil got err: %v", err)
			}
			if dNumeral.String() != want {
				t.Errorf("got: %s, want: %s", dNumeral.String(), want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	sumTests := []struct {
		number1 string
		number2 string
		want    string
	}{
		{"1", "1", "2"},
		{"a", "z", "19"},
		{"10", "10", "20"},
	}

	for _, tt := range sumTests {
		t.Run(tt.want, func(t *testing.T) {
			want := tt.want
			number1, err := numeral.NewNumeral(testValues, tt.number1)
			if err != nil {
				t.Errorf("expected nil got err: %v", err)
			}
			number2, err := numeral.NewNumeral(testValues, tt.number2)
			if err != nil {
				t.Errorf("expected nil got err: %v", err)
			}
			sum, _ := numeral.Sum(testValues, *number1, *number2)
			if sum.String() != want {
				t.Errorf("got: %s, want: %s", sum, want)
			}
		})
	}
}

func TestDiff(t *testing.T) {
	absDifferenceTests := []struct {
		number1 string
		number2 string
		want    string
	}{
		{"2", "9", "7"},
		{"10", "z", "1"},
		{"10", "1", "z"},
	}

	for _, tt := range absDifferenceTests {
		t.Run(tt.want, func(t *testing.T) {
			want := tt.want
			number1, err := numeral.NewNumeral(testValues, tt.number1)
			if err != nil {
				t.Errorf("expected nil got err: %v", err)
			}
			number2, err := numeral.NewNumeral(testValues, tt.number2)
			if err != nil {
				t.Errorf("expected nil got err: %v", err)
			}
			absDiff, _ := numeral.Diff(testValues, *number1, *number2)
			if absDiff.String() != want {
				t.Errorf("got: %s, want: %s", absDiff, want)
			}
		})
	}
}

func TestNumeralAdd(t *testing.T) {
	addTests := []struct {
		number1 string
		number2 string
		want    string
	}{
		{"1", "1", "2"},
		{"a", "z", "19"},
		{"10", "10", "20"},
	}

	for _, tt := range addTests {
		t.Run(tt.want, func(t *testing.T) {
			want := tt.want
			number1, err := numeral.NewNumeral(testValues, tt.number1)
			if err != nil {
				t.Errorf("expected nil got err: %v", err)
			}
			number2, err := numeral.NewNumeral(testValues, tt.number2)
			if err != nil {
				t.Errorf("expected nil got err: %v", err)
			}
			err = number1.Add(*number2)
			if err != nil {
				t.Errorf("expected nil got err: %v", err)
			}
			if number1.String() != want {
				t.Errorf("got: %s, want: %s", number1.String(), want)
			}
		})
	}
}

func TestNewNumeralWrongRunes(t *testing.T) {
	testValues := []rune{'3', '4', '5', '6', '7', '8'}
	_, err := numeral.NewNumeral(testValues, "0000")
	if err == nil {
		t.Errorf("expected error to be thrown on NewNumeral")
	}
}

func TestDecrementOnZeroThrowsErr(t *testing.T) {
	number, _ := numeral.NewNumeral(testValues, "0")
	err := number.Decrement()
	if err == nil {
		t.Errorf("expected error to be thrown on Decrement")
	}
	if number.String()!="0"{
		t.Errorf("expected: 0, got: %s ",number.String())
	}
}
