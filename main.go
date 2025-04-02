package main

import (
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		printHelp()
		os.Exit(1)
	}
	if args[1] == "--help" || args[1] == "h" || args[1] == "-h" {
		printHelp()
		os.Exit(1)
	}

	num1, num2, op, err := parseArgs(args[1:])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	result, err := doMath(num1, num2, op)
	if err != nil {
		fmt.Println("Invalid operator: " + op)
		os.Exit(1)
	}
	fmt.Println(result)
}

func parseNum(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func parseArgs(args []string) (float64, float64, string, error) {
	var num1 float64
	var num2 float64
	var err error

	if len(args) == 1 {
		args, err = parseMathWhenItsABigOlBlob(args[0])
		if err != nil {
			return 0, 0, "", err
		}
	}

	// assume +
	if len(args) == 2 {
		args = []string{args[0], "+", args[1]}
	}

	if len(args) > 3 {
		return 0, 0, "", errInvalidNumberOfArguments
	}

	if len(args) == 3 {
		num1, err = parseNum(args[0])
		if err != nil {
			return 0, 0, "", fmt.Errorf("first argument %w", errArgumentIsNotANumber)
		}
		num2, err = parseNum(args[2])
		if err != nil {
			return 0, 0, "", fmt.Errorf("third argument %w", errArgumentIsNotANumber)
		}
		return num1, num2, args[1], nil
	}
	return num1, num2, "0", nil
}

func doMath(n1, n2 float64, operator string) (float64, error) {
	a := big.NewFloat(n1)
	b := big.NewFloat(n2)
	var result *big.Float

	if operator == "+" {
		result = new(big.Float).Add(a, b)
	}
	if operator == "-" {
		result = new(big.Float).Sub(a, b)
	}
	if result == nil {
		return 0, errors.New("invalid operator")
	}

	result.SetPrec(64)
	output, _ := strconv.ParseFloat(result.Text('f', 2), 64)
	return output, nil
}

func parseMathWhenItsABigOlBlob(blob string) ([]string, error) {
	parts := strings.Split(blob, "+")
	if len(parts) > 1 {
		return parts, nil
	}
	parts = strings.Split(blob, "-")
	if len(parts) > 1 {
		return parts, nil
	}
	return nil, errNoOperandInBlob
}

func printHelp() {
	fmt.Println("                          ==============================")
	fmt.Println("                                    MATH TIME!")
	fmt.Println("                          ==============================")
	fmt.Println("This program lets you do simple math from the command line in a few different ways")
	fmt.Println("")
	fmt.Println("math h")
	fmt.Println("math -h")
	fmt.Println("math --help     Print this help and exit")
	fmt.Println("")
	fmt.Println("Here's some examples:")
	fmt.Println("---------------------")
	fmt.Println("math 40 + 2      42")
	fmt.Println("math 40+2        42")
	fmt.Println("math 40 2        42 (note, without an operand, the program assumes addition)")
	fmt.Println("math 44-2        42")
}
