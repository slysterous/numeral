package customnumber_test

import (
	"fmt"
	"github.com/slysterous/custom-number"
	"testing"
)

var testValues = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

func TestCustomNumberString(t *testing.T) {
	want := "150000"
	number, _ := customnumber.NewNumber(testValues, want)
	if got, want := number.String(), want; got != want {
		t.Errorf("String of custom number, want: %s got: %s", want, got)
	}
}

var incrementTests = []struct {
	number string
	want   string
}{
	{"150000", "150001"}, // Regular Increment
	{"zz", "100"},        // Increment that adds a digit
	{"0", "1"},           // Increment that adds a digit

}

func TestIncrement(t *testing.T) {
	for _, tt := range incrementTests {
		t.Run(tt.number, func(t *testing.T) {
			want := tt.want
			number, _ := customnumber.NewNumber(testValues, tt.number)
			number.Increment()
			if got, want := number.String(), want; got != want {
				t.Errorf("String of custom number, want: %s got: %s", want, got)
			}
		})
	}
}

var decrementTests = []struct {
	number string
	want   string
}{
	{"150001", "150000"}, // Regular decrement
	{"15000", "14zzz"},   // Decrement that borrows carry
	{"1000", "0zzz"},     // Decrement creates leading carry
}

func TestDecrement(t *testing.T) {
	for _, tt := range decrementTests {
		t.Run(tt.number, func(t *testing.T) {
			want := tt.want
			number, _ := customnumber.NewNumber(testValues, tt.number)
			err := number.Decrement()
			if err != nil {
				t.Errorf("Unexpected error on Decrement, err:%v", err)
			}
			if got, want := number.String(), want; got != want {
				t.Errorf("String of custom number, want: %s got: %s", want, got)
			}
		})
	}
}

var decimalTests = []struct {
	number string
	want   int
}{
	{"1", 1},       // Regular decrement
	{"100", 1296},  // Decrement that borrows carry
	{"abc", 13368}, // Decrement creates leading carry
}

func TestDecimal(t *testing.T) {
	for _, tt := range decimalTests {
		t.Run(tt.number, func(t *testing.T) {
			want := tt.want
			number, _ := customnumber.NewNumber(testValues, tt.number)
			dNumber, err := number.Decimal()
			if err != nil {
				t.Errorf("unexpected error on Decimal, err: %v", err)
			}
			if dNumber != want {
				t.Errorf("expected: %d, got: %d", want, dNumber)
			}
		})
	}
}

var fromDecimalTests = []struct {
	number int
	want   string
}{
	{1, "1"},       // Regular decrement
	{1296, "100"},  // Decrement that borrows carry
	{13368, "abc"}, // Decrement creates leading carry
}

func TestFromDecimal(t *testing.T) {
	for _, tt := range fromDecimalTests {
		t.Run(tt.want, func(t *testing.T) {
			want := tt.want
			dNumber, err := customnumber.NewFromDecimal(testValues, tt.number)
			if err != nil {
				t.Errorf("unexpected error on Decimal, err: %v", err)
			}
			if dNumber.String() != want {
				t.Errorf("expected: %s, got: %s", want, dNumber.String())
			}
		})
	}
}

func TestNewNumberWrongRunes(t *testing.T) {
	testValues := []rune{'3', '4', '5', '6', '7', '8'}
	_, err := customnumber.NewNumber(testValues, "0000")
	if err == nil {
		t.Errorf("expected error to be thrown on NewNumber")
	}
}

func TestDecrementOnZeroThrowsErr(t *testing.T) {
	number, _ := customnumber.NewNumber(testValues, "0000")
	err := number.Decrement()
	if err == nil {
		t.Errorf("expected error to be thrown on Decrement")
	}
}

func ExampleNumber_Increment() {
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	number, _ := customnumber.NewNumber(testValues, "123z")
	number.Increment()
	fmt.Printf(number.String())
	// Output: 1240
}

func ExampleNumber_Decrement() {
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	number, _ := customnumber.NewNumber(testValues, "1230")
	err := number.Decrement()
	if err != nil {
		// do whatever you need with the error
	}
	fmt.Printf(number.String())
	// Output: 122z
}

func ExampleNumber_Decimal() {
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	number, _ := customnumber.NewNumber(testValues, "100")
	dec, err := number.Decimal()

	if err != nil {
		// do whatever you need with the error
	}

	fmt.Printf("%d", dec)
	// Output: 1296
}

func ExampleNewFromDecimal() {
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	number, _ := customnumber.NewFromDecimal(testValues, 100)
	fmt.Printf("custom number: %v", number)
}

func ExampleNewNumber() {
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	number, _ := customnumber.NewNumber(testValues, "100")
	fmt.Printf("custom number: %v", number)
}
