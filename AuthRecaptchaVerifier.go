package gojsfirebase

import (
	"errors"
	"syscall/js"
)

type AuthRecaptchaVerifier struct {
	JSValue *js.Value
}

func NewAuthRecaptchaVerifier(
	container string,
	parameters *map[string]interface{},
	app *App,
) (*AuthRecaptchaVerifier, error) {

	RecaptchaVerifier := js.Global().Get("RecaptchaVerifier")
	if RecaptchaVerifier.IsNull() || RecaptchaVerifier.IsUndefined() {
		return nil, errors.New("couldn't instantinate Firebase Auth RecaptchaVerifier")
	}

	rv := RecaptchaVerifier.New(container, parameters, app)

	self := &AuthRecaptchaVerifier{
		JSValue: &rv,
	}

	return self, nil
}

func (self *AuthRecaptchaVerifier) GetJSValue() *js.Value {
	return self.JSValue
}
