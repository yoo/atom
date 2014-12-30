package atom

import "github.com/gopherjs/gopherjs/js"

var atom js.Object

// https://atom.io/docs/api/v0.165.0/Atom#instance-commands
var Commands *CommandRegistry

// https://atom.io/docs/api/v0.165.0/Atom#instance-workspace
var Workspace *WorkspaceT

func init() {
	atom = js.Global.Get("atom")
	Commands = new(CommandRegistry)
	Commands.o = atom.Get("commands")
	Workspace = new(WorkspaceT)
	Workspace.o = atom.Get("workspace")
}

func Activate(f func()) {
	js.Module.Get("exports").Set("activate", f)
}
