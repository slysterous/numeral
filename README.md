# custom-number [![PkgGoDev](https://pkg.go.dev/badge/github.com/slysterous/custom-number)](https://pkg.go.dev/github.com/slysterous/custom-number)
[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://lbesson.mit-license.org/)
[![CircleCI](https://circleci.com/gh/slysterous/custom-number.svg?style=shield)](https://circleci.com/gh/slysterous/custom-numbers)
[![codecov](https://codecov.io/gh/slysterous/custom-number/branch/main/graph/badge.svg?token=057BbZbRE4)](https://codecov.io/gh/slysterous/custom-number)
[![Go Report Card](https://goreportcard.com/badge/github.com/slysterous/custom-number)](https://goreportcard.com/report/github.com/slysterous/custom-number)


Custom numbers based on custom positional (numeral) systems.
## Table of Contents ##

## About ##

### Why
There are times, where we need to iterate over an X amount of possible combinations of values over a specific number or a specific string etc.
<strong> (iterators, rainbow tables, numeral system converters) </strong>
 
For example, if you have an identifier that can contain values that are either 0-9 and a-z and A-Z and you would want to know each possible 
combination.

Your base (or radix) is the amount of possible values a digit can have. In this example:
```.env
0-9 : 10

a-z : 26

A-Z : 26
```
That means that the base of the numeral system is 62. The possible values of a digit are:
```.env
0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ
```
Now, if you wanted to find all the possible combinations of numbers, for 4 digits, for the above example you would want to iterate from:
```.env
0000 to ZZZZ
```
These are <strong>916.132.832</strong> different possible combinations.

### How
A <strong>numeral</strong> is a symbol or group of symbols that <strong>represents a number</strong>. Numerals are not the
same as numbers just as words are not the same with the things they refer to. The symbols <strong>"11", "1011" and "B"</strong>
are different numerals, all representing the same number.

A <strong>numeral</strong> system (or system of numeration) is a <strong>framework</strong> where a set of numbers are 
represented by numerals in a consistent manner. It can be seen as the context that allows the numeral <strong>"11"</strong> to be interpreted
as the binary numeral for <strong>three</strong>, the decimal numeral for <strong>eleven</strong>, or other numbers in different bases.

Ideally, such a system will:

* Represent a useful set of numbers (e.g. all whole numbers, integers, or real numbers)

* Give every number represented a unique representation (or at least a standard representation)

* Reflect the algebraic and arithmetic structure of the numbers.

In mathematical numeral systems the <strong>base or radix</strong> is usually the number of unique digits, including zero, that a positional numeral system uses to represent numbers. 

For example, for the decimal system the radix is 10, because it uses the 10 digits from 0 through 9. 
When a number "hits" 9, the next number will not be another different symbol, but a "1" followed by a "0". 

In binary, the radix is 2, since after it hits "1", instead of "2" or another written symbol, it jumps straight to "10", followed by "11" and "100".

The highest symbol of a positional numeral system usually has the value one less than the value of the base of that numeral system. 

The standard positional numeral systems differ from one another only in the base they use.

The base is an integer that is greater than 1 (or less than negative 1), since a radix of zero would not have any digits, and a radix of 1 would only have the zero digit. Negative bases are rarely used.

In base-10 (decimal) positional notation, there are 10 decimal digits and the number

![decimal equation](https://latex.codecogs.com/svg.latex?2056=(2%20*%2010^3%20)+(5%20*%2010^2)%20+(0%20*%2010%20^1)+(6*10^0))

In base-16 (hexadecimal), there are 16 hexadecimal digits (0–9 and A–F) and the number

![hex equation](https://latex.codecogs.com/svg.latex?171B=(1*16^3)+(7*16^2)+(1*16^1)+(B*16^0)) (where B represents the number eleven as a single symbol)

In general, in base-b, there are b digits and the number

![base equation](https://latex.codecogs.com/svg.latex?a_3a_2a_1a_0=(a_3%20*%20b^3)+(a_2*b^2)+(a_1*b^1)+(a_0*b^0))
(Note that ![base digits](https://latex.codecogs.com/svg.latex?a_3a_2a_1a_0) represents a sequence of digits, not multiplication)

### What
This library, <strong>customnumber</strong> provides the ability to create custom positional numeral systems in an efficient and performant way.
You can create custom numbers based on custom numeral systems and use them at will. 

All you need is the possible values of a digit (e.g. 0123456789ABCDEF) and an initial number (e.g. 14FF)

To implement our HOW we utilize 2 standard library packages:
* [container/ring](https://golang.org/pkg/container/ring/) (circular list)
* [container/list](https://golang.org/pkg/container/list/) (doubly linked list)

Each digit represented as a circular list that contains the all the possible numerals.

Each number is represented as a doubly linked list of circular lists. 

When a digit rotates back to it's first digit as a result of an addition, then an arithmetic holding is generated
and because of the doubly linked list it rotates the next, in weight digit, once. The opposite thing happens when a subtraction is happening.

## Getting Started ##
All you need is at least <strong>GO 1.13</strong>
## Usage ##
Get the package
```bash
go get github.com/slysterous/custom-number
```
Then all you need to do is create a custom number
```gotemplate
// create a slice of runes.
digitValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

number := customnumber.NewNumber(digitValues, "128z")

// will make the number 1290.
number.Increment()

// will make the number 128y.
number.Decrement()

//will give you the string representation of the number.
strnumber:=number.String()
```
### Make Utilities
```bash
ci                             run ci
fmt                            gofmt all files excluding vendor
lint                           perform linting
test                           run tests
```
## Contributing ##
Refer to [Contributing](https://github.com/slysterous/custom-number/blob/main/CONTRIBUTING.md).
## Report bugs using Github's [issues](https://github.com/slysterous/custom-number/issues)
We use GitHub issues to track public bugs. Report a bug by [opening a new issue](https://github.com/slysterous/custom-number/issues);
## License ##
This library is distributed under the MIT license found in the [LICENSE](./LICENSE)
file.
