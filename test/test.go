package test

import (
	"testing"
)

//Errorfg this function executes Errorf function with the function name in which it is called and error info
func Errorfg(t *testing.T, info string) {
	funcName := t.Name()
	t.Errorf("%s: %s", funcName, info)
}

//Fatalfg this function executes Fatalf function with the function name in which it is called and error info
func Fatalfg(t *testing.T, info string) {
	funcName := t.Name()
	t.Fatalf("%s: %s", funcName, info)
}