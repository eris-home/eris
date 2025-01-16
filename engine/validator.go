package engine

import (
	"fmt"
	"regexp"
)

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) ValidateInput(input string) bool {
	if len(input) == 0 {
		fmt.Println("Validator: Input is empty")
		return false
	}
	
	isValid, _ := regexp.MatchString(`^[a-zA-Z0-9\s]+$`, input)
	if !isValid {
		fmt.Println("Validator: Input contains invalid characters")
	}
	
	return isValid
}