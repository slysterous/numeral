package customnumber_test

import (
	"fmt"
	"github.com/slysterous/custom-number"
	"testing"
)

func TestCustomNumberString(t *testing.T) {
	want := "150000"
	values := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	number := customnumber.NewNumber(values, want)

	if got, want := number.String(), want; got != want {
		t.Errorf("String of custom number, want: %s got: %s", want, got)
	}
}

func TestCustomNumberSmartString(t *testing.T) {
	initialValue := "000010"
	want := "10"
	values := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	number := customnumber.NewNumber(values, initialValue)

	if got, want := number.SmartString(), want; got != want {
		t.Errorf("Smart string of custom number, want: %s got: %s", want, got)
	}
}

var incrementTests = []struct {
	number string
	want   string
}{
	{"150000", "150001"}, // Regular Increment
	{"zz", "100"},        // Increment that adds a digit
}

func TestIncrement(t *testing.T) {
	for _, tt := range incrementTests {
		t.Run(tt.number, func(t *testing.T) {
			want := tt.want
			values := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
			number := customnumber.NewNumber(values, tt.number)
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
			values := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
			number := customnumber.NewNumber(values, tt.number)
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

func TestDecrementOnZeroThrowsErr(t *testing.T) {
	values := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	number := customnumber.NewNumber(values, "0000")
	err := number.Decrement()
	if err == nil {
		t.Errorf("Expected error to be thrown on Decrement")
	}
}

func ExampleNumber_Increment() {
	values := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	number := customnumber.NewNumber(values, "123z")
	err := number.Decrement()
	if err == nil {
		// do whatever you need with the error
	}
	fmt.Printf(number.String())
	// Output: 1240
}

func ExampleNumber_Decrement() {
	values := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	number := customnumber.NewNumber(values, "1230")
	err := number.Decrement()
	if err == nil {
		// do whatever you need with the error
	}
	fmt.Printf(number.String())
	// Output: 122z
}
