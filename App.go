package gojsfirebase

import (
	"syscall/js"
)

type App struct {
	parent *Firebase
	app    js.Value
	cfg    map[string]interface{}
}

func (self *App) Auth() (*Auth, error) {
	auth := self.app.Call("auth")
	ret := &Auth{
		parent: self,
		auth:   auth,
	}
	return ret, nil
}
