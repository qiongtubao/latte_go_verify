package latte_go_verify

import "errors"

type IntVerifyFactory struct {
}
type IntVerifyClass struct {
	Min        int
	Max        int
	Expression string
}

func (verify IntVerifyClass) Verify(data interface{}) error {
	var i = data.(int)
	if verify.Min != -1 {
		if i < verify.Min {
			return errors.New("Len Below the minimum")
		}
	}
	if verify.Max != -1 {
		if i > verify.Max {
			return errors.New("Len Exceeding the maximum")
		}
	}
	return nil
}

func (verify IntVerifyFactory) Create(config map[string]interface{}) VerifyInterface {
	var intVerify = IntVerifyClass{-1, -1, ""}
	if v, ok := config["max"]; ok {
		intVerify.Max = v.(int)
	}
	if v, ok := config["min"]; ok {
		intVerify.Min = v.(int)
	}
	return intVerify
}

func init() {
	factory := IntVerifyFactory{}
	Register("int", factory)
}
