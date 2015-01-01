package atom

import "github.com/gopherjs/gopherjs/js"

// OpenOptions doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-open
type OpenOptions struct {
	PathsToOpen []string `js:"pathsToOpen"`
	NewWindow   bool     `js:"newWindow"`
	DevMode     bool     `js:"devMode"`
	SafeMode    bool     `js:"safeMode"`
}

// ConfirmOption doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-confirm
// If the callbacks are set to nil the index of the pressed button is returned else 0.
// TODO: Is passing nil as function possible or dose it need to be a pointer?
type ConfirmOption struct {
	Message         string
	DetailedMessage string
	Buttons         []Button
}

func (m *ConfirmOption) toJsObject() js.Object {
	var jsObject js.Object
	jsObject.Set("message", m.Message)
	jsObject.Set("detailedMessage", m.DetailedMessage)

	if len(m.Buttons) > 0 {
		if m.Buttons[0].Callback == nil {
			jsObject.Set("buttons", buttonsArray(m.Buttons))
		} else {
			jsObject.Set("buttons", buttonsObject(m.Buttons))
		}
	}
	return jsObject
}

func buttonsArray(buttons []Button) js.Object {
	var jsButtons js.Object
	for i, button := range buttons {
		jsButtons.SetIndex(i, button.Name)
	}

	return jsButtons
}

func buttonsObject(buttons []Button) js.Object {
	var jsButtons js.Object
	for _, button := range buttons {
		jsButtons.Set(button.Name, button.Callback)
	}
	return jsButtons
}

// Button is a button inside ConfirmOptios.
type Button struct {
	// The name and text of the button.
	Name string
	// A callback function or nil.
	Callback *func()
}
