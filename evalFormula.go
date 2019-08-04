package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Arithmetric abstraction
type Arithmetric interface {
	Sum() float64
	Subtract() float64
	Multiply() float64
	Divide() (float64, error)
}

// Formula format
type Formula struct {
	firstNumber, secondNumber float64
	operator                  string
}

// Sum operator
func (f *Formula) Sum() float64 {
	return f.firstNumber + f.secondNumber
}

// Subtract operator
func (f *Formula) Subtract() float64 {
	return f.firstNumber - f.secondNumber
}

// Multiply operator
func (f *Formula) Multiply() float64 {
	return f.firstNumber * f.secondNumber
}

// Divide operator
func (f *Formula) Divide() (float64, error) {
	if f.secondNumber == 0 {
		return 0, errors.New("Cannot divide a number by zero")
	}
	return f.firstNumber / f.secondNumber, nil
}

// calc result based on user formula
func calc(operator string, arithmetric Arithmetric) (float64, error) {
	switch operator {
	case "+":
		return arithmetric.Sum(), nil
	case "-":
		return arithmetric.Subtract(), nil
	case "*":
		return arithmetric.Multiply(), nil
	case "/":
		return arithmetric.Divide()
	default:
		return 0, errors.New("Invalid operator")
	}
}

// checkFormat check formula format
func checkFormat(formula *Formula, arr []string) {
	if len(arr) == 3 {
		// check valid number
		if first, err := strconv.ParseFloat(arr[0], 64); err == nil {
			formula.firstNumber = first
		} else {
			panic(err)
		}

		// check valid operator
		switch arr[1] {
		case "+", "-", "*", "/":
			formula.operator = arr[1]
		default:
			panic(errors.New("Invalid operator"))
		}

		// check valid number
		if second, err := strconv.ParseFloat(arr[2], 64); err == nil {
			formula.secondNumber = second
		} else {
			panic(err)
		}
	} else {
		panic(errors.New("Invalid formula format length"))
	}

}

func main() {
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// read user formula
		str := scanner.Text()
		arr := strings.Split(str, " ")

		// check formula format
		var formula Formula
		checkFormat(&formula, arr)

		// calc result
		result, err := calc(formula.operator, &formula)
		if err != nil {
			panic(err)
		} else {
			fmt.Println(result)
		}

		// type new formula or Ctrl+C to exit program
		fmt.Print("> ")
	}
}
