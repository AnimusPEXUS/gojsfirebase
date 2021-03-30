package widgets

import (
	// "syscall/js"

	"log"
	"syscall/js"

	firebaseapp "github.com/AnimusPEXUS/gojsfirebase/app"
	firebaseauth "github.com/AnimusPEXUS/gojsfirebase/auth"
	"github.com/AnimusPEXUS/gojstools/elementtreeconstructor"
	gojstoolsutils "github.com/AnimusPEXUS/gojstools/utils"
)

type FirebasePhoneLoginPhone00Options struct {
	Etc *elementtreeconstructor.ElementTreeConstructor

	// LoadFirebaseCB    func() error
	GetFirebaseAppCB  func() (*firebaseapp.App, error)
	GetFirebaseAuthCB func() (*firebaseauth.Auth, error)
	OnEndSuccess      func(res *firebaseauth.UserCredential)
	OnEndFailure      func()
}

type FirebasePhoneLoginPhone00 struct {
	options *FirebasePhoneLoginPhone00Options

	Element *elementtreeconstructor.ElementMutator

	phone_input  *elementtreeconstructor.ElementMutator
	code_input   *elementtreeconstructor.ElementMutator
	phone_button *elementtreeconstructor.ElementMutator
	code_button  *elementtreeconstructor.ElementMutator

	cr *firebaseauth.ConfirmationResult
}

func NewFirebasePhoneLoginPhone00(options *FirebasePhoneLoginPhone00Options) (
	*FirebasePhoneLoginPhone00,
	error,
) {

	self := &FirebasePhoneLoginPhone00{
		options: options,
	}

	etc := self.options.Etc

	self.Element = etc.CreateElement("table")

	self.Element.AppendChildren(
		etc.CreateElement("tr").
			AppendChildren(
				etc.CreateElement("td").
					AppendChildren(
						etc.CreateTextNode("Phone:"),
					),
				etc.CreateElement("td").
					AppendChildren(
						etc.CreateElement("input").
							AssignSelf(&self.phone_input).
							SetAttribute("type", "text"),
					),
				etc.CreateElement("td").
					AppendChildren(
						etc.CreateElement("button").
							AssignSelf(&self.phone_button).
							AppendChildren(
								etc.CreateTextNode("Begin Phone Authentication Process"),
							),
					),
			),
		etc.CreateElement("tr").
			AppendChildren(
				etc.CreateElement("td").
					SetAttribute("id", "phone-input-captcha-placement").
					SetAttribute("colspan", "3"),
			),
		etc.CreateElement("tr").
			AppendChildren(
				etc.CreateElement("td").
					AppendChildren(
						etc.CreateTextNode("Confirmation Code:"),
					),
				etc.CreateElement("td").
					AppendChildren(
						etc.CreateElement("input").
							AssignSelf(&self.code_input).
							SetAttribute("type", "text"),
					),
				etc.CreateElement("td").
					AppendChildren(
						etc.CreateElement("button").
							AssignSelf(&self.code_button).
							AppendChildren(
								etc.CreateTextNode("Submit"),
							),
					),
			),
	)

	self.phone_button.Set("onclick", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go self.onphoneclick()
		return false
	}))

	self.code_button.Set("onclick", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go self.oncodeclick()
		return false
	}))

	return self, nil
}

func (self *FirebasePhoneLoginPhone00) onphoneclick() {
	// TODO: error handling
	app, err := self.options.GetFirebaseAppCB()
	if err != nil {
		panic(err)
	}

	auth, err := self.options.GetFirebaseAuthCB()
	if err != nil {
		panic(err)
	}

	rvo := map[string]interface{}{}

	vfy, err := firebaseauth.NewRecaptchaVerifier(
		"phone-input-captcha-placement",
		rvo,
		app,
		auth,
	)
	if err != nil {
		panic(err)
	}

	phone := self.phone_input.SelfJsValue().Get("value").String()
	// log.Println("Phone is:", phone)

	promise, err := auth.SignInWithPhoneNumber(phone, vfy)
	if err != nil {
		panic(err)
	}
	promise.Then(
		gojstoolsutils.JSFuncLiteralToPointer(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			log.Println("firebase phone auth start success")
			if len(args) != 0 {
				cr := firebaseauth.NewConfirmationResultFromJSValue(
					gojstoolsutils.JSValueLiteralToPointer(args[0]),
				)
				self.cr = cr
			}
			return false
		})),
		gojstoolsutils.JSFuncLiteralToPointer(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			log.Println("firebase phone auth start failure")
			return false
		})),
	)
}

func (self *FirebasePhoneLoginPhone00) oncodeclick() {
	// TODO: error handling

	if self.cr == nil {
		log.Println("cr is not available")
		return
	}

	code := self.code_input.SelfJsValue().Get("value").String()

	promise := self.cr.Confirm(code)

	promise.Then(
		gojstoolsutils.JSFuncLiteralToPointer(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			log.Println("firebase phone auth end success")
			if len(args) != 0 {
				ucr := firebaseauth.NewUserCredentialFromJSValue(
					gojstoolsutils.JSValueLiteralToPointer(args[0]),
				)
				go self.options.OnEndSuccess(ucr)
			}
			return false
		})),
		gojstoolsutils.JSFuncLiteralToPointer(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			log.Println("firebase phone auth end failure")
			return false
		})),
	)

}

func (self *FirebasePhoneLoginPhone00) Destroy() {
	self.Element.Parent().Remove(self.Element)

	self.Element = nil

	self.phone_input = nil
	self.code_input = nil
	self.phone_button = nil
	self.code_button = nil
	self.cr = nil
}
