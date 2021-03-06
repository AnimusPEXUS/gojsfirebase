package gojsfirebase

import (
	"errors"
	"log"
	"syscall/js"

	"github.com/AnimusPEXUS/gojstools/elementtreeconstructor"
	"github.com/AnimusPEXUS/gojswebapi/dom"
)

// type FirebaseConfig struct {
// }

func loadPart(doc *dom.Document, version string, name string) error {
	etc := elementtreeconstructor.NewElementTreeConstructor(doc)
	body := doc.GetBody()

	waiter := make(chan struct{})

	js0 := etc.CreateElement("script")
	js0.SetAttribute("src", "https://www.gstatic.com/firebasejs/"+version+"/firebase-"+name+".js")
	js0.Set("onload", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		waiter <- struct{}{}
		return nil
	}))
	log.Println("firebase loading. waiting for", name)
	body_mutator := elementtreeconstructor.NewElementMutatorFromElement(body)
	body_mutator.AppendChildren(js0)
	<-waiter
	log.Println("    loaded", name)
	return nil
}

func LoadParts(doc *dom.Document, version string, parts ...string) error {
	for _, i := range parts {
		err := loadPart(doc, version, i)
		if err != nil {
			return err
		}
	}
	return nil
}

type Firebase struct {
	JSValue *js.Value
}

func NewFirebaseFromGlobalObjectP0() (*Firebase, error) {
	jsvalue := js.Global().Get("firebase")
	if jsvalue.IsUndefined() {
		return nil, errors.New("firebase variable is undefined")
	}
	return NewFirebaseFromGlobalObjectP1(&jsvalue)
}

func NewFirebaseFromGlobalObjectP1(jsvalue *js.Value) (*Firebase, error) {
	self := &Firebase{JSValue: jsvalue}
	return self, nil
}
