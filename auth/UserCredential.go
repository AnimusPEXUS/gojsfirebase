package auth

import (
	"syscall/js"

	firebase "github.com/AnimusPEXUS/gojsfirebase"
)

type UserCredential struct {
	JSValue *js.Value
}

func NewUserCredentialFromJSValue(jsvalue *js.Value) *UserCredential {
	return &UserCredential{JSValue: jsvalue}
}

func (self *UserCredential) GetUser() *firebase.User {

	user_js := self.JSValue.Get("user")

	ret := firebase.NewUserFromJSValue(&user_js)
	return ret
}
