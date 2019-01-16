package latte_go_verify

import (
	"errors"
	"fmt"
	"reflect"
)

var VerifyType = make(map[string]func(reflect.StructField, interface{}) error)

func Register(str string, fun func(reflect.StructField, interface{}) error) {
	VerifyType[str] = fun
}

type Verify struct {
	data interface{}
}

func (verify Verify) Set(key string, value interface{}) error {
	fmt.Println("type:", reflect.TypeOf(reflect.Indirect(reflect.ValueOf(verify.data)).Interface()))
	tt := reflect.TypeOf(reflect.Indirect(reflect.ValueOf(verify.data)).Interface())
	if field, boolean := tt.FieldByName(key); boolean {
		bb := field.Type
		fmt.Println(bb.String())
		if fun, b := VerifyType[bb.String()]; b {
			return fun(field, value)
		} else {
			return errors.New("not the type")
		}

	} else {
		return errors.New("not attribute")
	}
	//fmt.Println(tt.NumField(), )
	// for i := 0; i < tt.NumField(); i++ {
	// 	bb := tt.Field(i).Tag
	// 	fmt.Println("name", tt.Field(i).Name)
	// 	fmt.Println("attribute:", reflect.ValueOf(bb).String())
	// }
	return nil
}
