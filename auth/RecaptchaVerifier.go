package auth

import (
	"errors"
	"syscall/js"

	firebaseapp "github.com/AnimusPEXUS/gojsfirebase/app"
)

type RecaptchaVerifier struct {
	JSValue *js.Value
}

func NewRecaptchaVerifier(
	container string,
	parameters map[string]interface{},
	app *firebaseapp.App,
) (*RecaptchaVerifier, error) {
	// auth *Auth,

	auth := app.Parent.JSValue.Get("auth")

	rv_js := auth.Get("RecaptchaVerifier")

	if rv_js.IsNull() || rv_js.IsUndefined() {
		return nil, errors.New("RecaptchaVerifier undefined")
	}

	rv := rv_js.New(container, js.ValueOf(parameters), *app.JSValue)

	ret := &RecaptchaVerifier{
		JSValue: &rv,
	}

	return ret, nil
}

func (self *RecaptchaVerifier) GetJSValue() *js.Value {
	return self.JSValue
}
