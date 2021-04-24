# numeral [![PkgGoDev](https://pkg.go.dev/badge/github.com/slysterous/numeral)](https://pkg.go.dev/github.com/slysterous/numeral)
[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://lbesson.mit-license.org/)
[![CircleCI](https://circleci.com/gh/slysterous/numeral.svg?style=shield)](https://circleci.com/gh/slysterous/numerals)
[![Coverage Status](https://coveralls.io/repos/github/slysterous/numeral/badge.svg?branch=main)](https://coveralls.io/github/slysterous/numeral?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/slysterous/numeral)](https://goreportcard.com/report/github.com/slysterous/numeral)

numerals based on custom positional (numeral) systems.

## 📝 About ##

### 🧠 Why
There are times, where we need to iterate over an X amount of possible combinations of values over a 
specific number or a specific string etc <strong> (iterators, rainbow tables, numeral system converters) </strong>

 You might also need to  have custom numerals and perform basic operations on them.

### ♟️ Numeral Positional Systems
A numeral is a symbol or group of symbols that represents a number. Numerals are not the same as numbers just as words are not the same with the things they refer to. The symbols "11(decimal)", "1011(binary)" and "B(hexadecimal)" are different numerals, all representing the same number. The number is the idea, the numeral is anything that represents that idea.

A positional numeral system denotes usually the extension to any base of the Hindu–Arabic numeral system (or decimal system). In a more general sense, a positional system is a numeral system in which the contribution of a digit to the value of a number is the value of the digit multiplied by a factor determined by the position of the digit.

In modern positional systems, such as the decimal system, the position of the digit means that its value must be multiplied by some value: in 555, the three identical symbols represent five hundreds, five tens, and five units, respectively, due to their different positions in the digit string.

![sa](https://cdn-images-1.medium.com/max/600/1*bXQb00XiL0am9VbiptYRDw.png)

According to its position of occurrence in the number, each digit is weighted. Towards the left, the weights increase by a constant factor equivalent to the base or radix. With the help of the radix point ('.'), the positions corresponding to integral weights (1) are differentiated from the positions corresponding to fractional weights (<1).

Any integer value that is greater than or equal to two can be used as the base or radix. 
The digit position 'n' has weight r ^n. The largest value of digit position is always 1 less than the base value. The value of a number is a weighted sum of its digits.
For example:

![decimal equation](https://latex.codecogs.com/png.latex?\bg_white&space;2056=(2%20*%2010^3%20)+(0%20*%2010^2)%20+(5%20*%2010%20^1)+(6*10^0))

Decimal (0–9 digits with radix 10) 2056 value breakdown

![hex equation](https://latex.codecogs.com/png.latex?\bg_white&space;171B=(1%20*%2016^3%20)+(7%20*%2016^2)%20+(1%20*%2016%20^1)+(B*16^0))


Hexadecimal(0–9 & A-F digits with radix 16) 171B value breakdown

![](https://cdn-images-1.medium.com/max/800/1*JrIpqe-RX5KgD-P6nQSyQw.png)

The general form of value breakdown, where b is the base or radix of the numeral system

### ✨ The Package
This package, <strong>numeral</strong> provides the ability to create custom positional numeral systems in an efficient and performant way.
You can create numerals based on custom numeral systems and use them at will. 

All you need is the possible values of a digit (e.g. 0123456789ABCDEF) and an initial number (e.g. 14FF)

To implement our HOW we utilize 2 standard library packages:
* [container/ring](https://golang.org/pkg/container/ring/) (circular list)
* [container/list](https://golang.org/pkg/container/list/) (doubly linked list)

Each digit represented as a circular list that contains the all the possible numeral.

Each number is represented as a doubly linked list of circular lists. 

When a digit rotates back to it's first digit as a result of an addition, then an arithmetic holding is generated
and because of the doubly linked list it rotates the next, in weight digit, once. The opposite thing happens when a subtraction is happening.

## 🎬 Getting Started ##
All you need is at least <strong>Go 1.13</strong>
## 🤓 Usage ##
Get the package
```bash
go get github.com/slysterous/numeral
```
Then all you need to do is create a numeral.
For full reference on usage and available methods visit [pkg.go.dev](https://pkg.go.dev/badge/github.com/slysterous/numeral)
```gotemplate
// create a slice of runes.
digitValues := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

number := numeral.NewNumeral(digitValues, "128z")

number2, err := numeral.NewFromDecimal(digitValues, 150)


//will give you the string representation of the number.
strnumber:=number2.String()

//will give you the decimal integer representation of the number.
intnumber:=number.Decimal()
```
### ⛩️ Make Utilities
```bash
ci                             run ci
fmt                            gofmt all files excluding vendor
lint                           perform linting
test                           run tests
```
## ℹ️ Contributing ##
Refer to [Contributing](https://github.com/slysterous/numeral/blob/main/CONTRIBUTING.md).
## 🐛 Report bugs using Github's [issues](https://github.com/slysterous/numeral/issues)
We use GitHub issues to track public bugs. Report a bug by [opening a new issue](https://github.com/slysterous/numeral/issues).
## ⚖️ License ##
This library is distributed under the MIT license found in the [LICENSE](./LICENSE)
file.
