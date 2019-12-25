package model

import (
	"testing"
	"io/ioutil"
)

func TestUnmarshal(t *testing.T) {
	type a struct {
		Name string `validate:"required"`
		Email string `validate:"required,email"`
	}
	files := []string{
		"./testdata/1.json",
		"./testdata/2.json",
		"./testdata/3.json",
		"./testdata/4.json",
	}	
	wants := []bool{
		false,
		true,
		true,
		true,
	}
	for i, v := range files {
		blob, _ := ioutil.ReadFile(v)
		var temp a
		err := UnmarshalAndValidate(blob, &temp)
		if (err != nil) != wants[i] {
			t.Errorf("the expected value is %t, but got %t at %d", wants[i], err != nil , i)
		}
	}
}