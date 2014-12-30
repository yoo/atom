package atom

import "github.com/gopherjs/gopherjs/js"

// Options doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-open
type Options struct {
	PathsToOpen []string
	NewWindow   bool
	DevMode     bool
	SafeMode    bool
}

func (o *Options) toJsObject() js.Object {
	var jsObject js.Object

	jsObject.Set("pathToOpen", o.PathsToOpen)
	jsObject.Set("newWindow", o.NewWindow)
	jsObject.Set("sevMode", o.DevMode)
	jsObject.Set("safeMode", o.SafeMode)

	return jsObject
}
