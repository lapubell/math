package main

import "errors"

var errInvalidNumberOfArguments = errors.New("invalid number of arguments")
var errArgumentIsNotANumber = errors.New("is not a number")
var errNoOperandInBlob = errors.New("couldn't find + or - in your args")
