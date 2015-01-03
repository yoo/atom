package atom

import "github.com/gopherjs/gopherjs/js"

// Disposable doc at: https://atom.io/docs/api/v0.165.0/Disposable
type Disposable interface {
	Dispose()
}

type disposableImpl struct {
	o js.Object
}

func (d *disposableImpl) Dispose() {
	d.o.Call("dispose")
}
