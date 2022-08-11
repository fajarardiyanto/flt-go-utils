package utils

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorValidation struct {
	Type    string `json:"type"`
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Message string `json:"message,omitempty"`
}

func ParseValidatorError(err error) (res error) {
	if _, ok := err.(*validator.InvalidValidationError); ok {
		return err
	}

	if vals, ok := err.(validator.ValidationErrors); ok {
		var errrs []ErrorValidation
		for _, err := range vals {
			var message string
			switch err.Tag() {
			case "required":
				message = "this field can't be empty"
			case "datetime":
				strc := strings.Split(err.Param(), ";")
				if len(strc) == 2 {
					message = fmt.Sprintf("format must be %s", strc[1])
				}
			}
			errrs = append(errrs, ErrorValidation{
				Type:    "validation",
				Field:   err.Field(),
				Tag:     err.Tag(),
				Message: message,
			})
		}
		bss, _ := json.Marshal(errrs)
		return fmt.Errorf("%s", bss)
	}

	return err
}
