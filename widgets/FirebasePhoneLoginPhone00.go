package widgets

import (
	// "syscall/js"

	"github.com/AnimusPEXUS/gojsfirebase"
	"github.com/AnimusPEXUS/gojstools/elementtreeconstructor"
)

type FirebasePhoneLoginPhone00Options struct {
	Etc *elementtreeconstructor.ElementTreeConstructor

	LoadFirebaseAuthonomously bool
	GetAuthCB                 func() (*gojsfirebase.Auth, error)
}

type FirebasePhoneLoginPhone00 struct {
	options *FirebasePhoneLoginPhone00Options

	Element *elementtreeconstructor.ElementMutator
}

// type LoginPasswordForm00 struct {
// 	etc *elementtreeconstructor.ElementTreeConstructor

// 	put_submit_button bool

// 	onedit         func()
// 	onloginedit    func()
// 	onpasswordedit func()
// 	onsubmitclick  func()

// 	login_input    *elementtreeconstructor.ElementMutator
// 	password_input *elementtreeconstructor.ElementMutator

// 	Element *elementtreeconstructor.ElementMutator
// }

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
							SetAttribute("type", "text"),
					),
				etc.CreateElement("td").
					AppendChildren(
						etc.CreateElement("button").
							AppendChildren(
								etc.CreateTextNode("Begin Phone Authentication Process"),
							),
					),
			),
		etc.CreateElement("tr").
			AppendChildren(
				etc.CreateElement("td"),
			),
		etc.CreateElement("tr").
			AppendChildren(
				etc.CreateElement("td").
					AppendChildren(
						etc.CreateTextNode("Confirnation Code:"),
					),
				etc.CreateElement("td").
					AppendChildren(
						etc.CreateElement("input").
							SetAttribute("type", "text"),
					),
				etc.CreateElement("td").
					AppendChildren(
						etc.CreateElement("button").
							AppendChildren(
								etc.CreateTextNode("Begin Phone Authentication Process"),
							),
					),
			),
	)

	return self, nil
}
