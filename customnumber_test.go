package customnumber_test

import (
	customnumber "github.com/slysterous/custom-number"
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

func TestIncrement(t *testing.T) {
	want := "150001"
	values := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	number := customnumber.NewNumber(values, "150000")
	number.Increment()
	if got, want := number.String(), want; got != want {
		t.Errorf("String of custom number, want: %s got: %s", want, got)
	}

}

func TestIncrementThatAddsNewDigit(t *testing.T) {
	want := "100"
	values := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	number := customnumber.NewNumber(values, "zz")
	number.Increment()
	if got, want := number.String(), want; got != want {
		t.Errorf("String of custom number, want: %s got: %s", want, got)
	}
}
