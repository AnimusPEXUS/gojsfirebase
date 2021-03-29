package gojsfirebase

import (
	"syscall/js"

	"github.com/AnimusPEXUS/gojswebapi/promise"
)

type Auth struct {
	parent *App
	auth   js.Value
}

func (self *Auth) SignInWithPhoneNumber(
	phonnenumber string,
	app_vfy ApplicationVerifierI,
) (*promise.Promise, error) {
	promise_js := self.auth.Call("signInWithPhoneNumber", *app_vfy.GetJSValue())
	promise, err := promise.NewPromiseFromJSValue(&promise_js)
	if err != nil {
		return nil, err
	}
	return promise, nil
}
