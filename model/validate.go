package model

import(
 "gopkg.in/go-playground/validator.v10"
 "encoding/json"
)

var validate *validator.Validate

func initValidater(){
	if validate == nil {
		validate = validator.New()
	}
}

//Unmarshal this function constructs a struct and validates it.
func Unmarshal(blob []byte, s interface{}) error {
	err := json.Unmarshal(blob, &s)
	if err != nil {
		return err
	}
	initValidater()
	err = validate.Struct(s)
	if err != nil {
		return err
	}	
	return err
}

