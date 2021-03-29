package gojsfirebase

import (
	"syscall/js"

	"github.com/AnimusPEXUS/gojswebapi/promise"
)

type User struct {
	JSValue *js.Value
}

func NewUserFromJSValue(jsvalue *js.Value) *User {

	ret := &User{
		JSValue: jsvalue,
	}

	return ret
}

func (self *User) GetIdToken(forceRefresh *bool) *promise.Promise {
	forceRefresh_js := js.Undefined()
	if forceRefresh != nil {
		forceRefresh_js = js.ValueOf(*forceRefresh)
	}
	// TODO: errors

	var res_js js.Value

	if forceRefresh_js.IsUndefined() {
		res_js = self.JSValue.Call("getIdToken")
	} else {
		res_js = self.JSValue.Call("getIdToken", forceRefresh_js)
	}

	ret, _ := promise.NewPromiseFromJSValue(&res_js)

	return ret
}
