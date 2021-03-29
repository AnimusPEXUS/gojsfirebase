package auth

import (
	"syscall/js"

	firebaseapp "github.com/AnimusPEXUS/gojsfirebase/app"
	"github.com/AnimusPEXUS/gojswebapi/promise"
)

type Auth struct {
	JSValue *js.Value
	Parent  *firebaseapp.App
}

func NewAuthFromApp(app *firebaseapp.App) (*Auth, error) {
	auth := app.JSValue.Call("auth")
	ret := &Auth{
		Parent:  app,
		JSValue: &auth,
	}
	return ret, nil
}

func (self *Auth) SignInWithPhoneNumber(
	phonnenumber string,
	app_vfy ApplicationVerifierI,
) (*promise.Promise, error) {
	promise_js := self.JSValue.Call(
		"signInWithPhoneNumber",
		phonnenumber,
		*app_vfy.GetJSValue(),
	)
	promise, err := promise.NewPromiseFromJSValue(&promise_js)
	if err != nil {
		return nil, err
	}
	return promise, nil
}
