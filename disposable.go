package atom

import "github.com/gopherjs/gopherjs/js"

// Disposable doc at: https://atom.io/docs/api/v0.165.0/Disposable
type Disposable interface {
	Dispose()
	getJsObject() js.Object
}

type disposableImpl struct {
	o js.Object
}

// Dispose doc at: https://atom.io/docs/api/v0.165.0/Disposable#instance-dispose
func (d *disposableImpl) Dispose() {
	d.o.Call("dispose")
}

func (d *disposableImpl) getJsObject() js.Object {
	return d.o
}
