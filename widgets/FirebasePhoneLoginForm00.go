package widgets

import (
	// "syscall/js"

	"syscall/js"

	"github.com/AnimusPEXUS/gojstools/elementtreeconstructor"
)

type FirebasePhoneLoginForm00Options struct {
	Etc *elementtreeconstructor.ElementTreeConstructor

	// LoadFirebaseCB    func() error
	// GetFirebaseAuthCB func() (*firebaseauth.Auth, error)
	// GetFirebaseAppCB  func() (*firebaseapp.App, error)

	// OnStartSuccess func(cr firebaseauth.ConfirmationResult)
	// OnStartFailure func()
	// OnEndSuccess   func(res *firebaseauth.UserCredential)
	// OnEndFailure   func()

	OnPhoneButtonClicked func(number string) error
	OnCodeButtonClicked  func(code string) error
}

type FirebasePhoneLoginForm00 struct {
	options *FirebasePhoneLoginForm00Options

	Element *elementtreeconstructor.ElementMutator

	PhoneInput   *elementtreeconstructor.ElementMutator
	CodeInput    *elementtreeconstructor.ElementMutator
	phone_button *elementtreeconstructor.ElementMutator
	code_button  *elementtreeconstructor.ElementMutator
}

func NewFirebasePhoneLoginForm00(options *FirebasePhoneLoginForm00Options) (
	*FirebasePhoneLoginForm00,
	error,
) {

	self := &FirebasePhoneLoginForm00{
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
							AssignSelf(&self.PhoneInput).
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
							AssignSelf(&self.CodeInput).
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

	self.phone_button.Set(
		"onclick",
		js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			if self.options.OnPhoneButtonClicked != nil {
				go self.options.OnPhoneButtonClicked(self.PhoneInput.GetJsValue("value").String())
			}
			return false
		}),
	)

	self.code_button.Set(
		"onclick",
		js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			if self.options.OnCodeButtonClicked != nil {
				go self.options.OnCodeButtonClicked(self.CodeInput.GetJsValue("value").String())
			}
			return false
		}),
	)

	return self, nil
}

func (self *FirebasePhoneLoginForm00) Destroy() {
	self.Element.Parent().Remove(self.Element)

	self.Element = nil

	self.PhoneInput = nil
	self.CodeInput = nil
	self.phone_button = nil
	self.code_button = nil

}
