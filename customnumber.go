//Package customnumber provides the ability to create custom positional numeral
// systems in an efficient and performant way. You can create custom numbers based
// on custom numeral systems and use them at will.
//
//
// Each digit represented as a circular list that contains the all the possible numerals.
//
// Each number is represented as a doubly linked list of circular lists.
//
// Example
//
//  // create a slice of runes.
//  digitValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
//
//  number := customnumber.NewNumber(digitValues, "128z")
//
//  // will make the number 1290.
//  number.Increment()
//
//  // will make the number 128y.
//  number.Decrement()
//
//  //will give you the string representation of the number.
//  strnumber:=number.String()
package customnumber

import (
	"bytes"
	"container/list"
	"container/ring"
	"fmt"
	"math"
	"strings"
)

// Number represents a custom number that is consisted by its digits
// and digit values.
type Number struct {
	digits      *list.List
	digitValues []rune
}

// NewNumber initializes a CustomNumber by providing the initial number in strings
// along with the possible values that each digit can have.
func NewNumber(values []rune, initial string) (*Number, error) {
	// initialise a new number.
	number := Number{
		digits:      list.New(),
		digitValues: values,
	}
	// add digits to the number along with their state.
	for i := 0; i < len(initial); i++ {
		digit, err := newDigit(values, rune(initial[i]))
		if err != nil {
			return nil, err
		}
		number.digits.PushBack(digit)
	}
	return &number, nil
}

// newDigit creates and initializes a new digit (ring).
func newDigit(values []rune, state rune) (*ring.Ring, error) {
	// initialize a new empty ring
	r := ring.New(len(values))

	// fill the ring with values
	for _, e := range values {
		r.Value = e
		r = r.Next()
	}

	if indexOf(state, values) == -1 {
		return nil, fmt.Errorf("invalid digit. value: %v does not exist in possible values: %v", state, values)
	}

	// roll the ring in desired "state" position.
	for range values {
		if r.Value == state {
			break
		}
		r = r.Next()
	}
	return r, nil
}

// Increment performs a +1 to the Number.
func (n *Number) Increment() error {
	// take the first digit from the right and keep going if there are any arithmetic holdings.
	for e := n.digits.Back(); e != nil; e = e.Prev() {
		// get current ring.
		r := e.Value.(*ring.Ring)

		// rotate and update.
		r = r.Next()
		e.Value = r

		// if the digit is not being reset (no arithmetic holdings) then there is no need to
		// proceed in adding on the others.
		if r.Value != n.digitValues[0] {
			break
		}

		// If needed add an extra new digit on the left side.
		if e.Prev() == nil {
			d, _ := newDigit(n.digitValues, n.digitValues[0])
			n.digits.PushFront(d)
		}
	}
	return nil
}

// Decrement performs a -1 to the Number.
func (n *Number) Decrement() error {
	// take the first digit from the right and keep going if there are any arithmetic holdings or if the number is 0.
	for d := n.digits.Back(); d != nil; d = d.Prev() {
		// get current ring.
		r := d.Value.(*ring.Ring)
		// rotate and update
		r = r.Prev()
		d.Value = r

		// if the digit has not returned to it's last state then
		// there is no need to continue.
		if r.Value != n.digitValues[len(n.digitValues)-1] {
			break
		}

		if d.Prev() == nil {
			return fmt.Errorf("customnumber: can not Decrement")
		}
	}
	return nil
}

// Decimal converts a custom number to a decimal integer.
func (n *Number) Decimal() (int, error) {
	dec := 0
	di := 0
	for d := n.digits.Back(); d != nil; d = d.Prev() {
		// get current ring.
		r := d.Value.(*ring.Ring)
		// get the index of the ring.
		i := indexOf(r.Value.(rune), n.digitValues)

		// Add digit's decimal counterpart to the decimal sum.
		dec = dec + i*powInt(len(n.digitValues), di)
		di++
	}
	return dec, nil
}

// NewFromDecimal creates a custom number from a decimal integer.
func NewFromDecimal(values []rune, decimal int) (*Number, error) {
	dividend := decimal
	quotient := decimal
	divisor := len(values)
	sNumber := new(strings.Builder)
	for quotient > 0 {

		if dividend < divisor {
			break
		}

		quotient := dividend / divisor
		remainder := dividend % divisor

		//prepend character
		s := new(strings.Builder)
		s.WriteRune(values[remainder])
		s.WriteString(sNumber.String())
		sNumber = s
		//previous remainder will be the new dividend
		dividend = quotient
	}

	//prepend last remaining character
	s := new(strings.Builder)
	s.WriteRune(values[dividend%divisor])
	s.WriteString(sNumber.String())
	sNumber = s

	customNumber, err := NewNumber(values, sNumber.String())
	if err != nil {
		return nil, fmt.Errorf("creating converted number from decimal, err: %v", err)
	}
	return customNumber, nil
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func indexOf(element rune, data []rune) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

// String prints a string representation of Number.
func (n Number) String() string {
	// Loop over container list.
	var numberBytes bytes.Buffer
	for e := n.digits.Front(); e != nil; e = e.Next() {
		r := e.Value.(*ring.Ring)
		v := r.Value.(rune)
		numberBytes.WriteString(string(v))
	}
	return numberBytes.String()
}
