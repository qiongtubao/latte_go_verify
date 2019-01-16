package latte_go_verify

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func VerifyString(field reflect.StructField, value interface{}) error {
	str := value.(string)
	fmt.Println(str, field.Tag.Get("minLen"))
	if minLenStr := field.Tag.Get("minLen"); minLenStr != "" {
		minLen, err := strconv.Atoi(minLenStr)
		fmt.Println("minLen", minLen)
		if err != nil {
			return errors.New("string minLen")
		}
		if len(str) < minLen {
			return errors.New("Len Below the minimum")
		}
	}
	if maxLenStr := field.Tag.Get("maxLen"); maxLenStr != "" {
		maxLen, err := strconv.Atoi(maxLenStr)
		if err != nil {
			return errors.New("string maxLen attribute is not int class")
		}
		if len(str) > maxLen {
			return errors.New("Len Exceeding the maximum")
		}
	}
	return nil
}
func init() {
	fmt.Println("init....")
	Register("string", VerifyString)
}
