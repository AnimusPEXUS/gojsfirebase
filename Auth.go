package gojsfirebase

import "syscall/js"

type Auth struct {
	parent *App
	auth   js.Value
}

func (self *Auth) SignInWithPhoneNumber(
	phonnenumber string,
	app_vfy ApplicationVerifierI,
) error {
	return nil
}
