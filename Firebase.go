package gojsfirebase

import (
	"syscall/js"

	"github.com/AnimusPEXUS/gojstools/elementtreeconstructor"
	"github.com/AnimusPEXUS/gojswebapi/dom"
)

// type FirebaseConfig struct {
// }

type Firebase struct {
	JSValue js.Value
	Version string
}

func NewFirebaseFromGlobalObjectP0() (*Firebase, error) {
	jsvalue := js.Global().Get("firebase")
	return NewFirebaseFromGlobalObjectP1(jsvalue)
}

func NewFirebaseFromGlobalObjectP1(jsvalue js.Value) (*Firebase, error) {
	self := &Firebase{JSValue: jsvalue}
	return self, nil
}

func NewFirebase(version string) (*Firebase, error) {
	self := &Firebase{}
	self.Version = version
	return self, nil
}

func (self *Firebase) InitWithConfig() error {
	return nil
}

func (self *Firebase) loadPart(doc *dom.Document, name string) error {
	etc := elementtreeconstructor.NewElementTreeConstructor(doc)
	body := doc.GetBody()
	js0 := etc.CreateElement("script")
	js0.SetAttribute("src", "https://www.gstatic.com/firebasejs/"+self.Version+"/firebase-"+name+".js")
	body_mutator := elementtreeconstructor.NewElementMutatorFromElement(body)
	body_mutator.AppendChildren(js0)
	return nil
}

func (self *Firebase) LoadCore(doc *dom.Document) error {
	return self.loadPart(doc, "app")
}

func (self *Firebase) LoadAuth(doc *dom.Document) error {
	return self.loadPart(doc, "auth")
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
