package latte_go_verify

import "testing"

func Test_StringVerify(t *testing.T) {
	verifyFactory := GetVerifyObject("string")
	if verifyFactory == nil {
		t.Log("string error")
		return
	}
	var config = map[string]interface{}{}
	config["minLength"] = 3
	config["maxLength"] = 4
	verify := verifyFactory.Create(config)
	err := verify.Verify("hello")
	if err != nil {
		t.Log("检验string失败")
		return
	}
	t.Log("验证string成功")
}

func Test_IntVerify(t *testing.T) {
	verifyFactory := GetVerifyObject("int")
	if verifyFactory == nil {
		t.Log("int error")
		return
	}
	var config = map[string]interface{}{}
	config["min"] = 5
	config["max"] = 6
	verify := verifyFactory.Create(config)
	err := verify.Verify(4)
	if err != nil {
		t.Log("检验int失败")
		return
	}
	t.Log("验证int成功")
}
