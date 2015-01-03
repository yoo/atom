package atom

import "github.com/gopherjs/gopherjs/js"

var atom js.Object

// Commands doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-commands
var Commands *CommandRegistry

// Workspace doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-workspace
var Workspace *WorkspaceT

func init() {
	atom = js.Global.Get("atom")
	Commands = new(CommandRegistry)
	Commands.o = atom.Get("commands")
	Workspace = new(WorkspaceT)
	Workspace.o = atom.Get("workspace")
}

func Activate(callback func()) {
	js.Module.Get("exports").Set("activate", callback)
}

// OnDidBeep doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-onDidBeep
func OnDidBeep(callback func()) Disposable {
	return &disposableImpl{o: atom.Call("onDidBeep", callback)}
}

// OnWillThrowError doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-onWillThrowError
// TODO: THIS NEEDS TESTING!
func OnWillThrowError(callback func(*Event)) Disposable {
	// catch the js event object and convert it into a Event type and call the orginal callback
	disposable := &disposableImpl{}
	disposable.o = atom.Call("onWillThrowError", func(jsEvent js.Object) {
		event := newEventFromJsObject(jsEvent)
		callback(event)
	})
	return disposable
}

// OnDidThrowError doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-onDidThrowError
// TODO: THIS NEEDS TESTING!
func OnDidThrowError(callback func(*Event)) Disposable {
	disposable := &disposableImpl{}
	disposable.o = atom.Call("onDidThrowError", func(jsEvent js.Object) {
		event := newEventFromJsObject(jsEvent)
		callback(event)
	})
	return disposable
}

// InDevMode doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-inDevMode
func InDevMode() bool {
	return atom.Call("inDevMode").Bool()
}

// InSafeMode doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-inSafeMode
func InSafeMode() bool {
	return atom.Call("inSafeMode").Bool()
}

// InSpecMode doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-inSpecMode
func InSpecMode() bool {
	return atom.Call("inSpecMode").Bool()
}

// GetVersion doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-getVersion
func GetVersion() string {
	return atom.Call("getVersion").Str()
}

// IsReleasedVersion doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-isReleasedVersion
func IsReleasedVersion() bool {
	return atom.Call("isReleasedVersion").Bool()
}

// GetWindowLoadTime doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-getWindowLoadTime
func GetWindowLoadTime() int {
	return atom.Call("getWindowLoadTime").Int()
}

// Open doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-open
func Open(options *OpenOptions) {
	if options == nil {
		options = &OpenOptions{}
	}
	atom.Call("open", options)
}

// Close doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-close
func Close() {
	atom.Call("close")
}

// GetSize doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-getSize
// The only difference is the values are retuned sperated not as an object.
func GetSize() (width, heigth int) {
	jsObject := atom.Call("getSize")
	width = jsObject.Get("width").Int()
	heigth = jsObject.Get("heigth").Int()
	return
}

// SetSize doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-setSize
func SetSize(width, heigth int) {
	atom.Call("setSize", width, heigth)
}

// GetPosition doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-getPosition
func GetPosition() (x, y int) {
	jsObject := atom.Call("getPosition")
	x = jsObject.Get("x").Int()
	y = jsObject.Get("y").Int()
	return
}

// Center doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-center
func Center() {
	atom.Call("center")
}

// Focus doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-focus
func Focus() {
	atom.Call("focus")
}

// Show doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-show
func Show() {
	atom.Call("show")
}

// Hide doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-hide
func Hide() {
	atom.Call("hide")
}

// Reload doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-reload
func Reload() {
	atom.Call("reload")
}

// IsMaximixed doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-isMaximixed
func IsMaximixed() bool {
	return atom.Call("isMaximixed").Bool()
}

// IsFullScreen doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-isFullScreen
func IsFullScreen() bool {
	return atom.Call("isFullScreen").Bool()
}

// SetFullScreen doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-setFullScreen
func SetFullScreen() {
	atom.Call("setFullScreen")
}

// ToggleFullScreen doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-toggleFullScreen
func ToggleFullScreen() {
	atom.Call("toggleFullScreen")
}

// Beep doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-beep
func Beep() {
	atom.Call("beep")
}

// Confirm doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-confirm
func Confirm(options *ConfirmOption) int {
	index := -1
	index = atom.Call("confirm", options.toJsObject()).Int()
	return index
}

// OpenDevTools doc at: https://atom.io/docs/api/v0.165.0/Atom#instance-openDevTools
func OpenDevTools() {
	atom.Call("openDevTools")
}

func ToggleDevTools() {
	atom.Call("toggleDevTools")
}

func ExecuteJavaScriptInDevTools() {
	atom.Call("executeJavaScriptInDevTools")
}
