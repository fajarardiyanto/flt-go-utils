package utils

import (
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

// default validator
var validate *validator.Validate

func init() {
	validate = validator.New()
	RegisterAll(validate)
}

func GetValidator() *validator.Validate {
	return validate
}

func RegisterAll(validate *validator.Validate) {
	validate.RegisterValidation("url-empty", UrlNotRequired)
	validate.RegisterValidation("duration", Duration)
	validate.RegisterValidation("jwtifnotempty", JwtIfNotEmpty)
	validate.RegisterValidation("emailoralphanum", EmailOrAlphaNum)
	validate.RegisterValidation("datetime", DataTime)
}

func DataTime(fl validator.FieldLevel) bool {
	spl := strings.Split(fl.Param(), ";")
	if len(spl) < 1 {
		return false
	}

	_, err := time.Parse(spl[0], fl.Field().String())
	return err == nil
}

func JwtIfNotEmpty(fl validator.FieldLevel) bool {
	if len(fl.Field().String()) == 0 {
		return true
	}
	err := GetValidator().Var(fl.Field().String(), "jwt")
	return err == nil
}

func Duration(fl validator.FieldLevel) bool {
	if len(fl.Field().String()) == 0 {
		return false
	}
	if _, err := time.ParseDuration(fl.Field().String()); err != nil {
		return false
	}
	return true
}

func UrlNotRequired(fl validator.FieldLevel) bool {
	if len(fl.Field().String()) == 0 {
		return true
	}

	if err := validate.Var(fl.Field().String(), "url"); err != nil {
		return false
	}
	return true
}

func EmailOrAlphaNum(fl validator.FieldLevel) bool {
	err := validate.Var(fl.Field().String(), "email")
	if err == nil {
		return true
	}
	err = validate.Var(fl.Field().String(), "alphanum")
	return err == nil
}
