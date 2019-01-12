package latte_go_verify

type VerifyFactory interface {
	Create(map[string]interface{}) VerifyInterface
}
type VerifyInterface interface {
	Verify(interface{}) error
}

var VerifyType = make(map[string]VerifyFactory)

func GetVerifyObject(name string) VerifyFactory {
	if v, ok := VerifyType[name]; ok {
		return v
	}
	return nil
}

func Register(name string, Verify VerifyFactory) {
	VerifyType[name] = Verify
}
