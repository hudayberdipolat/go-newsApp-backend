package validate

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func ValidateStruct(data interface{}) error {
	validate := validator.New()
	err := validate.Struct(data)
	return err
}

func ValidatePhoneNumber(phone string) bool {
	if string(phone[3]) == "6" {
		r := regexp.MustCompile("[9936]{4}[1-5][0-9]{6}$")
		return r.MatchString(phone)
	}
	if string(phone[3]) == "7" {
		return string(phone[4]) == "1"
	}
	return false
}
