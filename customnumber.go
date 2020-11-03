package customnumber

import (
	"bytes"
	"container/list"
	"container/ring"
	"fmt"
	"strings"
)

// => 0 1 2 3 4 5 6 7 8 9 a b c d e f g h i j k l m n o p q r s t u v w x y z  arithmetic system
// var value = "0a9esd"
// value.increment() => 0a8ese
//==================================
// var value = "0a9esz"
// value.increment() => 0a8et0
//=====================================
//start scrapping
// 000000
// .
// .
// 00000z
// 000010

// Number represents a custom number that is consisted by its digits
// and digit values.
type Number struct {
	Digits      *list.List
	DigitValues []rune
}

// NewNumber initializes a CustomNumber by providing the initial number in strings
// along with the possible values that each digit can have.
func NewNumber(values []rune, initial string) Number {
	// initialise a new number.
	number := Number{
		Digits:      list.New(),
		DigitValues: values,
	}
	// add digits to the number along with their state.
	for i := 0; i < len(initial); i++ {
		digit := newDigit(values, rune(initial[i]))
		number.Digits.PushBack(digit)
	}

	return number
}

// newDigit creates and initializes a new digit (ring).
func newDigit(values []rune, state rune) ring.Ring {
	// initialize a new empty ring
	r := ring.New(len(values))

	// fill the ring with values
	for _, e := range values {
		r.Value = e
		r = r.Next()
	}

	// roll the ring in desired "state" position.
	for range values {
		if r.Value == state {
			break
		}
		r = r.Next()
	}

	return *r
}

// Increment performs a +1 to the Number.
func (p *Number) Increment() error {
	// take the first digit from the right and keep going if there are any arithmetic holdings.
	for e := p.Digits.Back(); e != nil; e = e.Prev() {
		// get current ring.
		r, ok := e.Value.(ring.Ring)
		if !ok {
			return fmt.Errorf("customnumber: could not get ring value")
		}

		// rotate and update.
		r = *r.Next()
		e.Value = r

		// if the digit is not being reset (no arithmetic holdings) then there is no need to
		// proceed in adding on the others.
		if r.Value != p.DigitValues[0] {
			return nil
		}

		// If needed add an extra new digit on the left side.
		if e.Prev() == nil {
			d := newDigit(p.DigitValues, p.DigitValues[0])
			p.Digits.PushFront(d)
		}
	}
	return nil
}

// Decrement performs a -1 to the Number.
func (p *Number)Decrement() error{
	// take the first digit from the right and keep going if there are any arithmetic holdings or if the number is 0.
	for e := p.Digits.Back(); e != nil; e = e.Prev() {
		// get current ring.
		r, ok := e.Value.(ring.Ring)
		if !ok {
			return fmt.Errorf("customnumber: could not get ring value")
		}

		// rotate and update
		r = *r.Prev()
		e.Value = r

		// if the digit has not returned to it's last state then
		// there is no need to continue.
		if r.Value != p.DigitValues[len(p.DigitValues)-1] {
			return nil
		}

		if e.Prev() == nil {
			return fmt.Errorf("customnumber: can not Decrement")
		}
	}
	return nil
}

// String prints a string representation of Number.
func (p Number) String() string {
	// Loop over container list.
	var numberBytes bytes.Buffer
	for e := p.Digits.Front(); e != nil; e = e.Next() {
		r, ok := e.Value.(ring.Ring)
		if !ok {
			return ""
		}

		v, ok := r.Value.(rune)
		if !ok {
			return ""
		}

		numberBytes.WriteString(string(v))
	}
	return numberBytes.String()
}

// SmartString returns a string representation of a customnumber while removing leading 0s.
func (p Number) SmartString() string {
	str := p.String()
	return strings.TrimLeft(str, "0")
}
