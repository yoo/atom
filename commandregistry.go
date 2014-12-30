package atom

import "github.com/gopherjs/gopherjs/js"

type CommandRegistry struct {
	o js.Object
}

func (c *CommandRegistry) Add(target string, commandName string, callback func()) {
	c.o.Call("add", target, commandName, callback)
}
