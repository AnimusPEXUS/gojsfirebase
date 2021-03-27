package gojsfirebase

import "syscall/js"

type ApplicationVerifierI interface {
	GetJSValue() *js.Value
}
