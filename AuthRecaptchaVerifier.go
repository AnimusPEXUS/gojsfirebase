package gojsfirebase

import (
	"errors"
	"syscall/js"
)

type AuthRecaptchaVerifier struct {
	JSValue *js.Value
}

func (self *Auth) RecaptchaVerifier(
	container string,
	parameters *map[string]interface{},
	app *App,
) (*AuthRecaptchaVerifier, error) {
	RecaptchaVerifier := self.parent.parent.JSValue.Get("auth").Get("RecaptchaVerifier")
	// RecaptchaVerifier := self.auth.Get("RecaptchaVerifier")
	if RecaptchaVerifier.IsNull() || RecaptchaVerifier.IsUndefined() {
		return nil, errors.New("RecaptchaVerifier undefined")
	}

	rv := RecaptchaVerifier.New(container, parameters, app.app)

	ret := &AuthRecaptchaVerifier{
		JSValue: &rv,
	}

	return ret, nil
}

func (self *AuthRecaptchaVerifier) GetJSValue() *js.Value {
	return self.JSValue
}
