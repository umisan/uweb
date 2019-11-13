package test

import (
	"testing"
	"fmt"
)

//Errorfg this function executes Errorf function with the function name in which it is called and error info
func Errorfg(t *testing.T, info string, a ...interface{}) {
	funcName := t.Name()
	s := fmt.Sprintf(info, a...)
	t.Errorf("%s: %s", funcName, s)
}

//Fatalfg this function executes Fatalf function with the function name in which it is called and error info
func Fatalfg(t *testing.T, info string, a ...interface{}) {
	funcName := t.Name()
	s := fmt.Sprintf(info, a...)
	t.Fatalf("%s: %s", funcName, s)
}