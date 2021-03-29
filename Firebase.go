package gojsfirebase

import (
	"syscall/js"

	"github.com/AnimusPEXUS/gojstools/elementtreeconstructor"
	"github.com/AnimusPEXUS/gojswebapi/dom"
)

// type FirebaseConfig struct {
// }

func loadPart(doc *dom.Document, version string, name string) error {
	etc := elementtreeconstructor.NewElementTreeConstructor(doc)
	body := doc.GetBody()
	js0 := etc.CreateElement("script")
	js0.SetAttribute("src", "https://www.gstatic.com/firebasejs/"+version+"/firebase-"+name+".js")
	body_mutator := elementtreeconstructor.NewElementMutatorFromElement(body)
	body_mutator.AppendChildren(js0)
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
	return NewFirebaseFromGlobalObjectP1(&jsvalue)
}

func NewFirebaseFromGlobalObjectP1(jsvalue *js.Value) (*Firebase, error) {
	self := &Firebase{JSValue: jsvalue}
	return self, nil
}

func (self *Firebase) NewApp(cfg map[string]interface{}) (*App, error) {
	// TODO: exception handling
	app := self.JSValue.Call("initializeApp", js.ValueOf(cfg))

	ret := &App{
		parent: self,
		app:    app,
		cfg:    cfg,
	}
	return ret, nil
}
