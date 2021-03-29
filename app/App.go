package app

import (
	"syscall/js"

	firebase "github.com/AnimusPEXUS/gojsfirebase"
)

type App struct {
	JSValue *js.Value
	Parent  *firebase.Firebase
	CFG     map[string]interface{}
}

func InitializeApp(fb *firebase.Firebase, cfg map[string]interface{}, name *string) (*App, error) {
	// TODO: exception handling

	name_js := js.Undefined()

	if name != nil {
		name_js = js.ValueOf(*name)
	}

	var app js.Value

	if name_js.IsUndefined() {
		app = fb.JSValue.Call("initializeApp", js.ValueOf(cfg))
	} else {
		app = fb.JSValue.Call("initializeApp", js.ValueOf(cfg), name_js)
	}

	ret := &App{
		JSValue: &app,
		Parent:  fb,
		CFG:     cfg,
	}

	return ret, nil
}
