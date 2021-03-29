package auth

import (
	"syscall/js"

	"github.com/AnimusPEXUS/gojswebapi/promise"
)

type ConfirmationResult struct {
	JSValue *js.Value
}

func (self *ConfirmationResult) GetVerificationId() string {
	return self.JSValue.Get("verificationId").String()
}

func (self *ConfirmationResult) Confirm(verificationCode string) *promise.Promise {

	// TODO: result

	res := self.JSValue.Call("confirm", verificationCode)

	ret, _ := promise.NewPromiseFromJSValue(&res)

	return ret
}
