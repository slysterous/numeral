package customnumber_test

import (
	"fmt"

	customnumber "github.com/slysterous/custom-number"
)

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
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	number, _ := customnumber.NewNumber(testValues, "100")
	dec := number.Decimal()

	fmt.Printf("%d", dec)
	// Output: 1296
}

func ExampleNumber_String() {
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	number, err := customnumber.NewFromDecimal(testValues, 2021)
	if err != nil {
		//handle the err
	}
	fmt.Printf("custom number as a string representation: %s", number.String())
}

func ExampleNumber_Add() {
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	number, err := customnumber.NewNumber(testValues, "123z")
	if err != nil {
		//handle the error
	}
	num2, err := customnumber.NewFromDecimal(testValues, 1)
	if err != nil {
		//handle error
	}
	err = number.Add(*num2)
	if err != nil {
		//handle error
	}
}

func ExampleAbsDifference() {
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	num, err := customnumber.NewFromDecimal(testValues, 22)
	if err != nil {
		//handle the error
	}
	num2, err := customnumber.NewFromDecimal(testValues, 12)
	if err != nil {
		//handle error
	}
	num3, err := customnumber.AbsDifference(testValues, *num2, *num)
	if err != nil {
		//handle error
	}
	fmt.Printf("decimal result should be 10, number: %d", num3.Decimal())
}

func ExampleNewFromDecimal() {
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	number, _ := customnumber.NewFromDecimal(testValues, 100)
	fmt.Printf("custom number: %v", number)
}

func ExampleNewNumber() {
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	number, _ := customnumber.NewNumber(testValues, "100")
	fmt.Printf("custom number: %v", number)
}

func ExampleSum() {
	testValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	number1, err := customnumber.NewNumber(testValues, "100")
	if err != nil {
		//handle the error
	}
	number2, err := customnumber.NewNumber(testValues, "100")
	if err != nil {
		//handle the error
	}
	sum, err := customnumber.Sum(testValues, *number1, *number2)
	if err != nil {
		//handle the error
	}

	fmt.Printf("sum is: %s", sum.String())
}
