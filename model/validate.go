package model

import (
	"encoding/json"
)

type ValidateError struct{}

func (v ValidateError) Error() string {
	return "バリデーションに失敗しました。"
}

type Validatable interface {
	Validate() error
}

//UnmarshalAndValidate this function constructs a struct and validates it.
func UnmarshalAndValidate(blob []byte, s Validatable) error {
	err := json.Unmarshal(blob, &s)
	if err != nil {
		return err
	}
	err = s.Validate()
	if err != nil {
		return err
	}
	return err
}
