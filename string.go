package latte_go_verify

import (
	"errors"
	"regexp"
)

type StringVerifyClass struct {
	MinLength int
	MaxLength int
	Regexp    string
}

func (class StringVerifyClass) Verify(data interface{}) error {
	var str = data.(string)
	if class.MaxLength != -1 {
		if len(str) > class.MaxLength {
			return errors.New("Len Exceeding the maximum")
		}
	}
	if class.MinLength != -1 {
		if len(str) < class.MinLength {
			return errors.New("Len Below the minimum")
		}
	}
	if class.Regexp != "" {
		if m, _ := regexp.MatchString(class.Regexp, str); !m {
			return errors.New("String does not match")
		}
	}
	return nil
}

type StringVerifyFactory struct {
}

func (verify StringVerifyFactory) Create(config map[string]interface{}) VerifyInterface {
	var stringVerify = StringVerifyClass{-1, -1, ""}
	if v, ok := config["maxLength"]; ok {
		stringVerify.MaxLength = v.(int)
	}
	if v, ok := config["minLength"]; ok {
		stringVerify.MinLength = v.(int)
	}
	if v, ok := config["regexp"]; ok {
		stringVerify.Regexp = v.(string)
	}
	return stringVerify
}

func init() {
	factory := StringVerifyFactory{}
	Register("string", factory)
}
