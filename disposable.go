package atom

import "github.com/gopherjs/gopherjs/js"

// Disposable doc at: https://atom.io/docs/api/v0.165.0/Disposable
type Disposable struct {
	o js.Object
}

// Constructor doc at: https://atom.io/docs/api/v0.165.0/Disposable#instance-constructor
func (d *Disposable) Constructor(disposalAction func()) {
	d.o.Set("constructor", disposalAction)
}

// Dispose doc at: https://atom.io/docs/api/v0.165.0/Disposable#instance-dispose
func (d *Disposable) Dispose() {
	d.o.Call("dispose")
}
